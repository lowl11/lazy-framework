package echo_server

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazylog/log"
	"github.com/lowl11/owl/data/domain"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/net/http2"
)

func (server *Server) Start(port string) {
	// start server
	log.Fatal(server.app.Start(port), "Start server error")
}

func (server *Server) StartHttp2(port string) {
	// check is HTTP 2.0 mode is ON
	if !server.useHttp2 {
		log.Fatal(errors.New("http 2.0 mode is turned off"), "Starting HTTP 2.0 app")
	}

	// check if config is NULL
	if server.http2Config.MaxReadFrameSize == 0 {
		log.Fatal(errors.New("config is NULL"), "Starting HTTP 2.0 app")
	}

	// initializing app config
	http2Server := &http2.Server{
		MaxConcurrentStreams: uint32(server.http2Config.MaxConcurrentStreams),
		MaxReadFrameSize:     uint32(server.http2Config.MaxReadFrameSize),
		IdleTimeout:          server.serverTimeout,
	}

	// start app
	log.Fatal(server.app.StartH2CServer(port, http2Server), "Start HTTP2 server error")
}

func (server *Server) SetHttp2Config(config domain.Http2Config) {
	// check if config is NULL
	if config.MaxReadFrameSize == 0 {
		log.Fatal(errors.New("config is NULL"), "Setting HTTP 2.0 config")
	}

	server.http2Config = config
}

func (server *Server) ActivateSwagger(customEndpoint ...string) {
	endpoint := "/swagger/*"
	if len(customEndpoint) > 0 {
		endpoint = customEndpoint[0]
	}

	// activating swagger endpoints
	server.app.GET(endpoint, echoSwagger.EchoWrapHandler())
}

func (server *Server) Get() *echo.Echo {
	return server.app
}
