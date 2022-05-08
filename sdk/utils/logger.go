package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

var (
	Debug = log.New(io.Discard, "Debug: ", log.LstdFlags|log.Lshortfile)
	Info  = log.New(os.Stdout, "Info: ", log.LstdFlags|log.Lshortfile)
	Warn  = log.New(os.Stderr, "Warn: ", log.LstdFlags|log.Lshortfile)
	Error = log.New(os.Stderr, "Error: ", log.LstdFlags|log.Lshortfile)

	File *log.Logger
)

func Init() {
	logFile, err := os.OpenFile("unknownLogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	PanicIfNotNil(err)

	File = log.New(bufio.NewWriter(logFile), "Info: ", log.LstdFlags|log.Lshortfile)

	if *IsDebug {
		Debug.SetOutput(os.Stdout)
	}
}
