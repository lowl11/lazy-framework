package domain

import (
	"github.com/lowl11/lazy-collection/array"
	"github.com/lowl11/lazy-framework/services/error_helper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type Exception struct {
	BusinessMessage string `json:"business_message,omitempty"`
	httpStatus      int
	grpcStatus      codes.Code

	withinErrors []error
}

func NewException(businessMessage string, httpStatus, grpcStatus int) *Exception {
	return &Exception{
		BusinessMessage: businessMessage,
		httpStatus:      httpStatus,
		grpcStatus:      codes.Code(grpcStatus),
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

func (exception *Exception) ToGrpc() error {
	grpcStatus := exception.grpcStatus
	if exception.HttpStatus() != http.StatusInternalServerError && exception.GrpcStatus() == codes.Internal {
		grpcStatus = error_helper.GrpcCode(exception.httpStatus)
	}
	if int(grpcStatus) == 0 {
		grpcStatus = codes.Internal
	}
	return status.Error(grpcStatus, exception.Error())
}

func (exception *Exception) Business() string {
	return exception.BusinessMessage
}

func (exception *Exception) Tech() string {
	return exception.techMessage(false)
}

func (exception *Exception) With(err error) *Exception {
	copyException := exception.copy()
	copyException.withinErrors = append(exception.withinErrors, err)
	return copyException
}

func (exception *Exception) HttpStatus() int {
	if exception.httpStatus == 0 {
		return http.StatusInternalServerError
	}

	return exception.httpStatus
}

func (exception *Exception) GrpcStatus() codes.Code {
	if exception.grpcStatus == 0 {
		return codes.Internal
	}

	return exception.grpcStatus
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
