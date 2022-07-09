package error

import (
	"net/http"

	"protobuf-v1/golang"
)

type errDescription struct {
	Description    string
	HttpCode       int32
	TranslationKey string
}

var errorMap map[golang.Error]errDescription

func init() {
	errorMap = make(map[golang.Error]errDescription)
	errorMap[golang.Error_ERROR_AUTH_ERROR] = getErrDescription(http.StatusUnauthorized, "unauthorized")
	errorMap[golang.Error_ERROR_TOKEN_HEADER_REQUIRED] = getErrDescription(http.StatusUnauthorized, "token header missing")
	errorMap[golang.Error_ERROR_TOKEN_INVALID] = getErrDescription(http.StatusUnauthorized, "invalid token")
	errorMap[golang.Error_ERROR_TOKEN_ERROR] = getErrDescription(http.StatusInternalServerError, "token error")

	errorMap[golang.Error_ERROR_INTERNAL_ERROR] = getErrDescription(http.StatusInternalServerError, "internal error")
	errorMap[golang.Error_ERROR_NOT_FOUND] = getErrDescription(http.StatusNotFound, "resource not found")
	errorMap[golang.Error_ERROR_INVALID_ARGS] = getErrDescription(http.StatusBadRequest, "invalid args")
	errorMap[golang.Error_ERROR_INVALID_ID] = getErrDescription(http.StatusBadRequest, "invalid id")
}

func getErrDescription(httpCode int32, message string) errDescription {
	return errDescription{HttpCode: httpCode, Description: message}
}
