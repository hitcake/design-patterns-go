package logger

import "testing"

func TestGetLogger(t *testing.T) {
	consoleLogger, err := GetLogger(CONSOLE, "", DEBUG)
	if err != nil {
		t.Error(err)
	}
	defer consoleLogger.Close()
	consoleLogger.Debug("test %s %d\n", "hello world", 123)
	fileLogger, err := GetLogger(FILE, "/Users/hit/GolandProjects/design-patterns-go/creative/simplefactory/logger/20250912.log", DEBUG)
	if err != nil {
		t.Error(err)
	}
	defer fileLogger.Close()
	fileLogger.Debug("test %s %d\n", "hello world", 123)

}
