package structs

import (
	"github.com/AkibaSummer/Danmu/sdk/utils"
	"github.com/AkibaSummer/Danmu/sdk/utils/logger"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func ReadableParser(message *logger.InternalLoggerChannelMessage) (ret *logger.InternalWriterChannelMessage) {
	if message.Level.String() == "Unknown" {
		message.Level = logger.LevelInfo
	}
	defer func() {
		ret.Message = message.Level.String() + ": " + ret.Message
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
			//case "INTERACT_WORD":
			//case "WATCHED_CHANGE":
			//case "STOP_LIVE_ROOM_LIST":
			default:
				return logger.NewInternalWriterChannelMessageNeedSkip()
			}
		}
	}

	return logger.NewInternalWriterChannelMessageNeedSkip()
}
