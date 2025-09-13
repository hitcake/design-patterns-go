package logger

type ConsoleLoggerFactory struct {
	level int
}

func (c *ConsoleLoggerFactory) GetLogger() (Logger, error) {
	return newConsoleLogger(c.level), nil
}
