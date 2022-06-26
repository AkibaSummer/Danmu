package structs

import (
	"fmt"
	"github.com/AkibaSummer/Danmu/sdk/utils"
	"github.com/thedevsaddam/gojsonq/v2"
)

// Common Struct

type CommentMeta struct { // 弹幕元信息
	Type      *int64
	TextSize  *int64
	Color     *int64
	Timestamp *int64
	ID        *int64
	SenderID  *string
}

func NewCommentMeta(prefix string, jsonQ *gojsonq.JSONQ) *CommentMeta {
	return &CommentMeta{
		Type:      utils.GetInt64(jsonQ.Copy().Find(prefix + ".[1]")),
		TextSize:  utils.GetInt64(jsonQ.Copy().Find(prefix + ".[2]")),
		Color:     utils.GetInt64(jsonQ.Copy().Find(prefix + ".[3]")),
		Timestamp: utils.GetInt64(jsonQ.Copy().Find(prefix + ".[4]")),
		ID:        utils.GetInt64(jsonQ.Copy().Find(prefix + ".[5]")),
		SenderID:  utils.GetString(jsonQ.Copy().Find(prefix + ".[7]")),
	}
}

type UserMeta struct { // 用户元信息
	UserID   *int64
	UserName *string
}

func NewUserMeta(prefix string, jsonQ *gojsonq.JSONQ) *UserMeta {
	return &UserMeta{
		UserID:   utils.GetInt64(jsonQ.Copy().Find(prefix + ".[0]")),
		UserName: utils.GetString(jsonQ.Copy().Find(prefix + ".[1]")),
	}
}

type GuardMeta struct { // 粉丝牌元信息
	GuardLevel *int64
	GuardName  *string
	UpName     *string
	UpRoomID   *int64
}

func NewGuardMeta(prefix string, jsonQ *gojsonq.JSONQ) *GuardMeta {
	return &GuardMeta{
		GuardLevel: utils.GetInt64(jsonQ.Copy().Find(prefix + ".[0]")),
		GuardName:  utils.GetString(jsonQ.Copy().Find(prefix + ".[1]")),
		UpName:     utils.GetString(jsonQ.Copy().Find(prefix + ".[2]")),
		UpRoomID:   utils.GetInt64(jsonQ.Copy().Find(prefix + ".[12]")),
	}
}

// Msg Struct

type Cmd struct {
	Cmd string `json:"cmd"`
	//Info []interface{} `json:"info"`
}

// Comment DANMU_MSG 普通弹幕类型
type Comment struct {
	CommentMeta *CommentMeta
	CommentText *string // 弹幕内容
	UserMeta    *UserMeta
	GuardMeta   *GuardMeta
}

func (o *Comment) String() string {
	str := ""
	if o.GuardMeta != nil && o.GuardMeta.GuardLevel != nil && o.GuardMeta.GuardName != nil {
		str = str + fmt.Sprintf("%v[%v]\t", *o.GuardMeta.GuardName, *o.GuardMeta.GuardLevel)
	}
	if o.UserMeta != nil && o.UserMeta.UserID != nil && o.UserMeta.UserName != nil {
		str = str + fmt.Sprintf("%v[%v]\t", *o.UserMeta.UserName, *o.UserMeta.UserID)
	}
	if o.CommentText != nil {
		str = str + *o.CommentText
	}

	return str
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
