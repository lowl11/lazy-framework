package main

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

func (lazy *Owl) getGrpcServer() interfaces.IGRPCServer {
	if !lazy.config.UseGRPC {
		panic("Set flag UseGRPC in config")
	}

	lazy.grpcServerMutex.Lock()
	defer lazy.grpcServerMutex.Unlock()
	return lazy.grpcServer
}

func (lazy *Owl) initLog() {
	logLevel := lazy.config.LogLevel
	if lazy.config.LogLevel == 0 && config.IsProduction() {
		logLevel = log_levels.INFO
	}

	log_internal.Init(log_internal.LogConfig{
		FileName:      lazy.config.LogFileName,
		FolderName:    lazy.config.LogFolderName,
		NoFile:        lazy.config.LogNoFile,
		NoTime:        lazy.config.LogNoTime,
		NoPrefix:      lazy.config.LogNoPrefix,
		JsonMode:      lazy.config.LogJson,
		LogLevel:      logLevel,
		CustomLoggers: lazy.config.CustomLoggers,
	})
}

func (lazy *Owl) initConfig() {
	config.SetEnvironmentName(lazy.config.EnvironmentName)
	config.SetEnvironmentDefault(lazy.config.EnvironmentDefault)
	config.SetEnvironmentFileName(lazy.config.EnvironmentFileName)
	config.Init()
}

func (lazy *Owl) initServer() {
	// HTTP server already exist
	if lazy.server != nil {
		return
	}

	// only gRPC server (no HTTP)
	if lazy.config.UseGRPC && lazy.config.OnlyGRPC {
		return
	}

	lazy.serverMutex.Lock()
	defer lazy.serverMutex.Unlock()

	timeoutDuration := time.Second * 60
	if lazy.config.ServerTimeout != 0 {
		timeoutDuration = lazy.config.ServerTimeout
	}

	if lazy.config.WebFramework == "" {
		lazy.config.WebFramework = defaultWebFramework
	}
	switch lazy.config.WebFramework {
	case frameworks.EchoFramework:
		lazy.server = echo_server.New(timeoutDuration, lazy.config.UseHttp2)
	}
	if lazy.server == nil {
		panic("Initialization error. Server is NULL")
	}

	// set http 2.0 server
	if lazy.config.UseHttp2 {
		// if config is empty, use default values
		if lazy.config.Http2Config == nil {
			lazy.config.Http2Config = &domain.Http2Config{
				MaxConcurrentStreams: http2MaxConcurrentStreams,
				MaxReadFrameSize:     http2MaxReadFrameSize,
			}
		}

		// set http 2.0 server config
		lazy.server.SetHttp2Config(lazy.config.Http2Config)
	}

	if lazy.config.UseSwagger {
		lazy.server.ActivateSwagger()
	}
}

func (lazy *Owl) initGrpcServer() {
	if lazy.grpcServer != nil {
		return
	}

	if !lazy.config.UseGRPC {
		return
	}

	error_helper.LogGrpc = lazy.config.LogGRPC

	lazy.grpcServerMutex.Lock()
	defer lazy.grpcServerMutex.Unlock()

	lazy.grpcServer = grpc_server.New()
}

func (lazy *Owl) runShutDownWaiter() {
	// create a channel to receive signals
	signalChannel := make(chan os.Signal, 1)

	// notify the signal channel when a SIGINT or SIGTERM signal is received
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	<-signalChannel

	// run shut down actions
	lazy.shutdownService.Run()

	// call shutdown
	os.Exit(0)
}
