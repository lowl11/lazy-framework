package log

import "github.com/lowl11/lazylog/logapi"

var logger logapi.ILogger

func Init(fileName, folderName string) {
	logger = logapi.New().File(fileName, folderName)
}

func Info(args ...string) {
	logger.Info(args...)
}

func Debug(args ...string) {
	logger.Debug(args...)
}

func Warn(args ...string) {
	logger.Warn(args...)
}

func Error(err error, args ...string) {
	logger.Error(err, args...)
}

func Fatal(err error, args ...string) {
	logger.Fatal(err, args...)
}
