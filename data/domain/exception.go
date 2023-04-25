package domain

import "net/http"

type Exception struct {
	TechMessage     string `json:"tech_message"`
	BusinessMessage string `json:"business_message"`
	httpStatus      int
}

func NewException(techMessage, businessMessage string, status int) *Exception {
	return &Exception{
		TechMessage:     techMessage,
		BusinessMessage: businessMessage,
		httpStatus:      status,
	}
}

func (exception *Exception) Error() string {
	return exception.TechMessage + " -> " + exception.BusinessMessage
}

func (exception *Exception) ToString() string {
	return exception.Error()
}

func (exception *Exception) ToError() error {
	return exception
}

func (exception *Exception) Business() string {
	return exception.BusinessMessage
}

func (exception *Exception) Tech() string {
	return exception.TechMessage
}

func (exception *Exception) With(err error) *Exception {
	with := exception.copy()
	with.TechMessage = err.Error() + " | " + with.TechMessage
	return with
}

func (exception *Exception) HttpStatus() int {
	if exception.httpStatus == 0 {
		return http.StatusInternalServerError
	}

	return exception.httpStatus
}

func (exception *Exception) copy() *Exception {
	errorCopy := &Exception{
		TechMessage:     exception.TechMessage,
		BusinessMessage: exception.BusinessMessage,
	}
	return errorCopy
}
