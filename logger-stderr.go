package logging

import "os"

type stderrLogger struct{}

// NewStderrLogger - logs to STDERR
func NewStderrLogger() CoreLogger {
	return &stderrLogger{}
}

func (thisRef stderrLogger) Log(logEntry LogEntry) LogEntry {
	os.Stderr.WriteString(logEntry.Message)

	return logEntry
}
