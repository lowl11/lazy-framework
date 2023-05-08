package framework

import (
	"errors"
	"github.com/lowl11/lazy-collection/safe_array"
	"github.com/lowl11/lazy-framework/config"
	frameworkConfig "github.com/lowl11/lazy-framework/config"
	"github.com/lowl11/lazy-framework/controllers"
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/events"
	"github.com/lowl11/lazy-framework/framework/echo_server"
	"github.com/lowl11/lazy-framework/log"
	"github.com/lowl11/lazy-framework/log/log_internal"
	"github.com/lowl11/lazylog/logapi/log_levels"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	defaultWebFramework = "echo"

	http2MaxConcurrentStreams = 250
	http2MaxReadFrameSize     = 1048576
)

var (
	_initDone bool

	_server      interfaces.IServer
	_serverMutex sync.Mutex

	_shutDownActions *safe_array.Array[func()]
)

func init() {
	_shutDownActions = safe_array.New[func()]()
}

func initFramework(frameworkConfig *Config) {
	defer func() {
		_initDone = true
	}()

	// log init
	initLog(frameworkConfig)
	log_internal.Init()

	// config init
	initConfig(frameworkConfig)
	config.Init()

	// events init
	events.Init()

	// controllers init
	controllers.Init()

	// server init
	timeoutDuration := time.Second * 60
	if frameworkConfig.ServerTimeout != 0 {
		timeoutDuration = frameworkConfig.ServerTimeout
	}

	if frameworkConfig.WebFramework == "" {
		frameworkConfig.WebFramework = defaultWebFramework
	}
	switch frameworkConfig.WebFramework {
	case EchoFramework:
		_server = echo_server.New(timeoutDuration, frameworkConfig.UseHttp2)
	}
	if _server == nil {
		panic("Initialization error. Server is NULL")
	}

	// set http 2.0 server
	if frameworkConfig.UseHttp2 {
		// if config is empty, use default values
		if frameworkConfig.Http2Config == nil {
			frameworkConfig.Http2Config = &domain.Http2Config{
				MaxConcurrentStreams: http2MaxConcurrentStreams,
				MaxReadFrameSize:     http2MaxReadFrameSize,
			}
		}

		// set http 2.0 server config
		_server.SetHttp2Config(frameworkConfig.Http2Config)
	}

	if frameworkConfig.UseSwagger {
		_server.ActivateSwagger()
	}

	runShutDownWaiter()
}

func initLog(config *Config) {
	// file logger
	log_internal.SetConfig(config.LogFileName, config.LogFolderName)

	// custom loggers
	if config.CustomLoggers != nil {
		log_internal.SetCustom(config.CustomLoggers...)
	}

	// modes
	if config.LogNoTime {
		log_internal.SetNoTimeMode()
	}

	if config.LogJson {
		log_internal.SetJsonMode()
	}

	if config.LogNoPrefix {
		log_internal.SetNoPrefixMode()
	}

	if config.LogNoFile {
		log_internal.SetNoFileMode()
	}

	if config.LogLevel > 0 {
		log_internal.SetLogLevel(config.LogLevel)
	} else {
		if frameworkConfig.IsProduction() {
			log_internal.SetLogLevel(log_levels.INFO)
		}
	}
}

func initConfig(frameworkConfig *Config) {
	config.SetEnvironmentName(frameworkConfig.EnvironmentName)
	config.SetEnvironmentDefault(frameworkConfig.EnvironmentDefault)
	config.SetEnvironmentFileName(frameworkConfig.EnvironmentFileName)
}

func addShutDownAction(action func()) {
	_shutDownActions.Push(action)
}

func runShutDownWaiter() {
	go func() {
		// Create a channel to receive signals
		signalChannel := make(chan os.Signal, 1)

		// Notify the signal channel when a SIGINT or SIGTERM signal is received
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		<-signalChannel

		_shutDownActions.Each(func(item func()) {
			defer func() {
				if value := recover(); value != nil {
					var err error
					if _, ok := value.(string); ok {
						err = errors.New(value.(string))
					} else if _, ok = value.(error); ok {
						err = value.(error)
					}
					log.Error(err, "Catch panic from shut down action")
				}
			}()

			// call action
			item()
		})

		// call shutdown
		os.Exit(0)
	}()
}
