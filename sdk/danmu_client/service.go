package danmu_client

import "github.com/go-playground/validator/v10"

type Client interface {
	Init() error
	GetLiveRoomInfo() LiveRoomInfo
	SetDanmuMsgCallback(func([]byte))
	SetLiveCallback(func())
	SetStopLiveCallback(func())
}

type Config struct {
	// user info
	UID      string `validate:"required,gt=0"`
	BUVID    string `validate:"required,gt=0"`
	SESSDATA string `validate:"required,gt=0"`

	// room info
	ShortID string `validate:"required,gt=0"`
}

type LiveRoomInfo struct {
	RoomID string

	Title          string
	AreaName       string
	ParentAreaName string
	LiveStatus     string // 1 开播 2 准备
	LiveStartTime  string // 单位 毫秒
}

func NewDanmuClient(conf Config) (Client, error) {
	err := validator.New().Struct(conf)
	if err != nil {
		return nil, err
	}
	return &DanmuClient{
		Config: conf,
	}, nil
}
