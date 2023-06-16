package owl

import (
	"github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazylog/log/log_internal"
	"github.com/lowl11/lazylog/logapi/log_levels"
	"github.com/lowl11/owl/data/domain"
	"github.com/lowl11/owl/data/enums/frameworks"
	"github.com/lowl11/owl/data/interfaces"
	"github.com/lowl11/owl/helpers/error_helper"
	"github.com/lowl11/owl/internal/echo_server"
	"github.com/lowl11/owl/internal/grpc_server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultWebFramework = "echo"

	http2MaxConcurrentStreams = 250
	http2MaxReadFrameSize     = 1048576
)

func (owl *Owl) getGrpcServer() interfaces.IGRPCServer {
	if !owl.config.UseGRPC {
		panic("Set flag UseGRPC in config")
	}

	owl.grpcServerMutex.Lock()
	defer owl.grpcServerMutex.Unlock()
	return owl.grpcServer
}

func (owl *Owl) initLog() {
	logLevel := owl.config.LogLevel
	if owl.config.LogLevel == 0 && config.IsProduction() {
		logLevel = log_levels.INFO
	}

	log_internal.Init(log_internal.LogConfig{
		FileName:      owl.config.LogFileName,
		FolderName:    owl.config.LogFolderName,
		NoFile:        owl.config.LogNoFile,
		NoTime:        owl.config.LogNoTime,
		NoPrefix:      owl.config.LogNoPrefix,
		JsonMode:      owl.config.LogJson,
		LogLevel:      logLevel,
		CustomLoggers: owl.config.CustomLoggers,
	})
}

func (owl *Owl) initConfig() {
	config.SetEnvironmentName(owl.config.EnvironmentName)
	config.SetEnvironmentDefault(owl.config.EnvironmentDefault)
	config.SetEnvironmentFileName(owl.config.EnvironmentFileName)
	config.Init()
}

func (owl *Owl) initServer() {
	// HTTP server already exist
	if owl.server != nil {
		return
	}

	// only gRPC server (no HTTP)
	if owl.config.UseGRPC && owl.config.OnlyGRPC {
		return
	}

	owl.serverMutex.Lock()
	defer owl.serverMutex.Unlock()

	timeoutDuration := time.Second * 60
	if owl.config.ServerTimeout != 0 {
		timeoutDuration = owl.config.ServerTimeout
	}

	if owl.config.WebFramework == "" {
		owl.config.WebFramework = defaultWebFramework
	}
	switch owl.config.WebFramework {
	case frameworks.EchoFramework:
		owl.server = echo_server.New(timeoutDuration, owl.config.UseHttp2)
	}
	if owl.server == nil {
		panic("Initialization error. Server is NULL")
	}

	// set http 2.0 server
	if owl.config.UseHttp2 {
		// if config is empty, use default values
		if owl.config.Http2Config == nil {
			owl.config.Http2Config = &domain.Http2Config{
				MaxConcurrentStreams: http2MaxConcurrentStreams,
				MaxReadFrameSize:     http2MaxReadFrameSize,
			}
		}

		// set http 2.0 server config
		owl.server.SetHttp2Config(owl.config.Http2Config)
	}

	if owl.config.UseSwagger {
		owl.server.ActivateSwagger()
	}
}

func (owl *Owl) initGrpcServer() {
	if owl.grpcServer != nil {
		return
	}

	if !owl.config.UseGRPC {
		return
	}

	error_helper.LogGrpc = owl.config.LogGRPC

	owl.grpcServerMutex.Lock()
	defer owl.grpcServerMutex.Unlock()

	owl.grpcServer = grpc_server.New()
}

func (owl *Owl) runShutDownWaiter() {
	// create a channel to receive signals
	signalChannel := make(chan os.Signal, 1)

	// notify the signal channel when a SIGINT or SIGTERM signal is received
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	<-signalChannel

	// run shut down actions
	owl.shutdownService.Run()

	// call shutdown
	os.Exit(0)
}
