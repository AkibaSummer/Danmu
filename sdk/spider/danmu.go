package spider

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/AkibaSummer/Danmu/sdk/structs"
	"github.com/AkibaSummer/Danmu/sdk/utils"
	"github.com/AkibaSummer/Danmu/sdk/utils/logger"

	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type DanmuSpider struct {
	ShortID int

	Live              []string  //直播流链接
	Live_qn           int       //当前直播流质量
	Live_want_qn      int       //期望直播流质量
	RoomID            int       //房间id
	Title             string    //直播标题
	Uname             string    //主播名
	UpUid             int       //主播uid
	Rev               float64   //营收
	Renqi             int       //人气
	GuardNum          int       //舰长数
	ParentAreaID      int       //父分区
	AreaID            int       //子分区
	Locked            bool      //直播间封禁
	Note              string    //分区排行
	Live_Start_Time   time.Time //直播开始时间
	Liveing           bool      //是否在直播
	Wearing_FansMedal int       //当前佩戴的粉丝牌
	Token             string    //弹幕钥
	WSURL             []string  //弹幕链接
	LIVE_BUVID        bool      //cookies含LIVE_BUVID

	Dial *websocket.Conn
}

func NewDanmuSpider(shortId int) *DanmuSpider {
	ret := &DanmuSpider{ShortID: shortId}
	ret.Init()
	return ret
}

/*
	整数 字节转换区
	32 4字节
	16 2字节
*/
func Itob32(num int32) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	utils.PanicIfNotNil(err)
	return buffer.Bytes()
}

func Itob16(num int16) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	utils.PanicIfNotNil(err)
	return buffer.Bytes()
}

func btoi32(b []byte) int32 {
	var buffer int32
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	utils.PanicIfNotNil(err)
	return buffer
}

func btoi16(b []byte) int16 {
	var buffer int16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	utils.PanicIfNotNil(err)
	return buffer
}

func Btoi32(b []byte, offset int) int32 {
	return btoi32(b[offset : offset+4])
}

func Btoi16(b []byte, offset int) int16 {
	return btoi16(b[offset : offset+2])
}

//认证生成与检查
func HelloGen(roomid int, key string) []byte {
	if roomid == 0 || key == "" {
		return []byte("")
	}

	var obj = fmt.Sprintf(`{"roomid":%d,"uid":0,"protover":3,"key":"%s","platform":"danmuji","type":2}`, roomid, key)

	return EncodeMessage(obj, WS_OP_USER_AUTHENTICATION)
}

type Message struct {
	PacketLen int32
	HeaderLen int16
	Ver       int16
	Op        int32
	Seq       int32
	Body      []byte
}

func EncodeMessage(msg string, Operation int) []byte {
	var buffer bytes.Buffer
	byteMsg := []byte(msg)
	buffer.Write(Itob32(int32(len(byteMsg) + WS_PACKAGE_HEADER_TOTAL_LENGTH)))
	buffer.Write(Itob16(WS_PACKAGE_HEADER_TOTAL_LENGTH))
	buffer.Write(Itob16(WS_HEADER_DEFAULT_VERSION))
	buffer.Write(Itob32(int32(Operation)))
	buffer.Write(Itob32(int32(WS_HEADER_DEFAULT_SEQUENCE)))
	buffer.Write(byteMsg)
	return buffer.Bytes()
}

func (d *DanmuSpider) MessageHandler(msg *Message) {
	//cmd := structs.Cmd{}
	//utils.PanicIfNotNil(json.Unmarshal(msg.Body, &cmd))
	Info <- logger.NewMsgInternalLoggerChannelMessage(string(msg.Body))
	//switch cmd.Cmd {
	//case "DANMU_MSG":
	//	Comment := structs.Comment{}
	//	Comment.CommentText = cmd.Info[1].(string)
	//	Comment.UserID = int64(cmd.Info[2].([]interface{})[0].(float64))
	//	Comment.UserName = cmd.Info[2].([]interface{})[1].(string)
	//	Comment.IsAdmin = int64(cmd.Info[2].([]interface{})[2].(float64)) == 1
	//	Comment.IsVIP = int64(cmd.Info[2].([]interface{})[3].(float64)) == 1
	//	Comment.UserGuardLevel = int64(cmd.Info[7].(float64))
	//
	//case "INTERACT_WORD":
	//	Interact := structs.Interact{}
	//	utils.PanicIfNotNil(json.Unmarshal(msg.Body, &Interact))
	//	logger.Info.Println(Interact.String())
	//case "WATCHED_CHANGE":
	//	WatchedChange := structs.WatchedChange{}
	//	utils.PanicIfNotNil(json.Unmarshal(msg.Body, &WatchedChange))
	//	logger.Info.Println(WatchedChange.String())
	//case "STOP_LIVE_ROOM_LIST":
	//	StopLiveRoomList := structs.StopLiveRoomList{}
	//	utils.PanicIfNotNil(json.Unmarshal(msg.Body, &StopLiveRoomList))
	//	logger.Info <- logger.NewSystemInternalLoggerChannelMessage()
	//	Println(StopLiveRoomList.String())
	//default:
	//	logger.Debug <- logger.NewInternalLoggerChannelMessage(logger.LevelInfo, logger.TypeMsg, string(msg.Body))
	//}
}

