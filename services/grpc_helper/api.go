package grpc_helper

import (
	"github.com/lowl11/lazy-framework/data/exceptions"
	"github.com/lowl11/lazy-framework/data/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func HttpCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Canceled:
		return 499
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.Aborted:
		return http.StatusConflict
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.InvalidArgument:
	case codes.FailedPrecondition:
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Internal:
	case codes.DataLoss:
	case codes.Unknown:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}

func ToException(err error) interfaces.IException {
	if err == nil {
		return nil
	}

	grpcError, ok := status.FromError(err)
	if grpcError == nil && !ok {
		return exceptions.FromError(err)
	}

	return exceptions.New(grpcError.Message(), HttpCode(grpcError.Code()))
}

func IsCode(err error, code codes.Code) bool {
	if grpcError, ok := status.FromError(err); ok {
		return grpcError.Code() == code
	}

	return false
}
