package httpres

import (
	"fmt"
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

func GenerateErrResponse(e error, m string) HttpResponse {
	return HttpResponse{
		Success: false,
		Error: &ErrorDetails{
			Code:    GetCaseCode(e),
			Message: m,
		},
	}
}

func GenerateOK(d any) HttpResponse {
	res := HttpResponse{
		Success: true,
		Data:    d,
		Error:   nil,
	}

	fmt.Println(res)
	return res
}

// func generateErr(c string, m string) HttpResponse {
// 	return HttpResponse{
// 		Success: false,
// 		Error: ErrorDetails{
// 			Code:    c,
// 			Message: m,
// 		},
// 	}
// }

// func GetErrStatusCode(err error) int {
// 	if err == nil {
// 		return http.StatusOK
// 	}

// 	errVal := err.Error()

// 	switch {
// 	case strings.Contains(errVal, "BAD_REQUEST"):
// 		return http.StatusBadRequest
// 	case strings.Contains(errVal, "UNAUTHORIZED"):
// 		return http.StatusUnauthorized
// 	case strings.Contains(errVal, "NOT_FOUND"):
// 		return http.StatusNotFound
// 	case strings.Contains(errVal, "INTERNAL_SERVER_ERROR"):
// 		return http.StatusInternalServerError
// 	case strings.Contains(errVal, "INSUFFICIENT_BALANCE"):
// 		return http.StatusUnprocessableEntity
// 	case strings.Contains(errVal, "CONFLICT"):
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }

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
