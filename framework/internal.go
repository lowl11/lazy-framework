package framework

import (
	"github.com/lowl11/lazy-framework/config"
	"github.com/lowl11/lazy-framework/controllers"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/data/models"
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

func initFramework() {
	defer func() {
		_initDone = true
	}()

	// log init
	log.Init()

	// config init
	config.Init()

	// events init
	events.Init()

	// controllers init
	controllers.Init()

	// server init
	switch _webFramework {
	case EchoFramework:
		_server = echo_server.Create(TimeoutDuration, _useHttp2)
	}
	if _server == nil {
		panic("Initialization error. Server is NULL")
	}

	// set http 2.0 server
	if _useHttp2 {
		// if config is empty, use default values
		if _http2Config == nil {
			_http2Config = &models.Http2Config{
				MaxConcurrentStreams: http2MaxConcurrentStreams,
				MaxReadFrameSize:     http2MaxReadFrameSize,
			}
		}

		// set http 2.0 server config
		_server.SetHttp2Config(_http2Config)
	}

	if _useSwagger {
		_server.ActivateSwagger()
	}
}

func warnInit() {
	if _initDone {
		panic("Framework initialization already was done, move setting above the initialization")
	}
}
