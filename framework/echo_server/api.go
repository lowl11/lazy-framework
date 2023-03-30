package echo_server

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/log"
)

func (server *Server) Start(port string) {
	log.Fatal(server.server.Start(port), "Start server error")
}

func (server *Server) ActivateSwagger() {
	//
}

func (server *Server) Get() *echo.Echo {
	return server.server
}
