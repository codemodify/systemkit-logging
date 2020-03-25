package logging

import (
	"sync"

	loggingP "github.com/codemodify/systemkit-logging/persisters"
	loggingC "github.com/codemodify/systemkit-logging/contracts"
	housekeeping "github.com/codemodify/systemkit-logging/local-house-keeping"
)

var instance loggingC.EasyLogger
var instanceSync sync.RWMutex

// Instance -
func Instance() loggingC.EasyLogger {
	instanceSync.Lock()
	defer instanceSync.Unlock()

	if instance == nil {
		instanceSync.Unlock()
		Init(NewConsoleLogger())
		instanceSync.Lock()
	}

	return instance
}

// Init -
func Init(logger loggingC.EasyLogger) {
	instanceSync.Lock()
	defer instanceSync.Unlock()

	instance = logger
}

// NewConsoleLogger -
func NewConsoleLogger() loggingC.EasyLogger {
	return NewConsoleLoggerCustomColors(loggingP.NewConsoleLoggerDefaultColors())
}

// NewConsoleLoggerCustomColors -
func NewConsoleLoggerCustomColors(colors map[loggingC.LogType]loggingP.ConsoleLoggerColorDelegate) loggingC.EasyLogger {
	return housekeeping.NewDefaultHelperImplmentation(
		loggingP.NewConsoleLogger(loggingC.TypeDebug, colors),
	)
}

// NewFileLogger -
func NewFileLogger() loggingC.EasyLogger {
	return housekeeping.NewDefaultHelperImplmentation(
		loggingP.NewFileLoggerDefaultName(loggingC.TypeDebug),
	)
}

// NewEasyLoggerForLogger -
func NewEasyLoggerForLogger(logger loggingC.Logger) loggingC.EasyLogger {
	return housekeeping.NewDefaultHelperImplmentation(logger)
}
