package log_internal

import "github.com/lowl11/lazylog/logapi"

func Init() {
	if _logger != nil { // _logger already exist
		return
	}

	// creating new _logger instance
	loggerInstance := logapi.New()
	for _, customLogger := range _customLoggers {
		loggerInstance.Custom(customLogger)
	}

	if !_noFileMode {
		loggerInstance.File(_fileName, _folderName)
	}

	if _noTimeMode {
		loggerInstance.NoTime()
	}

	if _jsonMode {
		loggerInstance.JSON()
	}

	if _noPrefixMode {
		loggerInstance.NoPrefix()
	}

	_logger = loggerInstance
}

func SetCustom(customLoggers ...logapi.ILogger) {
	_customLoggers = customLoggers
}

func SetConfig(fileName, folderName string) {
	if len(fileName) > 0 {
		_fileName = fileName
	}

	if len(folderName) > 0 {
		_folderName = folderName
	}
}

func SetJsonMode() {
	_jsonMode = true
}

func SetNoTimeMode() {
	_noTimeMode = true
}

func SetNoPrefixMode() {
	_noPrefixMode = true
}

func SetNoFileMode() {
	_noFileMode = true
}

func Info(args ...any) {
	if _logger == nil {
		Init()
	}
	_logger.Info(args...)
}

func Debug(args ...any) {
	if _logger == nil {
		Init()
	}
	_logger.Debug(args...)
}

func Warn(args ...any) {
	if _logger == nil {
		Init()
	}
	_logger.Warn(args...)
}

func Error(err error, args ...any) {
	if _logger == nil {
		Init()
	}
	_logger.Error(err, args...)
}

func Fatal(err error, args ...any) {
	if _logger == nil {
		Init()
	}
	_logger.Fatal(err, args...)
}
