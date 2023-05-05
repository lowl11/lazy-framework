package domain

import (
	"github.com/lowl11/lazy-collection/array"
	"net/http"
	"strings"
)

type Exception struct {
	BusinessMessage string `json:"business_message,omitempty"`
	httpStatus      int

	withinErrors []error
}

func NewException(businessMessage string, status int) *Exception {
	return &Exception{
		BusinessMessage: businessMessage,
		httpStatus:      status,
		withinErrors:    make([]error, 0),
	}
}

func (exception *Exception) Error() string {
	return exception.BusinessMessage + exception.techMessage(true)
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
	return exception.techMessage(false)
}

func (exception *Exception) With(err error) *Exception {
	exception.withinErrors = append(exception.withinErrors, err)
	return exception.copy()
}

func (exception *Exception) HttpStatus() int {
	if exception.httpStatus == 0 {
		return http.StatusInternalServerError
	}

	return exception.httpStatus
}

func (exception *Exception) copy() *Exception {
	errorCopy := &Exception{
		BusinessMessage: exception.BusinessMessage,
		httpStatus:      exception.httpStatus,
		withinErrors:    exception.withinErrors,
	}
	return errorCopy
}

func (exception *Exception) techMessage(fullMessage bool) string {
	withinMessages := make([]string, 0, len(exception.withinErrors))
	array.NewWithList[error](exception.withinErrors...).Each(func(item error) {
		withinMessages = append(withinMessages, item.Error())
	})

	var withinMessage string
	if len(withinMessages) > 0 {
		if fullMessage {
			withinMessage = " --> "
		}
		withinMessage += strings.Join(withinMessages, " | ")
	}

	return withinMessage
}
