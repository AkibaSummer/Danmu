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

// Gift SEND_GIFT 发送礼物类型
type SendGift struct {
	Cmd   string `json:"cmd"`
	Danmu struct {
		Area int `json:"area"`
	} `json:"danmu"`
	Data struct {
		Action         string      `json:"action"`
		BagGift        interface{} `json:"bag_gift"`
		BatchComboId   string      `json:"batch_combo_id"`
		BatchComboSend struct {
			Action        string      `json:"action"`
			BatchComboId  string      `json:"batch_combo_id"`
			BatchComboNum int         `json:"batch_combo_num"`
			BlindGift     interface{} `json:"blind_gift"`
			GiftId        int         `json:"gift_id"`
			GiftName      string      `json:"gift_name"`
			GiftNum       int         `json:"gift_num"`
			SendMaster    interface{} `json:"send_master"`
			Uid           int         `json:"uid"`
			Uname         string      `json:"uname"`
		} `json:"batch_combo_send"`
		BeatId           string      `json:"beatId"`
		BizSource        string      `json:"biz_source"`
		BlindGift        interface{} `json:"blind_gift"`
		BroadcastId      int         `json:"broadcast_id"`
		CoinType         string      `json:"coin_type"`
		ComboResourcesId int         `json:"combo_resources_id"`
		ComboSend        struct {
			Action     string      `json:"action"`
			ComboId    string      `json:"combo_id"`
			ComboNum   int         `json:"combo_num"`
			GiftId     int         `json:"gift_id"`
			GiftName   string      `json:"gift_name"`
			GiftNum    int         `json:"gift_num"`
			SendMaster interface{} `json:"send_master"`
			Uid        int         `json:"uid"`
			Uname      string      `json:"uname"`
		} `json:"combo_send"`
		ComboStayTime  int    `json:"combo_stay_time"`
		ComboTotalCoin int    `json:"combo_total_coin"`
		CritProb       int    `json:"crit_prob"`
		Demarcation    int    `json:"demarcation"`
		DiscountPrice  int    `json:"discount_price"`
		Dmscore        int    `json:"dmscore"`
		Draw           int    `json:"draw"`
		Effect         int    `json:"effect"`
		EffectBlock    int    `json:"effect_block"`
		Face           string `json:"face"`
		FaceEffectId   int    `json:"face_effect_id"`
		FaceEffectType int    `json:"face_effect_type"`
		FaceEffectV2   struct {
			Id   int `json:"id"`
			Type int `json:"type"`
		} `json:"face_effect_v2"`
		FloatScResourceId int    `json:"float_sc_resource_id"`
		GiftId            int    `json:"giftId"`
		GiftName          string `json:"giftName"`
		GiftType          int    `json:"giftType"`
		GiftInfo          struct {
			EffectId      int    `json:"effect_id"`
			Gif           string `json:"gif"`
			HasImagedGift int    `json:"has_imaged_gift"`
			ImgBasic      string `json:"img_basic"`
			Webp          string `json:"webp"`
		} `json:"gift_info"`
		GiftTag        []interface{} `json:"gift_tag"`
		Gold           int           `json:"gold"`
		GroupMedal     interface{}   `json:"group_medal"`
		GuardLevel     int           `json:"guard_level"`
		IsFirst        bool          `json:"is_first"`
		IsJoinReceiver bool          `json:"is_join_receiver"`
		IsNaming       bool          `json:"is_naming"`
		IsSpecialBatch int           `json:"is_special_batch"`
		Magnification  int           `json:"magnification"`
		MedalInfo      struct {
			AnchorRoomid     int    `json:"anchor_roomid"`
			AnchorUname      string `json:"anchor_uname"`
			GuardLevel       int    `json:"guard_level"`
			IconId           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       int    `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Special          string `json:"special"`
			TargetId         int    `json:"target_id"`
		} `json:"medal_info"`
		NameColor        string `json:"name_color"`
		Num              int    `json:"num"`
		OriginalGiftName string `json:"original_gift_name"`
		Price            int    `json:"price"`
		Rcost            int    `json:"rcost"`
		ReceiveUserInfo  struct {
			Uid   int    `json:"uid"`
			Uname string `json:"uname"`
		} `json:"receive_user_info"`
		ReceiverUinfo struct {
			Base struct {
				Face         string `json:"face"`
				IsMystery    bool   `json:"is_mystery"`
				Name         string `json:"name"`
				NameColor    int    `json:"name_color"`
				NameColorStr string `json:"name_color_str"`
				OfficialInfo struct {
					Desc  string `json:"desc"`
					Role  int    `json:"role"`
					Title string `json:"title"`
					Type  int    `json:"type"`
				} `json:"official_info"`
				OriginInfo struct {
					Face string `json:"face"`
					Name string `json:"name"`
				} `json:"origin_info"`
				RiskCtrlInfo struct {
					Face string `json:"face"`
					Name string `json:"name"`
				} `json:"risk_ctrl_info"`
			} `json:"base"`
			Guard       interface{} `json:"guard"`
			GuardLeader interface{} `json:"guard_leader"`
			Medal       interface{} `json:"medal"`
			Title       interface{} `json:"title"`
			UheadFrame  interface{} `json:"uhead_frame"`
			Uid         int         `json:"uid"`
			Wealth      interface{} `json:"wealth"`
		} `json:"receiver_uinfo"`
		Remain      int         `json:"remain"`
		Rnd         string      `json:"rnd"`
		SendMaster  interface{} `json:"send_master"`
		SenderUinfo struct {
			Base struct {
				Face         string `json:"face"`
				IsMystery    bool   `json:"is_mystery"`
				Name         string `json:"name"`
				NameColor    int    `json:"name_color"`
				NameColorStr string `json:"name_color_str"`
				OfficialInfo struct {
					Desc  string `json:"desc"`
					Role  int    `json:"role"`
					Title string `json:"title"`
					Type  int    `json:"type"`
				} `json:"official_info"`
				OriginInfo struct {
					Face string `json:"face"`
					Name string `json:"name"`
				} `json:"origin_info"`
				RiskCtrlInfo struct {
					Face string `json:"face"`
					Name string `json:"name"`
				} `json:"risk_ctrl_info"`
			} `json:"base"`
			Guard       interface{} `json:"guard"`
			GuardLeader interface{} `json:"guard_leader"`
			Medal       struct {
				Color              int    `json:"color"`
				ColorBorder        int    `json:"color_border"`
				ColorEnd           int    `json:"color_end"`
				ColorStart         int    `json:"color_start"`
				GuardIcon          string `json:"guard_icon"`
				GuardLevel         int    `json:"guard_level"`
				HonorIcon          string `json:"honor_icon"`
				Id                 int    `json:"id"`
				IsLight            int    `json:"is_light"`
				Level              int    `json:"level"`
				Name               string `json:"name"`
				Ruid               int    `json:"ruid"`
				Score              int    `json:"score"`
				Typ                int    `json:"typ"`
				UserReceiveCount   int    `json:"user_receive_count"`
				V2MedalColorBorder string `json:"v2_medal_color_border"`
				V2MedalColorEnd    string `json:"v2_medal_color_end"`
				V2MedalColorLevel  string `json:"v2_medal_color_level"`
				V2MedalColorStart  string `json:"v2_medal_color_start"`
				V2MedalColorText   string `json:"v2_medal_color_text"`
			} `json:"medal"`
			Title      interface{} `json:"title"`
			UheadFrame interface{} `json:"uhead_frame"`
			Uid        int         `json:"uid"`
			Wealth     interface{} `json:"wealth"`
		} `json:"sender_uinfo"`
		Silver            int         `json:"silver"`
		Super             int         `json:"super"`
		SuperBatchGiftNum int         `json:"super_batch_gift_num"`
		SuperGiftNum      int         `json:"super_gift_num"`
		SvgaBlock         int         `json:"svga_block"`
		Switch            bool        `json:"switch"`
		TagImage          string      `json:"tag_image"`
		Tid               string      `json:"tid"`
		Timestamp         int         `json:"timestamp"`
		TopList           interface{} `json:"top_list"`
		TotalCoin         int         `json:"total_coin"`
		Uid               int         `json:"uid"`
		Uname             string      `json:"uname"`
		WealthLevel       int         `json:"wealth_level"`
	} `json:"data"`
	MsgId    string `json:"msg_id"`
	PIsAck   bool   `json:"p_is_ack"`
	PMsgType int    `json:"p_msg_type"`
	SendTime int64  `json:"send_time"`
}

func (o *SendGift) String() string {
	return fmt.Sprintf("%s(%d) 赠送了 %d 个 %s", o.Data.SenderUinfo.Base.Name, o.Data.SenderUinfo.Uid, o.Data.Num, o.Data.GiftName)
}
