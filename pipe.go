package logging

type pipe struct {
	loggers []CoreLogger
}

// NewPipe - returns a pipe
func NewPipe(loggers []CoreLogger) CoreLogger {
	return &pipe{
		loggers: loggers,
	}
}

func (thisRef pipe) Log(logEntry LogEntry) LogEntry {
	for _, logger := range thisRef.loggers {
		logEntry = logger.Log(logEntry)
	}

	return logEntry
}
