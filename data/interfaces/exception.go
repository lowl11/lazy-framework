package interfaces

import (
	"github.com/lowl11/owl/data/domain"
	"google.golang.org/grpc/codes"
)

type IException interface {
	ToString() string
	ToError() error
	Business() string
	Tech() string
	With(err error) *domain.Exception
	HttpStatus() int
	GrpcStatus() codes.Code
	ToGrpc() error
}
