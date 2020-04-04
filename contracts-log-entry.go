package logging

import (
	"time"
)

// LogEntry -
type LogEntry struct {
	Time    time.Time // time.Now()
	Type    LogType   // TypeDisable .. -> .. TypeDebug
	Message string    //
}
