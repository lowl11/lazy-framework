package framework

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/log"
	"github.com/lowl11/lazy-rmq/rabbit_event"
	"github.com/lowl11/lazylog/logapi"
	"time"
)

type Config struct {
	UseSwagger bool

	// log
	LogFileName   string
	LogFolderName string
	LogJson       bool
	LogNoTime     bool
	LogNoPrefix   bool
	LogNoFile     bool
	LogLevel      uint
	CustomLoggers []logapi.ILogger

	// environment
	EnvironmentName     string
	EnvironmentDefault  string
	EnvironmentFileName string

	// server
	UseHttp2      bool
	Http2Config   *domain.Http2Config
	WebFramework  string
	ServerTimeout time.Duration

	// grpc
	UseGRPC  bool
	OnlyGRPC bool
	LogGRPC  bool
}

func Init(config *Config) {
	if _initDone {
		return
	}

	initFramework(config)
}

func StartServer(port string) {
	Server().Start(port)
}

func StartHttp2Server(port string) {
	Server().StartHttp2(port)
}

func StartGRPC(port string) {
	ShutDownAction(func() {
		if err := GrpcServer().Close(); err != nil {
			log.Error(err, "Close gRPC server connection error")
			return
		}
		log.Info("gRPC server connection closed!")
	})

	GrpcServer().Start(port)
}

func ServerEcho() *echo.Echo {
	return Server().(interfaces.IEchoServer).Get()
}

func Server() interfaces.IServer {
	_serverMutex.Lock()
	defer _serverMutex.Unlock()

	if !_initDone {
		panic("Framework initialization was not done!")
	}

	return _server
}

func GrpcServer() interfaces.IGRPCServer {
	_grpcServerMutex.Lock()
	defer _grpcServerMutex.Unlock()

	if !_initDone {
		panic("Framework initialization was not done!")
	}

	if !_useGrpc {
		panic("Set the flag \"UseGRPC\" in framework config")
	}

	return _grpcServer
}

func Rabbit(connectionString string, heartbeat time.Duration) (*rabbit_event.Event, error) {
	rmqEvent, err := rabbit_event.NewConfig(connectionString, heartbeat)
	if err != nil {
		return nil, err
	}

	return rmqEvent, nil
}

func ShutDownAction(action func()) {
	addShutDownAction(action)
}

func IsGrpc() bool {
	return _useGrpc
}
