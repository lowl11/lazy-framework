package framework

import (
	"github.com/lowl11/lazy-framework/controllers"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/events"
	"github.com/lowl11/lazy-framework/framework/echo_server"
	"github.com/lowl11/lazy-framework/log"
	"sync"
)

const (
	defaultWebFramework = "echo"
)

var (
	_server      interfaces.IServer
	_serverMutex sync.Mutex

	_webFramework = defaultWebFramework
)

func initFramework() {
	// log init
	log.Init()

	// events init
	events.Init()

	// controllers init
	controllers.Init()

	// server init
	switch _webFramework {
	case EchoFramework:
		_server = echo_server.Create(TimeoutDuration)
	}
	if _server == nil {
		panic("Initialization error. Server is NULL")
	}

	if useSwagger {
		_server.ActivateSwagger()
	}
}
