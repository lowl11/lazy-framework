package echo_server

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/lowl11/lazy-framework/controllers"
	"github.com/lowl11/lazy-framework/framework/middlewares/echo_middlewares"
)

func (server *Server) setMiddlewares() {
	server.server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	server.server.Use(middleware.Secure())
	server.server.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	server.server.Use(echo_middlewares.Timeout(server.serverTimeout))
}

func (server *Server) setEndpoints() {
	server.server.GET("/health", controllers.Static.Health)
	server.server.RouteNotFound("*", controllers.Static.RouteNotFound)
}
