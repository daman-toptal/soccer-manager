package error

import (
	"net/http"

	grpcRoot "protobuf-v1/golang"

	grpcCodes "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

func ErrorWithStatus(code grpcRoot.Error, grpcCode grpcCodes.Code, message string, enMessage ...string) error {
	if message == "" {
		httpStatus, ok := httpStatusFromStatusCode[grpcCode]
		if !ok {
			httpStatus = 500
		}

		message = http.StatusText(httpStatus)
	}
	status := grpcStatus.New(grpcCode, message)
	return status.Err()
}

func NewRESTError(httpMethod string, code grpcCodes.Code, Error grpcRoot.Error, message string) *grpcRoot.HttpError {
	httpStatus, ok := httpStatusFromStatusCode[code]
	if !ok {
		httpStatus = 500
	}

	if message == "" {
		message = http.StatusText(httpStatus)
	}

	return &grpcRoot.HttpError{
		Code:    Error,
		Message: message,
	}
}

var httpStatusFromStatusCode = map[grpcCodes.Code]int{
	grpcCodes.OK:                 http.StatusOK,
	grpcCodes.Canceled:           http.StatusInternalServerError,
	grpcCodes.Unknown:            http.StatusInternalServerError,
	grpcCodes.InvalidArgument:    http.StatusBadRequest,
	grpcCodes.DeadlineExceeded:   http.StatusInternalServerError,
	grpcCodes.NotFound:           http.StatusNotFound,
	grpcCodes.AlreadyExists:      http.StatusBadRequest,
	grpcCodes.PermissionDenied:   http.StatusForbidden,
	grpcCodes.ResourceExhausted:  http.StatusInternalServerError,
	grpcCodes.FailedPrecondition: http.StatusPreconditionFailed,
	grpcCodes.Aborted:            http.StatusInternalServerError,
	grpcCodes.OutOfRange:         http.StatusRequestedRangeNotSatisfiable,
	grpcCodes.Unimplemented:      http.StatusNotImplemented,
	grpcCodes.Internal:           http.StatusInternalServerError,
	grpcCodes.Unavailable:        http.StatusServiceUnavailable,
	grpcCodes.DataLoss:           http.StatusInternalServerError,
	grpcCodes.Unauthenticated:    http.StatusUnauthorized,
}

func NewHttpErrorFromError(httpMethod string, err error) *grpcRoot.HttpError {
	httpStatus := grpcStatus.Code(err)
	if httpStatus == grpcCodes.Unknown {
		httpStatus = 500
	}

	code := grpcRoot.Error_ERROR_UNSPECIFIED
	message := http.StatusText(int(httpStatus))
	status := grpcStatus.Convert(err)
	if status != nil {
		if status.Message() != "" {
			message = status.Message()
		}

	}

	return &grpcRoot.HttpError{
		Code:    code,
		Message: message,
	}
}
