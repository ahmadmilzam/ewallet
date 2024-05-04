package response

import (
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
)

type RestAPIReponse struct {
	Success bool              `json:"success"`
	Data    any               `json:"data,omitempty"`
	Error   *httperrors.Error `json:"error,omitempty"`
}

func Success(data any) RestAPIReponse {
	return RestAPIReponse{
		Success: true,
		Data:    data,
		Error:   nil,
	}
}

func Error(err *httperrors.Error) RestAPIReponse {
	return RestAPIReponse{
		Success: false,
		Error:   err,
		Data:    nil,
	}
}
