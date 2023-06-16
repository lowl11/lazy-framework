package framework

import (
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/helpers/error_helper"
	"github.com/lowl11/lazy-framework/internal/controllers"
	"github.com/lowl11/lazy-framework/internal/echo_server"
	"github.com/lowl11/lazy-framework/internal/grpc_server"
	"github.com/lowl11/lazy-framework/internal/shutdown_service"
	"github.com/lowl11/lazyconfig/config"
	frameworkConfig "github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazylog/log/log_internal"
	"github.com/lowl11/lazylog/logapi/log_levels"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	defaultWebFramework = "echo"

	http2MaxConcurrentStreams = 250
	http2MaxReadFrameSize     = 1048576
)

var (
	_initDone bool

	_server      interfaces.IServer
	_serverMutex sync.Mutex

	_useGrpc         bool
	_grpcServer      interfaces.IGRPCServer
	_grpcServerMutex sync.Mutex

	_shutdownService *shutdown_service.Service
)

func initFramework(frameworkConfig *Config) {
	defer func() {
		_initDone = true
	}()

	// config init
	initConfig(frameworkConfig)

	// log init
	initLog(frameworkConfig)

	// controllers init
	controllers.Init()

	// server init
	initServer(frameworkConfig)

	// gRPC server init
	initGrpcServer(frameworkConfig)

	_shutdownService = shutdown_service.New()
	go runShutDownWaiter()
}

func initLog(config *Config) {
	logLevel := config.LogLevel
	if config.LogLevel == 0 && frameworkConfig.IsProduction() {
		logLevel = log_levels.INFO
	}

	log_internal.Init(log_internal.LogConfig{
		FileName:      config.LogFileName,
		FolderName:    config.LogFolderName,
		NoFile:        config.LogNoFile,
		NoTime:        config.LogNoTime,
		NoPrefix:      config.LogNoPrefix,
		JsonMode:      config.LogJson,
		LogLevel:      logLevel,
		CustomLoggers: config.CustomLoggers,
	})
}

func initConfig(frameworkConfig *Config) {
	config.SetEnvironmentName(frameworkConfig.EnvironmentName)
	config.SetEnvironmentDefault(frameworkConfig.EnvironmentDefault)
	config.SetEnvironmentFileName(frameworkConfig.EnvironmentFileName)
	config.Init()
}

func initServer(frameworkConfig *Config) {
	// HTTP server already exist
	if _server != nil || _initDone {
		return
	}

	// only gRPC server (no HTTP)
	if frameworkConfig.UseGRPC && frameworkConfig.OnlyGRPC {
		return
	}

	_serverMutex.Lock()
	defer _serverMutex.Unlock()

	timeoutDuration := time.Second * 60
	if frameworkConfig.ServerTimeout != 0 {
		timeoutDuration = frameworkConfig.ServerTimeout
	}

	if frameworkConfig.WebFramework == "" {
		frameworkConfig.WebFramework = defaultWebFramework
	}
	switch frameworkConfig.WebFramework {
	case EchoFramework:
		_server = echo_server.New(timeoutDuration, frameworkConfig.UseHttp2)
	}
	if _server == nil {
		panic("Initialization error. Server is NULL")
	}

	// set http 2.0 server
	if frameworkConfig.UseHttp2 {
		// if config is empty, use default values
		if frameworkConfig.Http2Config == nil {
			frameworkConfig.Http2Config = &domain.Http2Config{
				MaxConcurrentStreams: http2MaxConcurrentStreams,
				MaxReadFrameSize:     http2MaxReadFrameSize,
			}
		}

		// set http 2.0 server config
		_server.SetHttp2Config(frameworkConfig.Http2Config)
	}

	if frameworkConfig.UseSwagger {
		_server.ActivateSwagger()
	}
}

func initGrpcServer(frameworkConfig *Config) {
	if _grpcServer != nil || _initDone {
		return
	}

	if !frameworkConfig.UseGRPC {
		return
	}

	error_helper.LogGrpc = frameworkConfig.LogGRPC

	_useGrpc = frameworkConfig.UseGRPC
	_grpcServerMutex.Lock()
	defer _grpcServerMutex.Unlock()

	_grpcServer = grpc_server.New()
}

func runShutDownWaiter() {
	// create a channel to receive signals
	signalChannel := make(chan os.Signal, 1)

	// notify the signal channel when a SIGINT or SIGTERM signal is received
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	<-signalChannel

	// run shut down actions
	_shutdownService.Run()

	// call shutdown
	os.Exit(0)
}
