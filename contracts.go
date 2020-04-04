package logging

// Logger -
type Logger interface {
	Log(logEntry LogEntry) LogEntry
}

// LoggerImplementation -
type LoggerImplementation interface {
	KeepOnlyLogs(logUntil LogType)

	Tracef(format string, v ...interface{})
	Panicf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Successf(format string, v ...interface{})
	Debugf(format string, v ...interface{})

	Trace(v ...interface{})
	Panic(v ...interface{})
	Fatal(v ...interface{})
	Error(v ...interface{})
	Warning(v ...interface{})
	Info(v ...interface{})
	Success(v ...interface{})
	Debug(v ...interface{})
}
