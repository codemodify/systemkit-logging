package logging

import (
	"os"
	"sync"
)

var instance LoggerImplementation
var instanceSync sync.RWMutex

// Instance -
func Instance() LoggerImplementation {
	instanceSync.Lock()
	defer instanceSync.Unlock()

	// dump to STDOUT if nothing was set
	if instance == nil {
		instanceSync.Unlock()
		SetLogger(NewDefaultLoggerImplementation(&stdoutLogger{}))
		instanceSync.Lock()
	}

	return instance
}

// SetLogger -
func SetLogger(logger LoggerImplementation) {
	instanceSync.Lock()
	defer instanceSync.Unlock()

	instance = logger
}

type stdoutLogger struct{}

func (thisRef stdoutLogger) Log(logEntry LogEntry) LogEntry {
	os.Stdout.WriteString(logEntry.Message)

	return logEntry
}
