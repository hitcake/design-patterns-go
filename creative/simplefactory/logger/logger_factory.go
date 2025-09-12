package logger

import (
	"errors"
	"fmt"
)

const (
	CONSOLE = "console"
	FILE    = "file"
)

func GetLogger(loggerType string, filePath string, level int) (Logger, error) {
	switch loggerType {
	case CONSOLE:
		return NewConsoleLogger(level), nil
	case FILE:
		return NewFileLogger(filePath, level)
	}
	return nil, errors.New(fmt.Sprintf("logger type %s not support", loggerType))
}
