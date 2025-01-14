package main

import (
	"log"
	"time"

	"github.com/AkibaSummer/Danmu/sdk/spider"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read configure from conf/conf.yaml error: ", err)
	}

	spider.Init()
	for {
		_ = spider.NewDanmuSpider(
			viper.GetInt("bili.ShortID"),
			viper.GetInt64("bili.UID"),
			viper.GetString("bili.BUVID"),
			viper.GetString("bili.SESSDATA"),
		)
		time.Sleep(time.Second * 5)
		log.Println("Failed, retry")
	}

	//client, err := danmu_client.NewDanmuClient(danmu_client.Config{
	//	UID:      viper.GetString("bili.UID"),
	//	BUVID:    viper.GetString("bili.BUVID"),
	//	SESSDATA: viper.GetString("bili.SESSDATA"),
	//	ShortID:  viper.GetString("bili.ShortID"),
	//})
	//if err != nil {
	//	log.Fatal("Danmu client create failed: ", err)
	//}
	//err = client.Init()
	//if err != nil {
	//	log.Fatal("Danmu client init failed: ", err)
	//}
	//client.SetDanmuMsgCallback(func(b []byte) {
	//	log.Println(string(b))
	//})
	//client.SetLiveCallback(func() {
	//	log.Println("start live")
	//})
	//log.Println("Finish init")
	//time.Sleep(100000000000)
}
