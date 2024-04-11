package httpres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type sampleData struct {
	ID    string
	Name  string
	Phone string
}

type dataArr []sampleData

func TestGenerateOK(t *testing.T) {
	data := sampleData{
		ID:    "123",
		Name:  "John Doe",
		Phone: "+6281281281200",
	}

	data2 := sampleData{
		ID:    "456",
		Name:  "Jane Doe",
		Phone: "+6281281281211",
	}

	dataSlice := dataArr{data, data2}

	actualFlat := GenerateOK(data)
	actualArr := GenerateOK(dataSlice)

	expectedFlat := HttpResponse{
		Success: true,
		Data:    data,
	}
	expectedArr := HttpResponse{
		Success: true,
		Data:    dataSlice,
	}
	// fe.InitServiceCode(issuer.TransactionManagementService)
	assert.Equal(t, expectedFlat, actualFlat)
	assert.Equal(t, expectedArr, actualArr)
}

// func TestGenerateErr(t *testing.T) {
// 	errMessage := "Bad Request"
// 	actual := GenerateErr(GenericBadRequest, errMessage)
// 	expected := HttpResponse{
// 		Success: false,
// 		Data:    nil,
// 		Error:   ErrorDetails{Code: GenericBadRequest, Message: errMessage},
// 	}
// 	// fe.InitServiceCode(issuer.TransactionManagementService)
// 	assert.Equal(t, expected, actual)
// }

// func TestGetStatusCode(t *testing.T) {
// 	assert.Equal(t, http.StatusBadRequest, GetErrStatusCode(GenericBadRequest))
// 	assert.Equal(t, http.StatusUnauthorized, GetErrStatusCode(GenericUnauthorized))
// 	assert.Equal(t, http.StatusNotFound, GetErrStatusCode(GenericNotFound))
// 	assert.Equal(t, http.StatusMethodNotAllowed, GetErrStatusCode(GenericMethodNotAllowed))
// 	assert.Equal(t, http.StatusConflict, GetErrStatusCode(GenericConflict))
// 	assert.Equal(t, http.StatusUnprocessableEntity, GetErrStatusCode(GenericUnprocessable))
// 	assert.Equal(t, http.StatusTooManyRequests, GetErrStatusCode(GenericTooManyRequests))
// 	assert.Equal(t, http.StatusInternalServerError, GetErrStatusCode(GenericInternalError))
// }
