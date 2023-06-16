package static_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/owl/data/exceptions"
)

func (controller *Controller) Health(ctx echo.Context) error {
	return controller.OkAny(ctx, "OK")
}

func (controller *Controller) RouteNotFound(ctx echo.Context) error {
	return controller.NotFound(ctx, exceptions.RouteNotFound)
}
