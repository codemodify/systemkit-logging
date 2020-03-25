// +build windows

package logging

import (
	loggingC "github.com/codemodify/systemkit-logging/contracts"
	housekeeping "github.com/codemodify/systemkit-logging/local-house-keeping"
	loggingP "github.com/codemodify/systemkit-logging/persisters"
)

// NewWindowsEventLogger -
func NewWindowsEventLogger() loggingC.EasyLogger {
	return housekeeping.NewDefaultHelperImplmentation(
		loggingP.NewWindowsEventlogLogger(loggingC.TypeDebug),
	)
}
