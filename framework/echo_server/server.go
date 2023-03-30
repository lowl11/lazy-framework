package echo_server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lowl11/lazy-framework/middlewares/echo_middlewares"
	"time"
)

type Server struct {
	server *echo.Echo
}

func Create(timeout time.Duration) *Server {
	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	server.Use(middleware.Secure())
	server.Use(middleware.RecoverWithConfig(middleware.DefaultRecoverConfig))
	server.Use(echo_middlewares.Timeout(timeout))

	return &Server{
		server: server,
	}
}
