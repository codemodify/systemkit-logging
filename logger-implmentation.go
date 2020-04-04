package logging

import (
	"fmt"
	"time"
)

type defaultLoggerImplementation struct {
	logger   Logger
	logUntil LogType
}

// NewDefaultLoggerImplementation -
func NewDefaultLoggerImplementation(logger Logger) LoggerImplementation {
	return &defaultLoggerImplementation{
		logger:   logger,
		logUntil: TypeDebug,
	}
}

func (thisRef *defaultLoggerImplementation) KeepOnlyLogs(logUntil LogType) {
	thisRef.logUntil = logUntil
}

func (thisRef defaultLoggerImplementation) Tracef(format string, v ...interface{}) {
	if canLog(TypeTrace, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeTrace,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Panicf(format string, v ...interface{}) {
	if canLog(TypePanic, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypePanic,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Fatalf(format string, v ...interface{}) {
	if canLog(TypeFatal, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeFatal,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Errorf(format string, v ...interface{}) {
	if canLog(TypeError, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeError,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Warningf(format string, v ...interface{}) {
	if canLog(TypeWarning, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeWarning,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Infof(format string, v ...interface{}) {
	if canLog(TypeInfo, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeInfo,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Successf(format string, v ...interface{}) {
	if canLog(TypeSuccess, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeSuccess,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Debugf(format string, v ...interface{}) {
	if canLog(TypeDebug, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeDebug,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Trace(v ...interface{}) {
	if canLog(TypeTrace, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeTrace,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Panic(v ...interface{}) {
	if canLog(TypePanic, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypePanic,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Fatal(v ...interface{}) {
	if canLog(TypeFatal, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeFatal,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Error(v ...interface{}) {
	if canLog(TypeError, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeError,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Warning(v ...interface{}) {
	if canLog(TypeWarning, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeWarning,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Info(v ...interface{}) {
	if canLog(TypeInfo, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeInfo,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Success(v ...interface{}) {
	if canLog(TypeSuccess, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeSuccess,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef defaultLoggerImplementation) Debug(v ...interface{}) {
	if canLog(TypeDebug, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeDebug,
			Message: fmt.Sprint(v...),
		})
	}
}

func canLog(logType LogType, logUntil LogType) bool {
	if logType == TypeDisable {
		return false
	}

	if logType == TypeTrace {
		return true
	}

	if logType <= logUntil {
		return true
	}

	return false
}
