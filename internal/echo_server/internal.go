package echo_server

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/lowl11/owl/internal/controllers"
	"github.com/lowl11/owl/internal/middlewares/echo_middlewares"
)

func (server *Server) setMiddlewares() {
	server.app.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	server.app.Use(middleware.Secure())
	server.app.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	server.app.Use(echo_middlewares.Timeout(server.serverTimeout))
}

func (server *Server) setEndpoints() {
	server.app.GET("/health", controllers.Static.Health)
	server.app.RouteNotFound("*", controllers.Static.RouteNotFound)
}
