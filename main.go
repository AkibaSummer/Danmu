package main

import (
	"github.com/AkibaSummer/Danmu/sdk/spider"
	"github.com/AkibaSummer/Danmu/sdk/utils"
)

func main() {
	if utils.InitFlags() {
		spider.Init()
		_ = spider.NewDanmuSpider(*utils.RoomID, *utils.UID, *utils.BUVID, *utils.SESSDATA)
	}
}
