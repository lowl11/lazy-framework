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

func WebFramework(webFramework string) {
	_webFramework = webFramework
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
	_serverMutex.Lock()
	defer _serverMutex.Unlock()

	if _server == nil {
		initFramework()
	}

	return _server
}
