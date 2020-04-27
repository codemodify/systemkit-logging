package logging

type emptyLogger struct{}

// NewEmptyLogger - returns a logger that does nothing
func NewEmptyLogger() CoreLogger {
	return &emptyLogger{}
}

func (thisRef emptyLogger) Log(logEntry LogEntry) LogEntry {
	return logEntry
}
