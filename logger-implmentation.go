package logging

import (
	"fmt"
	"time"
)

type loggerImplementation struct {
	logger   CoreLogger
	logUntil LogType
}

// NewLoggerImplementation -
func NewLoggerImplementation(logger CoreLogger) Logger {
	return &loggerImplementation{
		logger:   logger,
		logUntil: TypeDebug,
	}
}

func (thisRef *loggerImplementation) KeepOnlyLogs(logUntil LogType) Logger {
	thisRef.logUntil = logUntil
	return thisRef
}

func (thisRef loggerImplementation) Tracef(format string, v ...interface{}) {
	if canLog(TypeTrace, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeTrace,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Panicf(format string, v ...interface{}) {
	if canLog(TypePanic, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypePanic,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Fatalf(format string, v ...interface{}) {
	if canLog(TypeFatal, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeFatal,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Errorf(format string, v ...interface{}) {
	if canLog(TypeError, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeError,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Warningf(format string, v ...interface{}) {
	if canLog(TypeWarning, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeWarning,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Infof(format string, v ...interface{}) {
	if canLog(TypeInfo, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeInfo,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Successf(format string, v ...interface{}) {
	if canLog(TypeSuccess, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeSuccess,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Debugf(format string, v ...interface{}) {
	if canLog(TypeDebug, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeDebug,
			Message: fmt.Sprintf(format, v...),
		})
	}
}

func (thisRef loggerImplementation) Trace(v ...interface{}) {
	if canLog(TypeTrace, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeTrace,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef loggerImplementation) Panic(v ...interface{}) {
	if canLog(TypePanic, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypePanic,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef loggerImplementation) Fatal(v ...interface{}) {
	if canLog(TypeFatal, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeFatal,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef loggerImplementation) Error(v ...interface{}) {
	if canLog(TypeError, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeError,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef loggerImplementation) Warning(v ...interface{}) {
	if canLog(TypeWarning, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeWarning,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef loggerImplementation) Info(v ...interface{}) {
	if canLog(TypeInfo, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeInfo,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef loggerImplementation) Success(v ...interface{}) {
	if canLog(TypeSuccess, thisRef.logUntil) {
		thisRef.logger.Log(LogEntry{
			Time:    time.Now(),
			Type:    TypeSuccess,
			Message: fmt.Sprint(v...),
		})
	}
}

func (thisRef loggerImplementation) Debug(v ...interface{}) {
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
