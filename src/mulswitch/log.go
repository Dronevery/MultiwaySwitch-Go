package mulswitch

import (
	"code.google.com/p/log4go"
	"github.com/Unknwon/goconfig"
)

var logger log4go.Logger

func initLogger() {
	logger = make(log4go.Logger)
	logger.AddFilter("stdout", log4go.FINEST, log4go.NewConsoleLogWriter())

	if logFileName := configCommon("role", INFO); logFileName != "" {
		logger.AddFilter("file", log4go.INFO, log4go.NewFileLogWriter(logFileName, false).SetFormat("[%D %T][%L]%M - %S"))
	}
	logger.Info("System start")
}
