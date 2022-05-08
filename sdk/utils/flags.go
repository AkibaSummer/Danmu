package utils

import "flag"

var RoomID = flag.Int("roomID", 0, "直播间ID[required]")
var IsDebug = flag.Bool("debug", false, "是否开启调试（打印更多日志）")

func InitFlags() bool {
	flag.Parse()
	if RoomID == nil || *RoomID == 0 {
		flag.Usage()
		return false
	}
	return true
}
