package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"github.com/lowl11/owl/data/domain"
	"google.golang.org/grpc"
)

type IServer interface {
	Start(port string)
	StartHttp2(port string)
	SetHttp2Config(config domain.Http2Config)
	ActivateSwagger(customEndpoint ...string)
}

type IEchoServer interface {
	Get() *echo.Echo
}

type IFiberServer interface {
	Get() *fiber.App
}

type IGRPCServer interface {
	Start(port string)
	Close() error
	Get() *grpc.Server
}
