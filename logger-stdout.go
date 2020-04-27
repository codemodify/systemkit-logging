package logging

import "os"

type stdoutLogger struct{}

// NewStdoutLogger - logs to STDOUT
func NewStdoutLogger() CoreLogger {
	return &stdoutLogger{}
}

func (thisRef stdoutLogger) Log(logEntry LogEntry) LogEntry {
	os.Stdout.WriteString(logEntry.Message)

	return logEntry
}
