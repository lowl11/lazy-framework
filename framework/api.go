package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/config"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/data/models"
	"github.com/lowl11/lazy-framework/log"
	"github.com/lowl11/lazylog/logapi"
	"time"
)

var (
	TimeoutDuration = time.Second * 60
	_useSwagger     bool
	_useHttp2       bool

	_http2Config *models.Http2Config
)

func UseSwagger() {
	_useSwagger = true
}

func UseHttp2(config *models.Http2Config) {
	_useHttp2 = true
	_http2Config = config
}

func WebFramework(webFramework string) {
	_webFramework = webFramework
}

func SetLogConfig(fileName, folderName string) {
	log.SetConfig(fileName, folderName)
}

func SetCustomLoggers(customLoggers ...logapi.ILogger) {
	log.SetCustom(customLoggers...)
}

func SetEnvironmentName(name string) {
	config.SetEnvironmentName(name)
}

func SetEnvironmentDefault(name string) {
	config.SetEnvironmentDefault(name)
}

func SetEnvironmentFileName(fileName string) {
	config.SetEnvironmentFileName(fileName)
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
