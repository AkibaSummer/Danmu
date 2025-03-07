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
	now := time.Now()
	newPath := filepath.Join(lm.logDir, fmt.Sprintf("%d-%02d-%02d.log",
		now.Year(), now.Month(), now.Day()))

	latestLink := filepath.Join(lm.logDir, "latest.log")
	tempLink := filepath.Join(lm.logDir, "latest.log.tmp")

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

	// 获取绝对路径
	absNewPath, err := filepath.Abs(newPath)
	if err != nil {
		log.Fatal("获取绝对路径失败:", err)
	}

	lm.currentFile = file

	// 先创建临时链接
	os.Remove(tempLink)
	if err := os.Symlink(absNewPath, tempLink); err != nil {
		log.Printf("创建临时链接失败: %v", err)
		return
	}

	// 原子性地重命名临时链接
	if err := os.Rename(tempLink, latestLink); err != nil {
		log.Printf("重命名链接失败: %v", err)
		os.Remove(tempLink)
		return
	}
}

func (lm *LogManager) Close() error {
	lm.mutex.Lock()
	defer lm.mutex.Unlock()

	if lm.currentFile != nil {
		return lm.currentFile.Close()
	}
	return nil
}
