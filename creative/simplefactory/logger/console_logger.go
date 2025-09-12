package logger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	*BaseLogger
}

func (cl *ConsoleLogger) Trace(format string, args ...interface{}) {
	if cl.LoggerEnabled(TRACE) {
		fmt.Printf(fmt.Sprintf("[TRACE] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
	}
}
func (cl *ConsoleLogger) Debug(format string, args ...interface{}) {
	if cl.LoggerEnabled(DEBUG) {
		fmt.Printf(fmt.Sprintf("[DEBUG] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
	}
}
func (cl *ConsoleLogger) Info(format string, args ...interface{}) {
	if cl.LoggerEnabled(INFO) {
		fmt.Printf(fmt.Sprintf("[INFO] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
	}
}
func (cl *ConsoleLogger) Warn(format string, args ...interface{}) {
	if cl.LoggerEnabled(WARN) {
		fmt.Printf(fmt.Sprintf("[WARN] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
	}
}

func (cl *ConsoleLogger) Error(format string, args ...interface{}) {
	if cl.LoggerEnabled(ERROR) {
		fmt.Printf(fmt.Sprintf("[ERROR] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
	}
}
func (cl *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if cl.LoggerEnabled(FATAL) {
		fmt.Printf(fmt.Sprintf("[FATAL] %s ", time.Now().Format("2006-01-02 15:04:05"))+format, args...)
	}
}
func (cl *ConsoleLogger) Close() {}
func NewConsoleLogger(level int) Logger {
	baseLogger := &BaseLogger{name: "console", level: level}
	return &ConsoleLogger{baseLogger}
}
