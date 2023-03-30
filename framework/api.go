package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/framework/echo_server"
	"github.com/lowl11/lazy-framework/log"
	"time"
)

var (
	LogFileName   = "info" // example: info.log
	LogFolderName = "logs" // example: logs/info_28-03-2023.log

	TimeoutDuration = time.Second * 60
)

var (
	server interfaces.IServer
)

func Init() error {
	// log init
	log.Init(LogFileName, LogFolderName)

	// server init
	server = echo_server.Create(TimeoutDuration)

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

func SetServerTimeout(timeout time.Duration) {
	TimeoutDuration = timeout
}

func ServerEcho() *echo.Echo {
	return Server().(interfaces.IEchoServer).Get()
}

func Server() interfaces.IServer {
	if server == nil {
		panic("Server is NULL")
	}
	return server
}
