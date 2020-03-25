package tests

import (
	"testing"

	logging "github.com/codemodify/systemkit-logging"
	loggingC "github.com/codemodify/systemkit-logging/contracts"
	loggingF "github.com/codemodify/systemkit-logging/formatters"
	loggingM "github.com/codemodify/systemkit-logging/mixers"
	loggingP "github.com/codemodify/systemkit-logging/persisters"
)

func Test_03(t *testing.T) {
	logging.Init(
		loggingF.NewSimpleFormatterLogger(
			loggingM.NewAsyncLogger(
				loggingM.NewBufferedLogger(
					loggingM.NewMultiLogger(
						[]loggingC.Logger{
							loggingP.NewConsoleLogger(loggingC.TypeDebug, loggingP.NewConsoleLoggerDefaultColors()),
							loggingP.NewFileLogger(loggingC.TypeDebug, "log.log"),
						},
					),
					loggingM.BufferedLoggerConfig{
						MaxLogEntries: 100,
					},
				),
			),
		),
	)

	logging.Instance().LogInfo("Info line")
	logging.Instance().LogWarning("Warning line")
	logging.Instance().LogError("Error line")
}
