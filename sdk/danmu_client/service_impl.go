package danmu_client

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/AkibaSummer/Danmu/sdk/utils"
	"github.com/AkibaSummer/Danmu/sdk/utils/logger"
	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"github.com/thedevsaddam/gojsonq/v2"
)

type DanmuClient struct {
	Config
	LiveRoomInfo
	DanmuInfo
}

type DanmuInfo struct {
	Token   string
	Host    string
	Port    string
	WssPort string
	WsPort  string

	Dial             *websocket.Conn
	msgCallback      func([]byte)
	liveCallback     func()
	stopLiveCallback func()
}

func (d *DanmuClient) httpGet(url string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getJsonPathAsString(bytes []byte, path string) string {
	result := gojsonq.New().FromString(string(bytes)).Reset().Find(path)
	return utils.JsonResToString(result)
}

func (d *DanmuClient) getRoomInfos(withLiveCallback bool) {
	body, err := d.httpGet(getInfoByRoomURL(d.ShortID), nil)
	utils.PanicIfNotNil(err)
	d.RoomID = getJsonPathAsString(body, "data.room_info.room_id")
	d.Title = getJsonPathAsString(body, "data.room_info.title")
	d.AreaName = getJsonPathAsString(body, "data.room_info.area_name")
	d.ParentAreaName = getJsonPathAsString(body, "data.room_info.parent_area_name")
	d.LiveStatus = getJsonPathAsString(body, "data.room_info.live_status")
	d.LiveStartTime = getJsonPathAsString(body, "data.room_info.live_start_time")
	if d.LiveStatus == "1" && withLiveCallback && d.liveCallback != nil {
		d.liveCallback()
	}
}

func (d *DanmuClient) getDanmuServer() {
	body, err := d.httpGet(getDanmuInfoURL(d.RoomID), map[string]string{
		"Cookie": fmt.Sprintf("SESSDATA=%s", d.SESSDATA),
	})
	utils.PanicIfNotNil(err)
	d.DanmuInfo.Host = getJsonPathAsString(body, "data.host_list.[0].host")
	d.DanmuInfo.Port = getJsonPathAsString(body, "data.host_list.[0].port")
	d.DanmuInfo.WssPort = getJsonPathAsString(body, "data.host_list.[0].wss_port")
	d.DanmuInfo.WsPort = getJsonPathAsString(body, "data.host_list.[0].ws_port")
	d.DanmuInfo.Token = getJsonPathAsString(body, "data.token")
}

// 认证生成与检查
func helloGen(roomid string, uid string, buvid string, key string) []byte {
	if roomid == "" || buvid == "" || key == "" {
		return []byte("")
	}

	var obj = fmt.Sprintf(`{"roomid":%s,"uid":%s,"buvid":"%s","protover":3,"key":"%s","platform":"web","type":2}`, roomid, uid, buvid, key)
	return encodeMessage(obj, WS_OP_USER_AUTHENTICATION)
}

type message struct {
	PacketLen int32
	HeaderLen int16
	Ver       int16
	Op        int32
	Seq       int32
	Body      []byte
}

func itob32(num int32) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	utils.PanicIfNotNil(err)
	return buffer.Bytes()
}

func itob16(num int16) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	utils.PanicIfNotNil(err)
	return buffer.Bytes()
}

func encodeMessage(msg string, Operation int) []byte {
	var buffer bytes.Buffer
	byteMsg := []byte(msg)
	buffer.Write(itob32(int32(len(byteMsg) + WS_PACKAGE_HEADER_TOTAL_LENGTH)))
	buffer.Write(itob16(WS_PACKAGE_HEADER_TOTAL_LENGTH))
	buffer.Write(itob16(WS_HEADER_DEFAULT_VERSION))
	buffer.Write(itob32(int32(Operation)))
	buffer.Write(itob32(int32(WS_HEADER_DEFAULT_SEQUENCE)))
	buffer.Write(byteMsg)
	return buffer.Bytes()
}

func (d *DanmuClient) send(msg []byte) {
	utils.PanicIfNotNil(d.Dial.WriteMessage(websocket.TextMessage, msg))
}

func (d *DanmuClient) handleMessage(msg []byte) {
	if d.msgCallback != nil {
		d.msgCallback(msg)
	}
	cmd := getJsonPathAsString(msg, "cmd")
	switch cmd {
	case "LIVE":
		d.getRoomInfos(d.LiveStatus != "1")
	case "PREPARING":
		d.LiveStatus = "2"
		if d.stopLiveCallback != nil {
			d.stopLiveCallback()
		}
	}
}

