package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/models"
)

type IController interface {
	Ok(ctx echo.Context, response interface{}, messages ...string) error
	Error(ctx echo.Context, err *models.Error, status ...int) error
	NotFound(ctx echo.Context, err *models.Error) error
	Unauthorized(ctx echo.Context, err *models.Error) error

	RequiredField(value interface{}, name string) error
	FilterString(value string) string
	FilterStringSimple(value string) string
}
