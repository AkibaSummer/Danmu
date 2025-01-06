package danmu_client

import (
	"fmt"
	"net/url"

	"github.com/AkibaSummer/Danmu/sdk/utils"
)

// from https://github.com/lovelyyoshino/Bilibili-Live-API/blob/master/API.WebSocket.md
const (
	WS_OP_HEARTBEAT                          = 2  //心跳
	WS_OP_HEARTBEAT_REPLY                    = 3  //心跳回应
	WS_OP_MESSAGE                            = 5  //弹幕消息等
	WS_OP_USER_AUTHENTICATION                = 7  //用户进入房间
	WS_OP_CONNECT_SUCCESS                    = 8  //进房回应
	WS_PACKAGE_HEADER_TOTAL_LENGTH           = 16 //头部字节大小
	WS_PACKAGE_OFFSET                        = 0
	WS_HEADER_OFFSET                         = 4
	WS_VERSION_OFFSET                        = 6
	WS_OPERATION_OFFSET                      = 8
	WS_SEQUENCE_OFFSET                       = 12
	WS_BODY_PROTOCOL_VERSION_NORMAL          = 0 //普通消息
	WS_BODY_PROTOCOL_VERSION_HEARTBEAT_REPLY = 1
	WS_BODY_PROTOCOL_VERSION_DEFLATE         = 2
	WS_BODY_PROTOCOL_VERSION_BROTLI          = 3 //brotli压缩信息
	WS_HEADER_DEFAULT_VERSION                = 1
	WS_HEADER_DEFAULT_OPERATION              = 1
	WS_HEADER_DEFAULT_SEQUENCE               = 1
	WS_AUTH_OK                               = 0
	WS_AUTH_TOKEN_ERROR                      = -101
)

//from player-loader-2.0.11.min.js
/*
	customAuthParam
*/
const (
	Protover = 2
	Platform = "web"
	Type     = 2
)

// API
const (
	getRoomByInfoBaseURL = "https://api.live.bilibili.com/xlive/web-room/v1/index/getInfoByRoom"
	getDanmuInfoBaseURL  = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo"
)

func getInfoByRoomURL(shortID any) string {
	params := url.Values{}

	Url, err := url.Parse(getRoomByInfoBaseURL)
	utils.PanicIfNotNil(err)

	params.Set("room_id", fmt.Sprint(shortID))

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	return urlPath
}

func getDanmuInfoURL(roomID any) string {
	params := url.Values{}

	Url, err := url.Parse(getDanmuInfoBaseURL)
	utils.PanicIfNotNil(err)

	params.Set("id", fmt.Sprint(roomID))
	params.Set("type", fmt.Sprint(0))

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	return urlPath
}
