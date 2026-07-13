package spider

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/AkibaSummer/Danmu/sdk/structs"
	"github.com/AkibaSummer/Danmu/sdk/utils"
	"github.com/AkibaSummer/Danmu/sdk/utils/logger"
	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
)

type FailureKind string

const (
	FailureAuth     FailureKind = "auth"
	FailureNetwork  FailureKind = "network"
	FailureAPI      FailureKind = "api"
	FailureProtocol FailureKind = "protocol"
	FailureConfig   FailureKind = "config"
)

type Failure struct {
	Kind FailureKind
	Op   string
	Err  error
}

func (e *Failure) Error() string { return fmt.Sprintf("%s: %v", e.Op, e.Err) }
func (e *Failure) Unwrap() error { return e.Err }

func KindOf(err error) FailureKind {
	var failure *Failure
	if errors.As(err, &failure) {
		return failure.Kind
	}
	var netErr net.Error
	if errors.As(err, &netErr) {
		return FailureNetwork
	}
	return FailureProtocol
}

func failure(kind FailureKind, op string, err error) error {
	return &Failure{Kind: kind, Op: op, Err: err}
}

type DanmuSpider struct {
	ShortID int
	UID     int64
	BUVID   string
	Cookie  string

	RoomID            int
	Token             string
	WSURL             []string
	EndpointStateFile string

	dial    *websocket.Conn
	writeMu sync.Mutex
	http    *http.Client
}

func NewDanmuSpider(shortID int, uid int64, buvid, cookie string) *DanmuSpider {
	return &DanmuSpider{
		ShortID: shortID,
		UID:     uid,
		BUVID:   buvid,
		Cookie:  cookie,
		http:    &http.Client{Timeout: 15 * time.Second},
	}
}

type Message struct {
	PacketLen int32
	HeaderLen int16
	Ver       int16
	Op        int32
	Seq       int32
	Body      []byte
}

func Itob32(num int32) []byte {
	var b bytes.Buffer
	_ = binary.Write(&b, binary.BigEndian, num)
	return b.Bytes()
}
func Itob16(num int16) []byte {
	var b bytes.Buffer
	_ = binary.Write(&b, binary.BigEndian, num)
	return b.Bytes()
}

func EncodeMessage(msg string, operation int) []byte {
	var buffer bytes.Buffer
	byteMsg := []byte(msg)
	buffer.Write(Itob32(int32(len(byteMsg) + WS_PACKAGE_HEADER_TOTAL_LENGTH)))
	buffer.Write(Itob16(WS_PACKAGE_HEADER_TOTAL_LENGTH))
	buffer.Write(Itob16(WS_HEADER_DEFAULT_VERSION))
	buffer.Write(Itob32(int32(operation)))
	buffer.Write(Itob32(WS_HEADER_DEFAULT_SEQUENCE))
	buffer.Write(byteMsg)
	return buffer.Bytes()
}

func HelloGen(roomID int, uid int64, buvid, key string) []byte {
	obj := fmt.Sprintf(`{"roomid":%d,"uid":%d,"buvid":"%s","protover":3,"key":"%s","platform":"web","type":2}`, roomID, uid, buvid, key)
	return EncodeMessage(obj, WS_OP_USER_AUTHENTICATION)
}

func (d *DanmuSpider) MessageHandler(msg *Message) {
	Info <- logger.NewMsgInternalLoggerChannelMessage(string(msg.Body))
}

func (d *DanmuSpider) Run(ctx context.Context, connected func()) error {
	if d.ShortID <= 0 {
		return failure(FailureConfig, "validate room", errors.New("ShortID must be greater than zero"))
	}
	if err := d.discover(ctx); err != nil {
		return err
	}

	var lastErr error
	for _, host := range d.WSURL {
		u := url.URL{Scheme: "wss", Host: host, Path: "/sub"}
		Debug <- logger.NewSystemInternalLoggerChannelMessage("尝试建立连接: ", u.String(), " 房间ID: ", d.RoomID)
		dialer := *websocket.DefaultDialer
		dialer.HandshakeTimeout = 15 * time.Second
		conn, _, err := dialer.DialContext(ctx, u.String(), http.Header{"User-Agent": []string{userAgent}})
		if err != nil {
			lastErr = err
			continue
		}
		d.dial = conn
		err = d.runConnection(ctx, connected)
		_ = conn.Close()
		if err == nil || ctx.Err() != nil || KindOf(err) == FailureAuth {
			return err
		}
		lastErr = err
	}
	if lastErr == nil {
		lastErr = errors.New("API returned no websocket hosts")
	}
	return failure(FailureNetwork, "connect websocket hosts", lastErr)
}

