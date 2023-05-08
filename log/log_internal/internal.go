package log_internal

import "github.com/lowl11/lazylog/logapi"

const (
	defaultFileName   = "info"
	defaultFolderName = "logs"
)

var (
	_logger        logapi.ILogger
	_customLoggers = make([]logapi.ILogger, 0)
)

var (
	_fileName   = defaultFileName
	_folderName = defaultFolderName

	_jsonMode     bool
	_noTimeMode   bool
	_noPrefixMode bool
	_noFileMode   bool
)
