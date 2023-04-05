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
		ErrorMessage: "Request reached timeout!",
		Timeout:      timeout,

		OnTimeoutRouteErrorHandler: func(err error, ctx echo.Context) {
			timeoutError := errors.New(err.Error())
			ctx.Error(timeoutError)
		},
	})
}
