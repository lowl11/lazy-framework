package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"time"
)

var (
	LogFileName   = "info" // example: info.log
	LogFolderName = "logs" // example: logs/info_28-03-2023.log

	TimeoutDuration = time.Second * 60
)

func SetLogConfig(fileName, folderName string) {
	if len(fileName) > 0 {
		LogFileName = fileName
	}

	if len(folderName) > 0 {
		LogFolderName = folderName
	}
}

func SetServerTimeout(timeout time.Duration) {
	TimeoutDuration = timeout
}

func ServerEcho() *echo.Echo {
	return Server().(interfaces.IEchoServer).Get()
}

func Server() interfaces.IServer {
	if err := initFramework(); err != nil {
		panic("Initialization lazy-framework error: " + err.Error())
	}

	if server == nil {
		panic("Initialization error. Server is NULL")
	}

	return server
}
