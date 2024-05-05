package api

import (
	"net/http"

	v1 "github.com/ahmadmilzam/ewallet/internal/rest/v1"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/ahmadmilzam/ewallet/pkg/httpserver"
	"github.com/ahmadmilzam/ewallet/pkg/response"
	"github.com/gorilla/mux"
)

func staticHandler(status int, code string, message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errRes := httperrors.GenerateError(code, message)
		httpserver.WriteJSON(w, r, status, response.Error(errRes))
	})
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	httpserver.WriteJSON(w, r, http.StatusOK, response.Success(nil))
}

func RegisterRoutes(router *mux.Router, usecase usecase.AppUsecaseInterface) {
	router.MethodNotAllowedHandler = staticHandler(
		http.StatusMethodNotAllowed,
		httperrors.GenericMethodNotAllowed,
		"method not allowed",
	)

	router.NotFoundHandler = staticHandler(
		http.StatusNotFound,
		httperrors.GenericNotFound,
		"path not found",
	)

	// K8s probe for kubernetes health checks
	router.HandleFunc("/ping", handlePing).Methods(http.MethodGet)
	v1RouterGroup := router.PathPrefix("/v1").Subrouter()

	v1.NewAccountHandler(v1RouterGroup, usecase)
	v1.NewTransferHandler(v1RouterGroup, usecase)
}
