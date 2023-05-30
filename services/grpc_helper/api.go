package grpc_helper

import (
	"github.com/lowl11/lazy-framework/data/exceptions"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"github.com/lowl11/lazy-framework/services/error_helper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToException(err error) interfaces.IException {
	if err == nil {
		return nil
	}

	grpcError, ok := status.FromError(err)
	if grpcError == nil && !ok {
		return exceptions.FromError(err)
	}

	return exceptions.New(grpcError.Message(), error_helper.HttpCode(grpcError.Code()))
}

func IsCode(err error, code codes.Code) bool {
	if grpcError, ok := status.FromError(err); ok {
		return grpcError.Code() == code
	}

	return false
}
