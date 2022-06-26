package logger

import (
	"io"
	"sync"
)

type Writer struct {
	Input  InternalWriterChannel
	Output []io.Writer

	mu   sync.Mutex
	init bool
}

func NewWriter(input InternalWriterChannel, output []io.Writer) *Writer {
	return &Writer{Input: input, Output: output}
}

func NewWriterWithDefaultInput(output []io.Writer) *Writer {
	return &Writer{Input: make(InternalWriterChannel, DefaultChannelBuffer), Output: output}
}

func (o *Writer) GetInput() InternalWriterChannel {
	return o.Input
}

func (o *Writer) AddOutput(channel io.Writer) {
	o.Output = append(o.Output, channel)
}

func (o *Writer) Run() {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.init {
		return
	}
	go func() {
		for {
			select {
			case message := <-o.Input:
				if message.Skip {
					continue
				}
				for _, writer := range o.Output {
					_, _ = writer.Write([]byte(message.Message + "\n"))
				}
			}
		}
	}()
	o.init = true
}
