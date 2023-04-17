package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/domain"
)

type IServer interface {
	Start(port string)
	StartHttp2(port string)
	SetHttp2Config(config *domain.Http2Config)
	ActivateSwagger()
}

type IEchoServer interface {
	Get() *echo.Echo
}
