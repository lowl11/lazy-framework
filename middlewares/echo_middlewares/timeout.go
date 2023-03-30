package echo_middlewares

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func Timeout(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Запрос достигнул таймаута",
		Timeout:      timeout,

		OnTimeoutRouteErrorHandler: func(err error, ctx echo.Context) {
			timeoutError := errors.New("request reached timeout | " + err.Error())
			ctx.Error(timeoutError)
		},
	})
}
