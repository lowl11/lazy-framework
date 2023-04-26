package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/log"
	"net/http"
)

type Base struct{}

func (controller *Base) Error(ctx echo.Context, err interfaces.IException) error {
	log.Error(err.ToError())

	errorObject := &domain.Response{
		Status:       "ERROR",
		Message:      err.Business(),
		InnerMessage: err.Tech(),
	}
	return ctx.JSON(err.HttpStatus(), errorObject)
}

func (controller *Base) NotFound(ctx echo.Context, err interfaces.IException) error {
	errorObject := &domain.Response{
		Status:       "ERROR",
		Message:      err.Business(),
		InnerMessage: err.Tech(),
	}
	return ctx.JSON(http.StatusNotFound, errorObject)
}

func (controller *Base) Unauthorized(ctx echo.Context, err interfaces.IException) error {
	errorObject := &domain.Response{
		Status:       "ERROR",
		Message:      err.Business(),
		InnerMessage: err.Tech(),
	}
	return ctx.JSON(http.StatusUnauthorized, errorObject)
}

func (controller *Base) Ok(ctx echo.Context, response any, messages ...string) error {
	defaultMessage := "OK"
	if len(messages) > 0 {
		defaultMessage = messages[0]
	}

	successObject := &domain.Response{
		Status:       "OK",
		Message:      defaultMessage,
		InnerMessage: defaultMessage,
		Body:         response,
	}
	return ctx.JSON(http.StatusOK, successObject)
}

func (controller *Base) OkAny(ctx echo.Context, response any) error {
	return ctx.JSON(http.StatusOK, response)
}
