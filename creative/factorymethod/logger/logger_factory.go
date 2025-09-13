package logger

type LoggerFactory interface {
	GetLogger() (Logger, error)
}
