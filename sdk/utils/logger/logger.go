package logger

import (
	"sync"
)

type LogParser func(message *InternalLoggerChannelMessage) *InternalWriterChannelMessage

type Logger struct {
	Input  InternalLoggerChannel
	Output []InternalWriterChannel
	Parser []LogParser

	mu   sync.Mutex
	init bool
}

func NewLogger(input InternalLoggerChannel, output []InternalWriterChannel, parser []LogParser) *Logger {
	if len(output) != len(parser) {
		panic("len(output) != len(parser)")
	}
	return &Logger{Input: input, Output: output, Parser: parser}
}

func NewLoggerWithDefaultInput(output []InternalWriterChannel, parser []LogParser) *Logger {
	if len(output) != len(parser) {
		panic("len(output) != len(parser)")
	}
	return &Logger{Input: make(InternalLoggerChannel, DefaultChannelBuffer), Output: output, Parser: parser}
}

func (o *Logger) GetInput() InternalLoggerChannel {
	return o.Input
}

func (o *Logger) AddOutput(channel InternalWriterChannel, parser LogParser) {
	o.Output = append(o.Output, channel)
	o.Parser = append(o.Parser, parser)
}

func (o *Logger) Run() {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.init {
		return
	}
	go func() {
		for {
			select {
			case message := <-o.Input:
				for i, writer := range o.Output {
					writer <- o.Parser[i](message)
				}
			}
		}
	}()
	o.init = true
}
