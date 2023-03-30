package echo_server

import (
	"github.com/labstack/echo/v4"
	"time"
)

type Server struct {
	server *echo.Echo

	useSwagger bool
}

func Create(timeout time.Duration) *Server {
	server := &Server{
		server: echo.New(),
	}

	server.setMiddlewares(timeout)
	server.setEndpoints()

	return server
}
