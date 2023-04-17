package static_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/exceptions"
	"net/http"
)

func (controller *Controller) Health(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

func (controller *Controller) RouteNotFound(ctx echo.Context) error {
	return controller.NotFound(ctx, exceptions.RouteNotFound)
}
