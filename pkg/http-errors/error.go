package httperrors

import (
	"fmt"
	"strconv"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s-%s ", e.Code, e.Message)
}

func GenerateError(code string, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func GetStatusCode(err error) int {
	errVal := err.Error()
	val, _ := strconv.Atoi(errVal[:3])
	return val
}

func GetCaseCode(err string) string {
	return err[:5]
}
