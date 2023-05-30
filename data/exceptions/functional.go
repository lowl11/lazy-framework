package exceptions

import (
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
)

func New(businessMessage string, statuses ...int) interfaces.IException {
	var httpStatus int
	var grpcStatus int

	if len(statuses) > 0 {
		httpStatus = statuses[0]
	}
	if len(statuses) > 1 {
		grpcStatus = statuses[1]
	}

	return domain.NewException(
		businessMessage,
		httpStatus,
		grpcStatus,
	)
}

func FromError(err error, status ...int) interfaces.IException {
	if err == nil {
		return nil
	}

	var httpStatus int
	var grpcStatus int

	if len(status) > 0 {
		httpStatus = status[0]
	}

	if len(status) > 1 {
		grpcStatus = status[1]
	}

	return domain.NewException(
		err.Error(),
		httpStatus,
		grpcStatus,
	)
}
