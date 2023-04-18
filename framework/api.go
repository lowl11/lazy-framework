package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazylog/logapi"
	"time"
)

var (
	_timeoutDuration = time.Second * 60
	_initDone        bool

	_http2Config *domain.Http2Config
)

type Config struct {
	UseSwagger bool

	// log
	LogFileName   string
	LogFolderName string
	LogJson       bool
	LogNoTime     bool
	LogNoPrefix   bool
	LogNoFile     bool
	CustomLoggers []logapi.ILogger

	// environment
	EnvironmentName     string
	EnvironmentDefault  string
	EnvironmentFileName string

	// server
	UseHttp2      bool
	Http2Config   *domain.Http2Config
	WebFramework  string
	ServerTimeout time.Duration
}

func Init(config *Config) {
	if _initDone {
		return
	}

	initFramework(config)
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
