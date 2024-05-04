package v1

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ahmadmilzam/ewallet/internal/usecase"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/ahmadmilzam/ewallet/pkg/httpserver"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/ahmadmilzam/ewallet/pkg/response"
	"github.com/ahmadmilzam/ewallet/pkg/validator"
	"github.com/gorilla/mux"
)

type AccountRoute struct {
	usecase usecase.AppUsecaseInterface
}

func NewAccountHandler(router *mux.Router, u usecase.AppUsecaseInterface) {
	route := &AccountRoute{u}
	accountRouter := router.PathPrefix("/accounts").Subrouter()

	accountRouter.HandleFunc("/", route.createAccount).Methods(http.MethodPost)
	accountRouter.HandleFunc("/{phone}", route.getAccount).Methods(http.MethodGet)
}

func (route *AccountRoute) getAccount(w http.ResponseWriter, r *http.Request) {
	phone := mux.Vars(r)["phone"]

	if !validator.IsValidPhone(phone) {
		err := httperrors.GenerateError(httperrors.InvalidPhone, "Invalid param {phone}")
		slog.Error("Bad request", logger.ErrAttr(err))
		httpserver.WriteJSON(w, r, httperrors.GetStatusCode(err), response.Error(err))
		return
	}

	c := context.Background()

	response := route.usecase.GetAccount(c, phone)
	status := http.StatusOK
	if !response.Success {
		status = httperrors.GetStatusCode(response.Error)
	}
	httpserver.WriteJSON(w, r, status, response)
}

func (route *AccountRoute) createAccount(w http.ResponseWriter, r *http.Request) {
	var err *httperrors.Error
	var status int
	var params usecase.CreateAccountRequest

	err = httpserver.DecodeJSON(w, r, &params)
	if err != nil {
		httpserver.WriteJSON(w, r, httperrors.GetStatusCode(err), response.Error(err))
		return
	}

	response := route.usecase.CreateAccount(context.Background(), params)

	status = http.StatusCreated
	if response.Error != nil {
		status = httperrors.GetStatusCode(response.Error)
	}

	httpserver.WriteJSON(w, r, status, response)
}
