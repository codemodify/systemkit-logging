package logging

import (
	"sync"
)

var instance LoggerImplementation
var instanceMutex sync.RWMutex

// Instance -
func Instance() LoggerImplementation {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	// dump to STDOUT if nothing was set
	if instance == nil {
		instanceMutex.Unlock()
		SetLogger(&stdoutLogger{})
		instanceMutex.Lock()
	}

	return instance
}

// SetLogger -
func SetLogger(logger Logger) Logger {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	instance = NewDefaultLoggerImplementation(logger)
	return logger
}
