package log

import "github.com/lowl11/lazylog/logapi"

const (
	FileName   = "info"
	FolderName = "logs"
)

var (
	_fileName   = FileName
	_folderName = FolderName
)

var logger logapi.ILogger

func Init() {
	if logger != nil { // logger already exist
		return
	}

	// creating new logger instance
	logger = logapi.New().File(_fileName, _folderName)
}

func SetConfig(fileName, folderName string) {
	if len(fileName) > 0 {
		_fileName = fileName
	}

	if len(folderName) > 0 {
		_folderName = folderName
	}
}

func Info(args ...string) {
	if logger == nil {
		Init()
	}
	logger.Info(args...)
}

func Debug(args ...string) {
	if logger == nil {
		Init()
	}
	logger.Debug(args...)
}

func Warn(args ...string) {
	if logger == nil {
		Init()
	}
	logger.Warn(args...)
}

func Error(err error, args ...string) {
	if logger == nil {
		Init()
	}
	logger.Error(err, args...)
}

func Fatal(err error, args ...string) {
	if logger == nil {
		Init()
	}
	logger.Fatal(err, args...)
}