func (d *DanmuSpider) discover(ctx context.Context) error {
	room := structs.NewGetInfoByRoomResp()
	roomURL, err := signedEndpoint(GetInfoByRoomURL(d.ShortID))
	if err != nil {
		return err
	}
	if err := d.getJSON(ctx, roomURL, "", room); err != nil {
		return err
	}
	if room.Code != 0 || room.Data.RoomInfo.RoomId == 0 {
		return failure(FailureAPI, "get room info", fmt.Errorf("code=%d message=%s room_id=%d", room.Code, room.Message, room.Data.RoomInfo.RoomId))
	}
	d.RoomID = room.Data.RoomInfo.RoomId

	info := structs.NewGetDanmuInfoResp()
	danmuURL, err := signedEndpoint(GetDanmuInfoURL(d.RoomID))
	if err != nil {
		return err
	}
	if err := d.getJSON(ctx, danmuURL, d.Cookie, info); err != nil {
		return err
	}
	if info.Code == -101 {
		return failure(FailureAuth, "get danmu token", errors.New("login expired"))
	}
	if info.Code != 0 || info.Data.Token == "" || len(info.Data.HostList) == 0 {
		apiErr := failure(FailureAPI, "get danmu token", fmt.Errorf("code=%d message=%s token=%t hosts=%d", info.Code, info.Message, info.Data.Token != "", len(info.Data.HostList)))
		if cacheErr := d.loadEndpointState(); cacheErr == nil {
			Debug <- logger.NewSystemInternalLoggerChannelMessage("弹幕节点接口受限，使用最近缓存的节点和令牌: ", apiErr)
			return nil
		}
		return apiErr
	}
	d.Token, d.WSURL = info.Data.Token, d.WSURL[:0]
	for _, host := range info.Data.HostList {
		if host.Host != "" {
			d.WSURL = append(d.WSURL, host.Host)
		}
	}
	if err := d.saveEndpointState(); err != nil {
		Debug <- logger.NewSystemInternalLoggerChannelMessage("保存弹幕节点缓存失败: ", err)
	}
	return nil
}

