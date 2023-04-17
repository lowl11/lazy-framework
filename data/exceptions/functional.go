package exceptions

import (
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/data/interfaces"
)

func New(techMessage, businessMessage string) interfaces.IException {
	return &domain.Exception{
		TechMessage:     techMessage,
		BusinessMessage: businessMessage,
	}
}
