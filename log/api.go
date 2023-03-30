package log

import "github.com/lowl11/lazylog/logapi"

const (
	FileName   = "info"
	FolderName = "logs"
)

var logger logapi.ILogger

func Init(fileName, folderName string) {
	if logger != nil { // logger already exist
		return
	}

	// creating new logger instance
	logger = logapi.New().File(fileName, folderName)
}

func Info(args ...string) {
	if logger == nil {
		Init(FileName, FolderName)
	}
	logger.Info(args...)
}

func Debug(args ...string) {
	if logger == nil {
		Init(FileName, FolderName)
	}
	logger.Debug(args...)
}

func Warn(args ...string) {
	if logger == nil {
		Init(FileName, FolderName)
	}
	logger.Warn(args...)
}

func Error(err error, args ...string) {
	if logger == nil {
		Init(FileName, FolderName)
	}
	logger.Error(err, args...)
}

func Fatal(err error, args ...string) {
	if logger == nil {
		Init(FileName, FolderName)
	}
	logger.Fatal(err, args...)
}
