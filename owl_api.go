package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazylog/log"
	"github.com/lowl11/owl/data/interfaces"
)

func (lazy *Owl) ShutdownAction(action func()) {
	lazy.shutdownService.Add(action)
}

func (lazy *Owl) Start(port string) {
	lazy.server.Start(port)
}

func (lazy *Owl) StartHttp2(port string) {
	lazy.server.StartHttp2(port)
}

func (lazy *Owl) StartGrpc(port string) {
	lazy.ShutdownAction(func() {
		if err := lazy.grpcServer.Close(); err != nil {
			log.Error(err, "Close gRPC server connection error")
			return
		}
		log.Info("gRPC server connection closed!")
	})

	lazy.getGrpcServer().Start(port)
}

func (lazy *Owl) Echo() *echo.Echo {
	return lazy.server.(interfaces.IEchoServer).Get()
}
