package echo_server

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/lowl11/lazy-framework/controllers"
	"github.com/lowl11/lazy-framework/middlewares/echo_middlewares"
	echoSwagger "github.com/swaggo/echo-swagger"
	"time"
)

func (server *Server) setMiddlewares(timeout time.Duration) {
	server.server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	server.server.Use(middleware.Secure())
	server.server.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	server.server.Use(echo_middlewares.Timeout(timeout))
}

func (server *Server) setEndpoints() {
	server.server.GET("/health", controllers.Static.Health)
	server.server.RouteNotFound("*", controllers.Static.RouteNotFound)

	if server.useSwagger {
		server.server.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	}
}