type EndpointState struct {
	RoomID    int       `json:"room_id"`
	Token     string    `json:"token"`
	Hosts     []string  `json:"hosts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FetchEndpoint(ctx context.Context, shortID int, uid int64, buvid, cookie string) (EndpointState, error) {
	client := NewDanmuSpider(shortID, uid, buvid, cookie)
	if err := client.discover(ctx); err != nil {
		return EndpointState{}, err
	}
	return EndpointState{RoomID: client.RoomID, Token: client.Token, Hosts: client.WSURL, UpdatedAt: time.Now()}, nil
}

func (d *DanmuSpider) loadEndpointState() error {
	if d.EndpointStateFile == "" {
		return errors.New("endpoint cache is disabled")
	}
	data, err := os.ReadFile(d.EndpointStateFile)
	if err != nil {
		return err
	}
	var state EndpointState
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}
	if state.RoomID != d.RoomID || state.Token == "" || len(state.Hosts) == 0 {
		return errors.New("endpoint cache is incomplete or for another room")
	}
	d.Token, d.WSURL = state.Token, append(d.WSURL[:0], state.Hosts...)
	return nil
}

func (d *DanmuSpider) saveEndpointState() error {
	if d.EndpointStateFile == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(d.EndpointStateFile), 0700); err != nil {
		return err
	}
	state := EndpointState{RoomID: d.RoomID, Token: d.Token, Hosts: d.WSURL, UpdatedAt: time.Now()}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	tmp := d.EndpointStateFile + ".tmp"
	if err := os.WriteFile(tmp, data, 0600); err != nil {
		return err
	}
	return os.Rename(tmp, d.EndpointStateFile)
}

func signedEndpoint(endpoint string) (string, error) {
	parsed, err := url.Parse(endpoint)
	if err != nil {
		return "", failure(FailureConfig, "parse API URL", err)
	}
	if err := utils.Sign(parsed); err != nil {
		return "", failure(FailureAPI, "sign API URL", err)
	}
	return parsed.String(), nil
}

func (d *DanmuSpider) getJSON(ctx context.Context, endpoint, cookie string, dst interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return failure(FailureConfig, "create HTTP request", err)
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", fmt.Sprintf("https://live.bilibili.com/%d", d.RoomID))
	req.Header.Set("Origin", "https://live.bilibili.com")
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	resp, err := d.http.Do(req)
	if err != nil {
		return failure(FailureNetwork, "HTTP GET", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return failure(FailureAPI, "HTTP GET", fmt.Errorf("status=%s", resp.Status))
	}
	decoder := json.NewDecoder(io.LimitReader(resp.Body, 4<<20))
	if err := decoder.Decode(dst); err != nil {
		return failure(FailureProtocol, "decode API JSON", err)
	}
	return nil
}

func (d *DanmuSpider) runConnection(ctx context.Context, connected func()) error {
	if err := d.write(websocket.BinaryMessage, HelloGen(d.RoomID, d.UID, d.BUVID, d.Token)); err != nil {
		return failure(FailureNetwork, "send authentication", err)
	}
	readErr := make(chan error, 1)
	authenticated := make(chan struct{}, 1)
	go d.readLoop(readErr, authenticated)

	select {
	case <-ctx.Done():
		return nil
	case err := <-readErr:
		return err
	case <-authenticated:
		Debug <- logger.NewSystemInternalLoggerChannelMessage("成功进入房间")
		if connected != nil {
			connected()
		}
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-readErr:
			return err
		case <-ticker.C:
			if err := d.write(websocket.BinaryMessage, EncodeMessage("[object Object]", WS_OP_HEARTBEAT)); err != nil {
				return failure(FailureNetwork, "send heartbeat", err)
			}
		}
	}
}

func (d *DanmuSpider) readLoop(result chan<- error, authenticated chan<- struct{}) {
	for {
		if err := d.dial.SetReadDeadline(time.Now().Add(75 * time.Second)); err != nil {
			result <- failure(FailureNetwork, "set read deadline", err)
			return
		}
		messageType, payload, err := d.dial.ReadMessage()
		if err != nil {
			result <- failure(FailureNetwork, "read websocket", err)
			return
		}
		switch messageType {
		case websocket.BinaryMessage:
			authOK, err := d.decodeMessage(payload)
			if err != nil {
				result <- err
				return
			}
			if authOK {
				select {
				case authenticated <- struct{}{}:
				default:
				}
			}
		case websocket.PingMessage:
			if err := d.write(websocket.PongMessage, payload); err != nil {
				result <- failure(FailureNetwork, "send pong", err)
				return
			}
		case websocket.CloseMessage:
			result <- failure(FailureNetwork, "read websocket", errors.New("server closed connection"))
			return
		}
	}
}

func (d *DanmuSpider) write(messageType int, payload []byte) error {
	d.writeMu.Lock()
	defer d.writeMu.Unlock()
	_ = d.dial.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return d.dial.WriteMessage(messageType, payload)
}

func (d *DanmuSpider) decodeMessage(payload []byte) (bool, error) {
	messages, err := unpackMessages(payload)
	if err != nil {
		return false, failure(FailureProtocol, "decode websocket packet", err)
	}
	authOK := false
	for _, msg := range messages {
		switch msg.Op {
		case WS_OP_HEARTBEAT_REPLY:
			if len(msg.Body) >= 4 {
				Info <- logger.NewSystemInternalLoggerChannelMessage("直播间人气: ", int32(binary.BigEndian.Uint32(msg.Body[:4])))
			}
		case WS_OP_MESSAGE:
			d.MessageHandler(&msg)
		case WS_OP_CONNECT_SUCCESS:
			var reply struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}
			if err := json.Unmarshal(msg.Body, &reply); err != nil {
				return false, failure(FailureProtocol, "decode auth reply", err)
			}
			if reply.Code != WS_AUTH_OK {
				return false, failure(FailureAuth, "websocket authentication", fmt.Errorf("code=%d message=%s", reply.Code, reply.Message))
			}
			authOK = true
		}
	}
	return authOK, nil
}

func unpackMessages(data []byte) ([]Message, error) {
	var result []Message
	for len(data) > 0 {
		if len(data) < WS_PACKAGE_HEADER_TOTAL_LENGTH {
			return nil, io.ErrUnexpectedEOF
		}
		packetLen := int(binary.BigEndian.Uint32(data[0:4]))
		headerLen := int(binary.BigEndian.Uint16(data[4:6]))
		if packetLen < headerLen || headerLen < WS_PACKAGE_HEADER_TOTAL_LENGTH || packetLen > len(data) {
			return nil, fmt.Errorf("invalid packet length packet=%d header=%d remaining=%d", packetLen, headerLen, len(data))
		}
		msg := Message{PacketLen: int32(packetLen), HeaderLen: int16(headerLen), Ver: int16(binary.BigEndian.Uint16(data[6:8])), Op: int32(binary.BigEndian.Uint32(data[8:12])), Seq: int32(binary.BigEndian.Uint32(data[12:16])), Body: data[headerLen:packetLen]}
		switch msg.Ver {
		case WS_BODY_PROTOCOL_VERSION_DEFLATE:
			reader, err := zlib.NewReader(bytes.NewReader(msg.Body))
			if err != nil {
				return nil, err
			}
			decoded, err := io.ReadAll(reader)
			closeErr := reader.Close()
			if err != nil {
				return nil, err
			}
			if closeErr != nil {
				return nil, closeErr
			}
			nested, err := unpackMessages(decoded)
			if err != nil {
				return nil, err
			}
			result = append(result, nested...)
		case WS_BODY_PROTOCOL_VERSION_BROTLI:
			decoded, err := io.ReadAll(brotli.NewReader(bytes.NewReader(msg.Body)))
			if err != nil {
				return nil, err
			}
			nested, err := unpackMessages(decoded)
			if err != nil {
				return nil, err
			}
			result = append(result, nested...)
		default:
			result = append(result, msg)
		}
		data = data[packetLen:]
	}
	return result, nil
}

// DecodeMessage remains for callers of the old SDK API. New code should use Run.
func (d *DanmuSpider) DecodeMessage(payload []byte) { _, _ = d.decodeMessage(payload) }
func (d *DanmuSpider) Send(msg []byte)              { _ = d.write(websocket.BinaryMessage, msg) }
func (d *DanmuSpider) HeartBeat() {
	_ = d.write(websocket.BinaryMessage, EncodeMessage("[object Object]", WS_OP_HEARTBEAT))
}
