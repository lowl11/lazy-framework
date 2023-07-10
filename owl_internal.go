package owl

import (
	"github.com/lowl11/lazyconfig/config"
	"github.com/lowl11/lazyconfig/config/config_internal"
	"github.com/lowl11/lazylog/log/log_internal"
	"github.com/lowl11/lazylog/logapi/log_levels"
	"github.com/lowl11/owl/data/domain"
	"github.com/lowl11/owl/data/enums/frameworks"
	"github.com/lowl11/owl/data/interfaces"
	"github.com/lowl11/owl/helpers/error_helper"
	"github.com/lowl11/owl/internal/echo_server"
	"github.com/lowl11/owl/internal/fiber_server"
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

func (app *App) getGrpcServer() interfaces.IGRPCServer {
	if !app.config.UseGRPC {
		panic("Set flag UseGRPC in config")
	}

	app.grpcServerMutex.Lock()
	defer app.grpcServerMutex.Unlock()
	return app.grpcServer
}

func (app *App) initLog() {
	logLevel := app.config.LogLevel
	if app.config.LogLevel == 0 && config.IsProduction() {
		logLevel = log_levels.INFO
	}

	log_internal.Init(log_internal.LogConfig{
		FileName:      app.config.LogFileName,
		FolderName:    app.config.LogFolderName,
		NoFile:        app.config.LogNoFile,
		NoTime:        app.config.LogNoTime,
		NoPrefix:      app.config.LogNoPrefix,
		JsonMode:      app.config.LogJson,
		LogLevel:      logLevel,
		CustomLoggers: app.config.CustomLoggers,
	})
}

func (app *App) initConfig() {
	config_internal.SetEnvironment(app.config.Environment)
	config_internal.SetEnvironmentVariableName(app.config.EnvironmentVariableName)
	config_internal.SetEnvironmentFileName(app.config.EnvironmentFileName)
	config_internal.SetBaseFolder(app.config.EnvironmentBaseFolder)

	config_internal.Init()
}

func (app *App) initServer() {
	// HTTP server already exist
	if app.server != nil {
		return
	}

	// only gRPC server (no HTTP)
	if app.config.UseGRPC && app.config.OnlyGRPC {
		return
	}

	app.serverMutex.Lock()
	defer app.serverMutex.Unlock()

	timeoutDuration := time.Second * 60
	if app.config.ServerTimeout != 0 {
		timeoutDuration = app.config.ServerTimeout
	}

	if app.config.WebFramework == "" {
		app.config.WebFramework = defaultWebFramework
	}

	switch app.config.WebFramework {
	case frameworks.Echo:
		app.server = echo_server.New(timeoutDuration, app.config.UseHttp2)
	case frameworks.Fiber:
		app.server = fiber_server.New(timeoutDuration, app.config.UseHttp2)
	}

	if app.server == nil {
		panic("Initialization error. Server is NULL")
	}

	// set http 2.0 server
	if app.config.UseHttp2 {
		// if config is empty, use default values
		if app.config.Http2Config.MaxReadFrameSize == 0 {
			app.config.Http2Config = domain.Http2Config{
				MaxConcurrentStreams: http2MaxConcurrentStreams,
				MaxReadFrameSize:     http2MaxReadFrameSize,
			}
		}

		// set http 2.0 server config
		app.server.SetHttp2Config(app.config.Http2Config)
	}

	if app.config.UseSwagger {
		app.server.ActivateSwagger()
	}
}

func (app *App) initGrpcServer() {
	if app.grpcServer != nil {
		return
	}

	if !app.config.UseGRPC {
		return
	}

	error_helper.LogGrpc = app.config.LogGRPC

	app.grpcServerMutex.Lock()
	defer app.grpcServerMutex.Unlock()

	app.grpcServer = grpc_server.New()
}

func (app *App) runShutDownWaiter() {
	// create a channel to receive signals
	signalChannel := make(chan os.Signal, 1)

	// notify the signal channel when a SIGINT or SIGTERM signal is received
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	<-signalChannel

	// run shut down actions
	app.shutdownService.Run()

	// call shutdown
	os.Exit(0)
}
