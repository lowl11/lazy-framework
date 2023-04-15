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
	_initDone       bool
	_useSwagger     bool
	_useHttp2       bool

	_http2Config *models.Http2Config
)

func Init() {
	if _initDone {
		return
	}

	initFramework()
}

func UseSwagger() {
	warnInit()
	_useSwagger = true
}

func UseHttp2(config *models.Http2Config) {
	warnInit()
	_useHttp2 = true
	_http2Config = config
}

func WebFramework(webFramework string) {
	warnInit()
	_webFramework = webFramework
}

func SetLogConfig(fileName, folderName string) {
	warnInit()
	log.SetConfig(fileName, folderName)
}

func SetLogJsonMode() {
	warnInit()
	log.SetJsonMode()
}

func SetLogNoTimeMode() {
	warnInit()
	log.SetNoTimeMode()
}

func SetCustomLoggers(customLoggers ...logapi.ILogger) {
	warnInit()
	log.SetCustom(customLoggers...)
}

func SetEnvironmentName(name string) {
	warnInit()
	config.SetEnvironmentName(name)
}

func SetEnvironmentDefault(name string) {
	warnInit()
	config.SetEnvironmentDefault(name)
}

func SetEnvironmentFileName(fileName string) {
	warnInit()
	config.SetEnvironmentFileName(fileName)
}

func SetServerTimeout(timeout time.Duration) {
	warnInit()
	TimeoutDuration = timeout
}

func StartServer(port string) {
	Server().Start(port)
}

func ServerEcho() *echo.Echo {
	return Server().(interfaces.IEchoServer).Get()
}

func Server() interfaces.IServer {
	_serverMutex.Lock()
	defer _serverMutex.Unlock()

	if !_initDone {
		panic("Framework initialization was not done!")
	}

	return _server
}
