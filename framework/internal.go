package framework

import (
	"github.com/lowl11/lazy-framework/controllers"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/events"
	"github.com/lowl11/lazy-framework/framework/echo_server"
	"github.com/lowl11/lazy-framework/log"
)

var (
	server interfaces.IServer
)

func initFramework() error {
	// log init
	log.Init(LogFileName, LogFolderName)

	// events init
	events.Init()

	// controllers init
	controllers.Init()

	// server init
	server = echo_server.Create(TimeoutDuration)
	return nil
}
