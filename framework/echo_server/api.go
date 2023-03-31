package echo_server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/models"
	"github.com/lowl11/lazy-framework/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/net/http2"
)

func (server *Server) Start(port string) {
	// start server
	log.Fatal(server.server.Start(port), "Start server error")
}

func (server *Server) StartHttp2(port string) {
	// check is HTTP 2.0 mode is ON
	if !server.useHttp2 {
		log.Fatal(errors.New("http 2.0 mode is turned off"), "Starting HTTP 2.0 server")
	}

	// check if config is NULL
	if server.http2Config == nil {
		log.Fatal(errors.New("config is NULL"), "Starting HTTP 2.0 server")
	}

	// initializing server config
	http2Server := &http2.Server{
		MaxConcurrentStreams: uint32(server.http2Config.MaxConcurrentStreams),
		MaxReadFrameSize:     uint32(server.http2Config.MaxReadFrameSize),
		IdleTimeout:          server.serverTimeout,
	}

	// start server
	log.Fatal(server.server.StartH2CServer(port, http2Server), "Start HTTP2 server error")
}

func (server *Server) SetHttp2Config(config *models.Http2Config) {
	// check if config is NULL
	if config == nil {
		log.Fatal(errors.New("config is NULL"), "Setting HTTP 2.0 config")
	}

	server.http2Config = config
}

func (server *Server) ActivateSwagger() {
	// activating swagger endpoints
	server.server.GET("/swagger/*", echoSwagger.EchoWrapHandler())
}

func (server *Server) Get() *echo.Echo {
	return server.server
}
