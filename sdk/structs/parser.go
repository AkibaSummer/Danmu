package structs

import (
	"fmt"
	"github.com/AkibaSummer/Danmu/sdk/utils"
	"github.com/AkibaSummer/Danmu/sdk/utils/logger"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"time"
)

func ReadableParser(message *logger.InternalLoggerChannelMessage) (ret *logger.InternalWriterChannelMessage) {
	if message.Level.String() == "Unknown" {
		message.Level = logger.LevelInfo
	}
	defer func() {
		ret.Message = message.Level.String() + ": " + time.Now().Format("2006-01-02 15:04:05") + " " + ret.Message
	}()

	switch message.MessageType {
	case logger.TypeSystem:
	case logger.TypeMsg:
		msg := gojsonq.New().FromString(message.Message)
		cmd, ok := msg.Copy().Find("cmd").(string)
		if ok {
			switch cmd {
			case "DANMU_MSG":
				return logger.NewInternalWriterChannelMessage((&Comment{
					CommentMeta: NewCommentMeta("info.[0]", msg),
					CommentText: utils.GetString(msg.Copy().Find("info.[1]")),
					UserMeta:    NewUserMeta("info.[2]", msg),
					GuardMeta:   NewGuardMeta("info.[3]", msg),
				}).String())
			case "SEND_GIFT":
				return logger.NewInternalWriterChannelMessage(
					fmt.Sprintf("%s(%d) 赠送了 %d 个 %s",
						utils.Unptr(utils.GetString(msg.Copy().Find("data.sender_uinfo.base.name"))),
						utils.Unptr(utils.GetInt64(msg.Copy().Find("data.sender_uinfo.uid"))),
						utils.Unptr(utils.GetInt64(msg.Copy().Find("data.num"))),
						utils.Unptr(utils.GetString(msg.Copy().Find("data.giftName"))),
					),
				)
			case "INTERACT_WORD":
			//case "WATCHED_CHANGE":
			//case "STOP_LIVE_ROOM_LIST":
			default:
				return logger.NewInternalWriterChannelMessageNeedSkip()
			}
		}
	}

	return logger.NewInternalWriterChannelMessageNeedSkip()
}
