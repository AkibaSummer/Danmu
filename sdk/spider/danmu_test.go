package spider

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"testing"
)

func TestDanmuSpider_Init(t *testing.T) {
	Init()
	_ = NewDanmuSpider(675014)
}

func TestJsonQ(t *testing.T) {
	const json = `{"cmd":"DANMU_MSG","info":[[0,1,25,16777215,1656249301670,0,0,"64e15af8",0,2,0,"",0,"{}","{}",{"mode":0,"show_player_type":0,"extra":"{\"send_from_me\":false,\"mode\":0,\"color\":16777215,\"dm_type\":0,\"font_size\":25,\"player_mode\":1,\"show_player_type\":0,\"content\":\"老板大气！点点红包抽礼物！\",\"user_hash\":\"1692490488\",\"emoticon_unique\":\"\",\"bulge_display\":0,\"recommend_score\":1,\"main_state_dm_color\":\"\",\"objective_state_dm_color\":\"\",\"direction\":0,\"pk_direction\":0,\"quartet_direction\":0,\"anniversary_crowd\":0,\"yeah_space_type\":\"\",\"yeah_space_url\":\"\",\"jump_to_url\":\"\",\"space_type\":\"\",\"space_url\":\"\"}"},{"activity_identity":"3802929","activity_source":2,"not_show":1}],"老板大气！点点红包抽礼物！",[1449658149,"旭出太洋",0,0,0,10000,1,""],[1,"法罗岛","AB电影馆",1857352,6067854,"",0,12632256,12632256,12632256,0,0,54489310],[0,0,9868950,"\u003e50000",0],["",""],0,0,null,{"ts":1656249301,"ct":"31B86770"},0,0,null,null,0,14]}`
	msg := gojsonq.New().FromString(json)
	//fmt.Println(msg.Find("cmd"))
	////fmt.Println(msg.Find("cmd"))
	//msg.Find("info")
	//fmt.Println(msg.Get())
	//msg = msg.Copy()
	//fmt.Println(msg.Get())
	//fmt.Println(msg.Find("cmd"))
	fmt.Println(msg.Find("info"))
	fmt.Println(msg.Find("[0]"))
	fmt.Println(msg.Find("[0]"))
	fmt.Println(msg.Copy().Find("[0]"))
	fmt.Println(msg.Get())
}
