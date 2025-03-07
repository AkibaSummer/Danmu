package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type LogManager struct {
	logDir      string
	currentFile *os.File
	mutex       sync.Mutex
}

func NewLogManager(logDir string) *LogManager {
	manager := &LogManager{
		logDir: logDir,
	}
	manager.rotateFile()
	return manager
}

func (lm *LogManager) Write(p []byte) (n int, err error) {
	lm.mutex.Lock()
	defer lm.mutex.Unlock()

	lm.rotateFile()
	return lm.currentFile.Write(p)
}

func (lm *LogManager) rotateFile() {
	now := time.Now().Unix() / 60
	newPath := filepath.Join(lm.logDir, fmt.Sprintf("%d.log", now))
	latestLink := filepath.Join(lm.logDir, "latest.log")

	// 如果当前文件已经存在且是同一个文件，则不需要重新打开
	if lm.currentFile != nil {
		if currentPath, err := filepath.Abs(newPath); err == nil {
			if stat1, err1 := lm.currentFile.Stat(); err1 == nil {
				if stat2, err2 := os.Stat(currentPath); err2 == nil {
					if os.SameFile(stat1, stat2) {
						return
					}
				}
			}
		}
		lm.currentFile.Close()
	}

	if err := os.MkdirAll(lm.logDir, 0755); err != nil {
		log.Fatal("创建日志目录失败:", err)
	}

	file, err := os.OpenFile(newPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("打开日志文件失败:", err)
	}

	lm.currentFile = file
	// 创建或更新符号链接
	os.Remove(latestLink) // 删除已存在的链接
	os.Symlink(newPath, latestLink)
}

func (lm *LogManager) Close() error {
	lm.mutex.Lock()
	defer lm.mutex.Unlock()

	if lm.currentFile != nil {
		return lm.currentFile.Close()
	}
	return nil
}
