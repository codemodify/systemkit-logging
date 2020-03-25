package mixers

import (
	housekeeping "github.com/codemodify/systemkit-logging/local-house-keeping"

	loggingC "github.com/codemodify/systemkit-logging/contracts"
)

type multiLogger struct {
	loggers []loggingC.Logger
}

// NewMultiLogger -
func NewMultiLogger(loggers []loggingC.Logger) loggingC.Logger {
	var thisRef = &multiLogger{
		loggers: loggers,
	}

	return housekeeping.NewDefaultHelperImplmentation(thisRef)
}

func (thisRef multiLogger) Log(logEntry loggingC.LogEntry) {
	for _, logger := range thisRef.loggers {
		logger.Log(logEntry)
	}
}
