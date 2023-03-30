package interfaces

import "github.com/labstack/echo/v4"

type IServer interface {
	Start(port string)
	ActivateSwagger()
}

type IEchoServer interface {
	Get() *echo.Echo
}
