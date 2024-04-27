package httpres

import (
	"strconv"
)

type HttpResponse struct {
	Success bool          `json:"success"`
	Error   *ErrorDetails `json:"error,omitempty"`
	Data    any           `json:"data,omitempty"`
}

type ErrorDetails struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func GenerateErrResponse(err error, message string) HttpResponse {
	return HttpResponse{
		Success: false,
		Error: &ErrorDetails{
			Code:    GetCaseCode(err),
			Message: message,
		},
	}
}

func GenerateOkResponse(data any) HttpResponse {
	res := HttpResponse{
		Success: true,
		Data:    data,
		Error:   nil,
	}
	return res
}

func GetStatusCode(e error) int {
	errVal := e.Error()
	val, _ := strconv.Atoi(errVal[:3])
	return val
}

func GetCaseCode(e error) int {
	errVal := e.Error()
	val, _ := strconv.Atoi(errVal[:5])
	return val
}
