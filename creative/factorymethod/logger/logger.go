package logger

type Logger interface {
	Trace(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type BaseLogger struct {
	name  string
	level int
}

func (bl *BaseLogger) LoggerEnabled(level int) bool {
	return bl.level <= level
}
