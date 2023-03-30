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

var _logger logapi.ILogger

func Init() {
	if _logger != nil { // _logger already exist
		return
	}

	// creating new _logger instance
	_logger = logapi.New().File(_fileName, _folderName)
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
	if _logger == nil {
		Init()
	}
	_logger.Info(args...)
}

func Debug(args ...string) {
	if _logger == nil {
		Init()
	}
	_logger.Debug(args...)
}

func Warn(args ...string) {
	if _logger == nil {
		Init()
	}
	_logger.Warn(args...)
}

func Error(err error, args ...string) {
	if _logger == nil {
		Init()
	}
	_logger.Error(err, args...)
}

func Fatal(err error, args ...string) {
	if _logger == nil {
		Init()
	}
	_logger.Fatal(err, args...)
}
