package logging

import "os"

type stdoutLogger struct{}

func (thisRef stdoutLogger) Log(logEntry LogEntry) LogEntry {
	os.Stdout.WriteString(logEntry.Message + "\n")

	return logEntry
}
