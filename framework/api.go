package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/log"
	"time"
)

var (
	TimeoutDuration = time.Second * 60
	useSwagger      bool
)

func UseSwagger() {
	useSwagger = true
}

func SetLogConfig(fileName, folderName string) {
	log.SetConfig(fileName, folderName)
}

func SetServerTimeout(timeout time.Duration) {
	TimeoutDuration = timeout
}

func ServerEcho() *echo.Echo {
	return Server().(interfaces.IEchoServer).Get()
}

func Server() interfaces.IServer {
	if server == nil {
		initFramework()
	}

	if server == nil {
		panic("Initialization error. Server is NULL")
	}

	return server
}
