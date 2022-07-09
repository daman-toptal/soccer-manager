package error

import (
	"context"
	"net/http"
	"protobuf-v1/golang"

	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

func NewError(ctx context.Context, code golang.Error, message ...string) error {
	if msg, ok := errorMap[code]; ok {
		if len(message) > 0 {
			return ErrorWithStatus(code, codes.Code(msg.HttpCode), msg.Description, message[0])
		}
		return ErrorWithStatus(code, codes.Code(msg.HttpCode), msg.Description, msg.Description)
	}
	if len(message) > 0 {
		return ErrorWithStatus(code, http.StatusInternalServerError, message[0], message[0])
	}
	return ErrorWithStatus(code, http.StatusInternalServerError, "")
}

func GetHttpErrCode(err error) codes.Code {
	return grpcStatus.Code(err)
}
