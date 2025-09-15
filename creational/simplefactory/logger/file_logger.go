package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileLogger struct {
	*BaseLogger
	filePath string
	file     *os.File
}

func (l *FileLogger) Trace(format string, args ...interface{}) {
	if l.LoggerEnabled(TRACE) {
		line := fmt.Sprintf(fmt.Sprintf("[TRACE] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
		_, _ = l.file.WriteString(line)
	}
}

func (l *FileLogger) Debug(format string, args ...interface{}) {
	if l.LoggerEnabled(DEBUG) {
		line := fmt.Sprintf(fmt.Sprintf("[DEBUG] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
		_, _ = l.file.WriteString(line)
	}
}
func (l *FileLogger) Info(format string, args ...interface{}) {
	if l.LoggerEnabled(INFO) {
		line := fmt.Sprintf(fmt.Sprintf("[INFO] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
		_, _ = l.file.WriteString(line)
	}
}
func (l *FileLogger) Warn(format string, args ...interface{}) {
	if l.LoggerEnabled(WARN) {
		line := fmt.Sprintf(fmt.Sprintf("[WARN] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
		_, _ = l.file.WriteString(line)
	}
}
func (l *FileLogger) Error(format string, args ...interface{}) {
	if l.LoggerEnabled(ERROR) {
		line := fmt.Sprintf(fmt.Sprintf("[ERROR] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
		_, _ = l.file.WriteString(line)
	}
}
func (l *FileLogger) Fatal(format string, args ...interface{}) {
	if l.LoggerEnabled(FATAL) {
		line := fmt.Sprintf(fmt.Sprintf("[FATAL] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
		_, _ = l.file.WriteString(line)
	}
}
func (l *FileLogger) Close() {
	if l.file != nil {
		err := l.file.Close()
		if err != nil {
			panic(err)
		}
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("close failed", err)
		}
	}()
}

func newFileLogger(filePath string, level int) (Logger, error) {
	baseLogger := &BaseLogger{name: "file", level: level}
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// 打开文件，追加模式
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{baseLogger, filePath, file}, nil
}
