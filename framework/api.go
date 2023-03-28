package framework

import (
	"github.com/lowl11/lazy-framework/log"
)

var (
	LogFileName   = "info" // example: info.log
	LogFolderName = "logs" // example: logs/info_28-03-2023.log
)

func Init() error {
	log.Init(LogFileName, LogFolderName)
	return nil
}

func SetLogConfig(fileName, folderName string) {
	if len(fileName) > 0 {
		LogFileName = fileName
	}

	if len(folderName) > 0 {
		LogFolderName = folderName
	}
}
