package logging

import "os"

type stderrLogger struct{}

func (thisRef stderrLogger) Log(logEntry LogEntry) LogEntry {
	os.Stderr.WriteString(logEntry.Message)

	return logEntry
}

// NewStderrLogger -
func NewStderrLogger() CoreLogger {
	return &stderrLogger{}
}
