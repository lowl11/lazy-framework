package errors

import "github.com/lowl11/lazy-framework/data/models"

func New(techMessage, businessMessage string) *models.Error {
	return &models.Error{
		TechMessage:     techMessage,
		BusinessMessage: businessMessage,
	}
}
