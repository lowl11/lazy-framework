package framework

import (
	"fmt"
	"github.com/lowl11/lazy-framework/config"
	"github.com/lowl11/lazy-framework/controllers"
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/events"
	"github.com/lowl11/lazy-framework/framework/echo_server"
	"github.com/lowl11/lazy-framework/log"
	"sync"
)

const (
	defaultWebFramework = "echo"

	http2MaxConcurrentStreams = 250
	http2MaxReadFrameSize     = 1048576
)

var (
	_server      interfaces.IServer
	_serverMutex sync.Mutex

	_webFramework = defaultWebFramework
)

func initFramework(frameworkConfig *Config) {
	defer func() {
		_initDone = true
	}()

	// log init
	initLog(frameworkConfig)
	log.Init()

	// config init
	initConfig(frameworkConfig)
	config.Init()

	// events init
	events.Init()

	// controllers init
	controllers.Init()

	// server init
	timeoutDuration := _timeoutDuration
	if frameworkConfig.ServerTimeout != 0 {
		timeoutDuration = frameworkConfig.ServerTimeout
	}

	fmt.Println("timeout duration:", timeoutDuration.String())

	switch _webFramework {
	case EchoFramework:
		_server = echo_server.New(timeoutDuration, frameworkConfig.UseHttp2)
	}
	if _server == nil {
		panic("Initialization error. Server is NULL")
	}

	// set http 2.0 server
	if frameworkConfig.UseHttp2 {
		// if config is empty, use default values
		if _http2Config == nil {
			_http2Config = &domain.Http2Config{
				MaxConcurrentStreams: http2MaxConcurrentStreams,
				MaxReadFrameSize:     http2MaxReadFrameSize,
			}
		}

		// set http 2.0 server config
		_server.SetHttp2Config(_http2Config)
	}

	if frameworkConfig.UseSwagger {
		_server.ActivateSwagger()
	}
}

func initLog(config *Config) {
	// file logger
	log.SetConfig(config.LogFileName, config.LogFolderName)

	// custom loggers
	if config.CustomLoggers != nil {
		log.SetCustom(config.CustomLoggers...)
	}

	// modes
	if config.LogNoTime {
		log.SetNoTimeMode()
	}

	if config.LogJson {
		log.SetJsonMode()
	}

	if config.LogNoPrefix {
		log.SetNoPrefixMode()
	}

	if config.LogNoFile {
		log.SetNoFileMode()
	}
}

func initConfig(frameworkConfig *Config) {
	config.SetEnvironmentName(frameworkConfig.EnvironmentName)
	config.SetEnvironmentDefault(frameworkConfig.EnvironmentDefault)
	config.SetEnvironmentFileName(frameworkConfig.EnvironmentFileName)
}
