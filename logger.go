package logging

var instance Logger

func init() {
	SetLogger(NewEmptyLogger())
}

// SetLogger -
func SetLogger(logger CoreLogger) Logger {
	instance = NewLoggerImplementation(logger)
	return instance
}

// Add shortcuts + duplicate `Logger`
func KeepOnlyLogs(logUntil LogType) Logger {
	instance.KeepOnlyLogs(logUntil)
	return instance
}

// Add shortcuts with formats + duplicate `Logger`
func Tracef(format string, v ...interface{})   { instance.Tracef(format, v...) }
func Panicf(format string, v ...interface{})   { instance.Panicf(format, v...) }
func Fatalf(format string, v ...interface{})   { instance.Fatalf(format, v...) }
func Errorf(format string, v ...interface{})   { instance.Errorf(format, v...) }
func Warningf(format string, v ...interface{}) { instance.Warningf(format, v...) }
func Infof(format string, v ...interface{})    { instance.Infof(format, v...) }
func Successf(format string, v ...interface{}) { instance.Successf(format, v...) }
func Debugf(format string, v ...interface{})   { instance.Debugf(format, v...) }

// Add shortcuts + duplicate `Logger`
func Trace(v ...interface{})   { instance.Trace(v...) }
func Panic(v ...interface{})   { instance.Panic(v...) }
func Fatal(v ...interface{})   { instance.Fatal(v...) }
func Error(v ...interface{})   { instance.Error(v...) }
func Warning(v ...interface{}) { instance.Warning(v...) }
func Info(v ...interface{})    { instance.Info(v...) }
func Success(v ...interface{}) { instance.Success(v...) }
func Debug(v ...interface{})   { instance.Debug(v...) }
