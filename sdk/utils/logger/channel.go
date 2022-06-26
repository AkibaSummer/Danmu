package logger

import "fmt"

const (
	DefaultChannelBuffer = 1024
)

// Log Channel

type LogLevel int64

const (
	LevelDefault = LogLevel(iota)
	LevelTrace
	LevelDebug
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelFatal
)

func (l LogLevel) String() string {
	switch l {
	case LevelTrace:
		return "Trace"
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelNotice:
		return "Notice"
	case LevelWarn:
		return "Warn"
	case LevelError:
		return "Error"
	case LevelFatal:
		return "Fatal"
	}
	return "Unknown"
}

type LogType int64

const (
	TypeDefault = LogType(iota)
	TypeSystem
	TypeMsg
)

func (l LogType) String() string {
	switch l {
	case TypeSystem:
		return "System"
	case TypeMsg:
		return "Msg"
	}
	return "Unknown"
}

type InternalLoggerChannelMessage struct {
	Level       LogLevel
	MessageType LogType
	Message     string
}

func NewInternalLoggerChannelMessage(level LogLevel, messageType LogType, message string) *InternalLoggerChannelMessage {
	return &InternalLoggerChannelMessage{Level: level, MessageType: messageType, Message: message}
}
func NewSystemInternalLoggerChannelMessage(message ...any) *InternalLoggerChannelMessage {
	return &InternalLoggerChannelMessage{MessageType: TypeSystem, Message: fmt.Sprint(message...)}
}
func NewMsgInternalLoggerChannelMessage(message string) *InternalLoggerChannelMessage {
	return &InternalLoggerChannelMessage{MessageType: TypeMsg, Message: message}
}

type InternalLoggerChannel chan *InternalLoggerChannelMessage

func NewInternalLoggerChannel() InternalLoggerChannel {
	return make(InternalLoggerChannel, DefaultChannelBuffer)
}

// Writer Channel

type InternalWriterChannelMessage struct {
	Skip    bool
	Message string
}

func NewInternalWriterChannelMessageNeedSkip() *InternalWriterChannelMessage {
	return &InternalWriterChannelMessage{Skip: true}
}

func NewInternalWriterChannelMessage(message string) *InternalWriterChannelMessage {
	return &InternalWriterChannelMessage{Message: message}
}

type InternalWriterChannel chan *InternalWriterChannelMessage

func NewInternalWriterChannel() InternalWriterChannel {
	return make(InternalWriterChannel, 1024)
}
