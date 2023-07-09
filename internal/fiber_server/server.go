package fiber_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lowl11/owl/data/domain"
	"time"
)

type Server struct {
	app           *fiber.App
	serverTimeout time.Duration
	http2Config   domain.Http2Config

	useHttp2 bool
}

func New(timeout time.Duration, useHttp2 bool) *Server {
	server := &Server{
		app:           fiber.New(),
		useHttp2:      useHttp2,
		serverTimeout: timeout,
	}

	return server
}
