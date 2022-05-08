package structs

import (
	"fmt"
	"github.com/AkibaSummer/Danmu/sdk/utils"
)

type Cmd struct {
	Cmd  string        `json:"cmd"`
	Info []interface{} `json:"info"`
}

// Comment DANMU_MSG 普通弹幕类型
type Comment struct {
	CommentText    string
	UserID         int64
	UserName       string
	IsAdmin        bool
	IsVIP          bool
	UserGuardLevel int64
}

func (o *Comment) String() string {
	return fmt.Sprintf("[%v]%v:\t%v", o.UserGuardLevel, o.UserName, o.CommentText)
}

type InteractMsgType int

func (o InteractMsgType) String() string {
	switch o {
	case 1:
		return "进入"
	case 2:
		return "关注"
	case 3:
		return "分享"
	case 4:
		return "特别关注"
	case 5:
		return "互相关注"
	}
	utils.Info.Println("未知互动类型:", o, "，记录于文件中待未来分析")
	utils.File.Println("未知互动类型：", o)
	return "<UNSET>"
}

// Interact INTERACT_WORD 观众互动信息
type Interact struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Contribution struct {
			Grade int `json:"grade"`
		} `json:"contribution"`
		Dmscore   int      `json:"dmscore"`
		FansMedal struct { // TODO: 怀疑可用于弹幕结构中
			AnchorRoomid     int    `json:"anchor_roomid"`
			GuardLevel       int    `json:"guard_level"`
			IconId           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       int    `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Score            int    `json:"score"`
			Special          string `json:"special"`
			TargetId         int    `json:"target_id"`
		} `json:"fans_medal"`
		Identities  []int           `json:"identities"`
		IsSpread    int             `json:"is_spread"`
		MsgType     InteractMsgType `json:"msg_type"`
		Roomid      int             `json:"roomid"`
		Score       int64           `json:"score"`
		SpreadDesc  string          `json:"spread_desc"`
		SpreadInfo  string          `json:"spread_info"`
		TailIcon    int             `json:"tail_icon"`
		Timestamp   int             `json:"timestamp"`
		TriggerTime int64           `json:"trigger_time"`
		Uid         int             `json:"uid"`
		Uname       string          `json:"uname"`
		UnameColor  string          `json:"uname_color"`
	} `json:"data"`
}

func (o *Interact) String() string {
	return fmt.Sprintf("%v %v了直播间", o.Data.Uname, o.Data.MsgType.String())
}

// WatchedChange WATCHED_CHANGE 观看人数变化
type WatchedChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Num       int    `json:"num"`
		TextSmall string `json:"text_small"`
		TextLarge string `json:"text_large"`
	} `json:"data"`
}

func (o *WatchedChange) String() string {
	return fmt.Sprintf("当前观看人次为 %v", o.Data.Num)
}

// StopLiveRoomList STOP_LIVE_ROOM_LIST 最近关闭的直播间
type StopLiveRoomList struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomIdList []int `json:"room_id_list"`
	} `json:"data"`
}

func (o *StopLiveRoomList) String() string {
	return fmt.Sprintf("有%d个直播间刚刚关闭", len(o.Data.RoomIdList))
}
