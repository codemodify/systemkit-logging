package logging

import (
	"fmt"
)

// LogType - Log type
type LogType int

// LogType -
const (
	TypeDisable LogType = iota // 0
	TypeTrace                  // 1 - log this no matter what
	TypePanic                  // 2
	TypeFatal                  // 3
	TypeError                  // 4
	TypeWarning                // 5
	TypeInfo                   // 6
	TypeSuccess                // 7
	TypeDebug                  // 8
)

// String - gives string representation of the log type
func (logType LogType) String() string {
	switch logType {
	case TypeDisable:
		return "DISABLED" // this is never logged

	// all values below are same STRING-length to improve logs reading
	case TypeTrace:
		return "TRACE"
	case TypePanic:
		return "PANIC"
	case TypeFatal:
		return "FATAL"
	case TypeError:
		return "ERROR"
	case TypeWarning:
		return "WAARN"
	case TypeInfo:
		return "INFOO"
	case TypeSuccess:
		return "SUCES"
	case TypeDebug:
		return "DEBUG"

	default:
		return fmt.Sprintf("%d", int(logType))
	}
}
