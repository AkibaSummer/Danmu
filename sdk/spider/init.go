package spider

import (
	"fmt"
	"github.com/AkibaSummer/Danmu/sdk/structs"
	"github.com/AkibaSummer/Danmu/sdk/utils"
	"github.com/AkibaSummer/Danmu/sdk/utils/logger"
	"io"
	"os"
	"time"
)

var (
	Info  logger.InternalLoggerChannel
	Debug logger.InternalLoggerChannel
)

func Init() {
	// Init system writer
	StdWriter := os.Stdout
	debugWriter, err := os.OpenFile("debugLogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utils.PanicIfNotNil(err)
	defaultWriter, err := os.OpenFile("defaultLogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	utils.PanicIfNotNil(err)

	// Init user writer
	debug := logger.NewWriterWithDefaultInput([]io.Writer{debugWriter})
	info := logger.NewWriterWithDefaultInput([]io.Writer{StdWriter, defaultWriter})
	debug.Run()
	info.Run()

	// Init logger
	debugLogger := logger.NewLoggerWithDefaultInput([]logger.InternalWriterChannel{debug.GetInput()}, []logger.LogParser{
		func(message *logger.InternalLoggerChannelMessage) *logger.InternalWriterChannelMessage {
			if message.Level.String() == "Unknown" {
				message.Level = logger.LevelDebug
			}
			return &logger.InternalWriterChannelMessage{Message: fmt.Sprintf("%s: %s %s %s", message.Level.String(), time.Now().Format("2006-01-02 15:04:05"), message.MessageType.String(), message.Message)}
		},
	})

	infoLogger := logger.NewLoggerWithDefaultInput([]logger.InternalWriterChannel{debug.GetInput(), info.GetInput()}, []logger.LogParser{
		func(message *logger.InternalLoggerChannelMessage) *logger.InternalWriterChannelMessage {
			if message.Level.String() == "Unknown" {
				message.Level = logger.LevelInfo
			}
			return &logger.InternalWriterChannelMessage{Message: fmt.Sprintf("%s: %s %s %s", message.Level.String(), time.Now().Format("2006-01-02 15:04:05"), message.MessageType.String(), message.Message)}
		},
		structs.ReadableParser,
	})

	debugLogger.Run()
	infoLogger.Run()

	Debug = debugLogger.GetInput()
	Info = infoLogger.GetInput()
}
