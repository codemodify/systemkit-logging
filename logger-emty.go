package logging

// NewEmptyLogger - returns a logger that does nothing
func NewEmptyLogger() CoreLogger {
	return &emptyLogger{}
}

type emptyLogger struct{}

func (thisRef emptyLogger) Log(logEntry LogEntry) LogEntry {
	return logEntry
}
