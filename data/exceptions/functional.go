package exceptions

import (
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
)

func New(businessMessage string, statuses ...int) interfaces.IException {
	var httpStatus int
	if len(statuses) > 0 {
		httpStatus = statuses[0]
	}

	return domain.NewException(
		businessMessage,
		httpStatus,
	)
}

func FromError(err error, status ...int) interfaces.IException {
	if err == nil {
		return nil
	}

	var errorStatus int
	if len(status) > 0 {
		errorStatus = status[0]
	}

	return domain.NewException(
		err.Error(),
		errorStatus,
	)
}
