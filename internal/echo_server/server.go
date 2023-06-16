package echo_server

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/owl/data/domain"
	"time"
)

type Server struct {
	server        *echo.Echo
	serverTimeout time.Duration
	http2Config   *domain.Http2Config

	useHttp2 bool
}

func New(timeout time.Duration, useHttp2 bool) *Server {
	server := &Server{
		server:   echo.New(),
		useHttp2: useHttp2,
	}

	server.serverTimeout = timeout
	server.setMiddlewares()
	server.setEndpoints()

	return server
}
