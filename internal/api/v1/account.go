package v1

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/gin-gonic/gin"
)

type AccountRoute struct {
	usecase usecase.AccountUsecaseInterface
}

func NewAccountRoute(handler *gin.RouterGroup, u usecase.AccountUsecaseInterface) {
	route := &AccountRoute{u}
	h := handler.Group("/accounts")
	{
		h.POST("/", route.createAccount)
		h.GET("/:phone", route.getAccount)
	}
}

type waResponse struct {
	Account entity.Account `json:"account_detail"`
	Wallet  entity.Wallet  `json:"wallet_detail"`
}

func (route *AccountRoute) createAccount(ctx *gin.Context) {
	var params usecase.CreateAccountReqParams
	c := context.Background()

	if err := ctx.ShouldBindJSON(&params); err != nil {
		slog.Error("Fail to parse", "error", err)
		ctx.JSON(
			http.StatusBadRequest,
			httpres.GenerateErrResponse(errors.New("40000_fail to parse request"), "Fail to parse request"),
		)
		return
	}

	a, w, err := route.usecase.CreateAccount(c, params)

	if err != nil {
		slog.Error("Fail to create account", "error", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to create account"),
		)
		return
	}

	ctx.JSON(httpres.GetStatusCode(err), httpres.GenerateOK(waResponse{
		Account: *a,
		Wallet:  *w,
	}))
}

func (route *AccountRoute) getAccount(ctx *gin.Context) {
	var req usecase.GetAccountReqParams
	c := context.Background()
	// phone := phonenumber.Parse(ctx.Param("phone"), "ID")
	if err := ctx.ShouldBindUri(&req); err != nil {
		er := errors.New("bad param phone")
		err := fmt.Errorf("%s: %w", httpres.GenericBadRequest, er)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

	fmt.Println("handler/phone: ", req.Phone)

	account, err := route.usecase.GetAccount(c, req.Phone)

	if err != nil {
		var msg string
		sc := httpres.GetStatusCode(err)
		if sc >= 500 {
			msg = "Internal server error"
		} else {
			msg = "Account not found"
		}

		ctx.Set("msg", "Unable to get account")
		ctx.Set("err", err)
		ctx.JSON(
			sc,
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	ctx.JSON(http.StatusOK, httpres.GenerateOK(account))
}
