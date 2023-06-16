package owl

import (
	"github.com/lowl11/lazylog/logapi"
	"github.com/lowl11/owl/data/domain"
	"github.com/lowl11/owl/data/interfaces"
	"github.com/lowl11/owl/internal/controllers"
	"github.com/lowl11/owl/internal/shutdown_service"
	"sync"
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

type Owl struct {
	config *Config

	server      interfaces.IServer
	serverMutex sync.Mutex

	useGrpc         bool
	grpcServer      interfaces.IGRPCServer
	grpcServerMutex sync.Mutex

	shutdownService *shutdown_service.Service
}

func New(config *Config) *Owl {
	owl := &Owl{
		config:          config,
		shutdownService: shutdown_service.New(),
	}

	controllers.Init()
	owl.initConfig()
	owl.initLog()
	owl.initServer()
	owl.initGrpcServer()
	go owl.runShutDownWaiter()
	return owl
}
