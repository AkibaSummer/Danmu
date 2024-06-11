package utils

import "flag"

var RoomID = flag.Int("roomID", 0, "直播间ID[required]")
var UID = flag.Int64("uid", 0, "用户UID[required]")
var BUVID = flag.String("buvid", "", "BUVID[required]")
var SESSDATA = flag.String("SESSDATA", "", "SESSDATA[required]")
var IsDebug = flag.Bool("debug", false, "是否开启调试（打印更多日志）")

func InitFlags() bool {
	flag.Parse()
	if RoomID == nil || *RoomID == 0 {
		flag.Usage()
		return false
	}
	if UID == nil || *UID == 0 {
		flag.Usage()
		return false
	}
	if BUVID == nil || *BUVID == "" {
		flag.Usage()
		return false
	}
	if SESSDATA == nil || *SESSDATA == "" {
		flag.Usage()
		return false
	}
	return true
}
