package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/domain"
	"google.golang.org/grpc"
)

type IServer interface {
	Start(port string)
	StartHttp2(port string)
	SetHttp2Config(config *domain.Http2Config)
	ActivateSwagger(customEndpoint ...string)
}

type IEchoServer interface {
	Get() *echo.Echo
}

type IGRPCServer interface {
	Start(port string)
	Close() error
	Get() *grpc.Server
}
