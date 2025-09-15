package logger

import "testing"

func TestGetLogger(t *testing.T) {
	consoleLoggerFactory := &ConsoleLoggerFactory{level: DEBUG}
	consoleLogger, err := consoleLoggerFactory.GetLogger()
	if err != nil {
		t.Error(err)
	}
	defer consoleLogger.Close()
	consoleLogger.Debug("test %s %d\n", "hello world", 123)
	fileLoggerFactory := &FileLoggerFactory{filePath: "/Users/hit/GolandProjects/design-patterns-go/creative/factorymethod/logger/20250912.log", level: DEBUG}
	fileLogger, err := fileLoggerFactory.GetLogger()
	if err != nil {
		t.Error(err)
	}
	defer fileLogger.Close()
	fileLogger.Debug("test %s %d\n", "hello world", 123)

}