func (d *DanmuSpider) DecodeMessage(msg []byte) {
	var err error
	m := Message{}
	reader := bytes.NewReader(msg)
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.PacketLen))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.HeaderLen))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.Ver))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.Op))
	utils.PanicIfNotNil(binary.Read(reader, binary.BigEndian, &m.Seq))
	m.Body, err = ioutil.ReadAll(reader)
	utils.PanicIfNotNil(err)
	switch m.Ver {
	case WS_BODY_PROTOCOL_VERSION_NORMAL:
	case WS_BODY_PROTOCOL_VERSION_HEARTBEAT_REPLY:
	case WS_BODY_PROTOCOL_VERSION_DEFLATE:
		zlibReader, err := zlib.NewReader(bytes.NewReader(m.Body))
		utils.PanicIfNotNil(err)
		m.Body, err = ioutil.ReadAll(zlibReader)
		utils.PanicIfNotNil(err)
	case WS_BODY_PROTOCOL_VERSION_BROTLI:
		m.Body, err = ioutil.ReadAll(brotli.NewReader(bytes.NewReader(m.Body)))
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
			d.MessageHandler(&m)
		} else {
			for bodyReader.Len() > 0 {
				m := Message{}
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
				d.MessageHandler(&m)
			}
		}
	case WS_OP_CONNECT_SUCCESS:
		Debug <- logger.NewSystemInternalLoggerChannelMessage("成功进入房间")
	default:
		Debug <- logger.NewSystemInternalLoggerChannelMessage("Unknown OpType", m.Op, string(m.Body))
	}
	//switch m.Op {
	//case WS_OP_MESSAGE:
	//	packetLen := 0
	//	for packet := binary.Read(reader, binary.BigEndian, packetLen)
	//case WS_OP_HEARTBEAT_REPLY:
	//	count := Btoi32(msg, 16)
	//	return fmt.Sprintf("当前人气数：%d", count)
	//}
	//
	//return buffer.Bytes()
}

func (d *DanmuSpider) Send(msg []byte) {
	utils.PanicIfNotNil(d.Dial.WriteMessage(websocket.TextMessage, msg))
}

func (d *DanmuSpider) HeartBeat() {
	Debug <- logger.NewSystemInternalLoggerChannelMessage("发送心跳包")
	d.Send(EncodeMessage("[object Object]", WS_OP_HEARTBEAT))
}

func (d *DanmuSpider) Init() {
	//Get RoomID
	{
		resp, err := http.Get(GetInfoByRoomURL(d.ShortID))
		utils.PanicIfNotNil(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		struttedResp := structs.NewGetInfoByRoomResp()
		utils.PanicIfNotNil(json.Unmarshal(body, &struttedResp))
		d.RoomID = struttedResp.Data.RoomInfo.RoomId
	}

	//Get DanmuServerURL
	{
		resp, err := http.Get(GetDanmuInfoURL(d.RoomID))
		utils.PanicIfNotNil(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		struttedResp := structs.NewGetDanmuInfoResp()
		utils.PanicIfNotNil(json.Unmarshal(body, &struttedResp))

		d.Token = struttedResp.Data.Token
		for i := range struttedResp.Data.HostList {
			d.WSURL = append(d.WSURL, struttedResp.Data.HostList[i].Host)
		}
	}

	// Connect To DanmuServer
	{
		u := url.URL{Scheme: "wss", Host: d.WSURL[0], Path: "/sub"}
		Debug <- logger.NewSystemInternalLoggerChannelMessage("尝试建立连接：", u.String(), "房间ID：", d.RoomID)
		var err error
		d.Dial, _, err = websocket.DefaultDialer.Dial(u.String(), make(http.Header))
		utils.PanicIfNotNil(err)
		defer d.Dial.Close()

		d.Send(HelloGen(d.RoomID, d.Token))

		done := make(chan struct{})
		go func() {
			defer close(done)
			for {
				utils.PanicIfNotNil(d.Dial.SetReadDeadline(time.Now().Add(time.Minute)))
				messageType, message, err := d.Dial.ReadMessage()
				if err != nil {
					if e, ok := err.(*websocket.CloseError); ok {
						switch e.Code {
						case websocket.CloseNormalClosure:
							Debug <- logger.NewSystemInternalLoggerChannelMessage("服务器连接关闭")
						case websocket.CloseAbnormalClosure:
							Debug <- logger.NewSystemInternalLoggerChannelMessage("服务器连接中断")
						default:
							Debug <- logger.NewSystemInternalLoggerChannelMessage("未知错误")
						}
					}
					return
				}
				utils.PanicIfNotNil(err)
				switch messageType {
				case websocket.TextMessage:
					Debug <- logger.NewSystemInternalLoggerChannelMessage("TextMessage Rec: ", string(message))
				case websocket.BinaryMessage:
					d.DecodeMessage(message)
				case websocket.CloseMessage:
					return
				case websocket.PingMessage:
					utils.PanicIfNotNil(d.Dial.WriteMessage(websocket.PongMessage, message))
				case websocket.PongMessage:
				}
			}
		}()

		ticker := time.NewTicker(time.Second * 30)
		defer ticker.Stop()

		for {
			select {
			case _ = <-ticker.C:
				d.HeartBeat()
			case <-done:
				return
			}
		}
	}
}
