package logger

type FileLoggerFactory struct {
	level    int
	filePath string
}

func (f *FileLoggerFactory) GetLogger() (Logger, error) {
	return newFileLogger(f.filePath, f.level)
}
