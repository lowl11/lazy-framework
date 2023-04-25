package exceptions

import (
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
)

func New(techMessage, businessMessage string, statuses ...int) interfaces.IException {
	var httpStatus int
	if len(statuses) > 0 {
		httpStatus = statuses[0]
	}

	return domain.NewException(
		techMessage,
		businessMessage,
		httpStatus,
	)
}