func (d *DanmuClient) decodeMessage(msg []byte) {
	var err error
	m := message{}
	reader := bytes.NewReader(msg)
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.PacketLen))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.HeaderLen))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.Ver))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.Op))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.Seq))
	m.Body, err = io.ReadAll(reader)
	utils.PanicIfNotNil(err)
	switch m.Ver {
	case WS_BODY_PROTOCOL_VERSION_NORMAL:
	case WS_BODY_PROTOCOL_VERSION_HEARTBEAT_REPLY:
	case WS_BODY_PROTOCOL_VERSION_DEFLATE:
		zlibReader, err := zlib.NewReader(bytes.NewReader(m.Body))
		utils.PanicIfNotNil(err)
		m.Body, err = io.ReadAll(zlibReader)
		utils.PanicIfNotNil(err)
	case WS_BODY_PROTOCOL_VERSION_BROTLI:
		m.Body, err = io.ReadAll(brotli.NewReader(bytes.NewReader(m.Body)))
		utils.PanicIfNotNil(err)
	}
	bodyReader := bytes.NewReader(m.Body)
	switch m.Op {
	case WS_OP_HEARTBEAT_REPLY:
		var count int32
		utils.PanicIfNotNil(binary.Read(bodyReader, binary.BigEndian, &count))
		logger.NewSystemInternalLoggerChannelMessage("直播间人气:", count)
	case WS_OP_MESSAGE:
		if m.Ver == WS_BODY_PROTOCOL_VERSION_NORMAL {
			d.handleMessage(m.Body)
		} else {
			for bodyReader.Len() > 0 {
				m := message{}
				utils.PanicIfNotNil(binary.Read(bodyReader, binary.BigEndian, &m.PacketLen))
				utils.PanicIfNotNil(binary.Read(bodyReader, binary.BigEndian, &m.HeaderLen))
				utils.PanicIfNotNil(binary.Read(bodyReader, binary.BigEndian, &m.Ver))
				utils.PanicIfNotNil(binary.Read(bodyReader, binary.BigEndian, &m.Op))
				utils.PanicIfNotNil(binary.Read(bodyReader, binary.BigEndian, &m.Seq))
				m.Body = make([]byte, m.PacketLen-int32(m.HeaderLen))
				n, err := bodyReader.Read(m.Body)
				utils.PanicIfNotNil(err)
				if n != int(m.PacketLen-int32(m.HeaderLen)) {
					panic("数据包读取长度不正确")
				}
				d.handleMessage(m.Body)
			}
		}
	case WS_OP_CONNECT_SUCCESS:
		log.Println("成功进入房间")
	default:
		log.Println("Unknown OpType", m.Op, string(m.Body))
	}
}

func (d *DanmuClient) Init() error {
	d.getRoomInfos(d.LiveStatus != "1")
	d.getDanmuServer()

	u := url.URL{Scheme: "wss", Host: d.Host, Path: "/sub"}
	var err error

	d.Dial, _, err = websocket.DefaultDialer.Dial(u.String(), make(http.Header))
	utils.PanicIfNotNil(err)

	d.send(helloGen(d.RoomID, d.UID, d.BUVID, d.DanmuInfo.Token))

	done := make(chan struct{})
	go func() {
		defer d.Dial.Close()
		defer close(done)
		for {
			utils.PanicIfNotNil(d.Dial.SetReadDeadline(time.Now().Add(time.Minute)))
			messageType, message, err := d.Dial.ReadMessage()
			if err != nil {
				if e, ok := err.(*websocket.CloseError); ok {
					switch e.Code {
					case websocket.CloseNormalClosure:
						log.Println("服务器连接关闭")
					case websocket.CloseAbnormalClosure:
						log.Println("服务器连接中断", e.Text)
					default:
						log.Println("未知错误")
					}
				}
				return
			}
			utils.PanicIfNotNil(err)
			switch messageType {
			case websocket.TextMessage:
			case websocket.BinaryMessage:
				d.decodeMessage(message)
			case websocket.CloseMessage:
				panic("websocket.CloseMessage")
			case websocket.PingMessage:
				utils.PanicIfNotNil(d.Dial.WriteMessage(websocket.PongMessage, message))
			case websocket.PongMessage:
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 30)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				d.send(encodeMessage("[object Object]", WS_OP_HEARTBEAT)) // 发送心跳包
			case <-done:
				return
			}
		}
	}()

	return nil
}

func (d *DanmuClient) GetLiveRoomInfo() LiveRoomInfo {
	return d.LiveRoomInfo
}

func (d *DanmuClient) SetDanmuMsgCallback(callback func([]byte)) {
	d.msgCallback = callback
}

func (d *DanmuClient) SetLiveCallback(callback func()) {
	d.liveCallback = callback
}

func (d *DanmuClient) SetStopLiveCallback(callback func()) {
	d.stopLiveCallback = callback
}
