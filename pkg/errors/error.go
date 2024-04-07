package errors

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var (
	ErrInternalServerError = errors.New("INTERNAL_SERVER_ERROR")
	ErrNotFound            = errors.New("NOT_FOUND")
	ErrBadRequest          = errors.New("BAD_REQUEST")
	ErrConflict            = errors.New("CONFLICT")
	ErrInsufficientFund    = errors.New("INSUFFICIENT_FUND")
	ErrUnauthorized        = errors.New("UNAUTHORIZED")
	errStruct              ErrorResponse
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorStruct `json:"error"`
}

// ErrorCodesStruct This is the struct for the error codes -.
type ErrorStruct struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetStatusCode Fetched the status code from the error -.
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	errCode := extractErroCode(err)

	switch errCode {
	case ErrInternalServerError.Error():
		return http.StatusInternalServerError
	case ErrNotFound.Error():
		return http.StatusNotFound
	case ErrConflict.Error():
		return http.StatusConflict
	case ErrInsufficientFund.Error():
		return http.StatusBadRequest
	case ErrUnauthorized.Error():
		return http.StatusUnauthorized
	case ErrBadRequest.Error():
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// Gets the error and extracts the error code from its json string -.
func extractErroCode(err error) string {
	s := err.Error()

	_ = json.Unmarshal([]byte(s), &errStruct)

	return strings.ToUpper(errStruct.Error.Code)
}

// ErrorCodeResponse The response message for the error -.
func ErrorCodeResponse(err error) ErrorResponse {
	s := err.Error()

	_ = json.Unmarshal([]byte(s), &errStruct)

	return errStruct
}
