package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/helpers/error_helper"
	"github.com/lowl11/lazy-framework/helpers/type_helper"
	"github.com/lowl11/lazy-framework/helpers/validation"
	"github.com/lowl11/lazy-framework/log"
	"google.golang.org/grpc/codes"
	"net/http"
)

type Base struct{}

func (controller *Base) Error(ctx echo.Context, err interfaces.IException) error {
	log.Error(err.ToError())

	httpStatus := err.HttpStatus()
	grpcStatus := err.GrpcStatus()
	if httpStatus == http.StatusInternalServerError && grpcStatus != codes.Internal {
		httpStatus = error_helper.HttpCode(grpcStatus)
	}

	errorObject := &domain.Response{
		Status:       "ERROR",
		Message:      err.Business(),
		InnerMessage: err.Tech(),
	}
	return ctx.JSON(httpStatus, errorObject)
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
	if validation.IsPrimitive(response) {
		return ctx.String(http.StatusOK, type_helper.ToString(response))
	}

	return ctx.JSON(http.StatusOK, response)
}
