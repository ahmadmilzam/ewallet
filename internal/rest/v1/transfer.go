package v1

import (
	"net/http"

	"github.com/ahmadmilzam/ewallet/internal/usecase"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/ahmadmilzam/ewallet/pkg/httpserver"
	"github.com/ahmadmilzam/ewallet/pkg/response"
	"github.com/gorilla/mux"
)

type TransferHandler struct {
	usecase usecase.AppUsecaseInterface
}

func NewTransferHandler(router *mux.Router, u usecase.AppUsecaseInterface) {
	handler := &TransferHandler{u}
	accountRouter := router.PathPrefix("/transfers").Subrouter()

	accountRouter.HandleFunc("/", handler.createTransfer).Methods(http.MethodPost)
}

func (handler *TransferHandler) createTransfer(w http.ResponseWriter, r *http.Request) {
	var params usecase.CreateTransferRequest
	var err *httperrors.Error
	var status int

	err = httpserver.DecodeJSON(w, r, &params)
	if err != nil {
		httpserver.WriteJSON(w, r, httperrors.GetStatusCode(err), response.Error(err))
		return
	}

	response := handler.usecase.CreateTransfer(r.Context(), params)

	status = http.StatusCreated
	if response.Error != nil {
		status = httperrors.GetStatusCode(response.Error)
	}

	httpserver.WriteJSON(w, r, status, response)
}
