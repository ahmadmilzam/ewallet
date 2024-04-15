package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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

func (route *AccountRoute) createAccount(ctx *gin.Context) {
	var params usecase.CreateAccountReqParams
	c := context.Background()

	if err := ctx.ShouldBindJSON(&params); err != nil {
		err = fmt.Errorf("%s: r.createAccount: %w", httpres.GenericBadRequest, err)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

	fmt.Println("Req payload parsed")
	fmt.Println("Begin calling u.CreateAccount")
	aw, err := route.usecase.CreateAccount(c, params)
	fmt.Println("Finished called u.CreateAccount")
	fmt.Println("AW Res: ", aw)

	if err != nil {
		ctx.Set("msg", "Fail to create account")
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to create account"),
		)
		return
	}

	ctx.JSON(http.StatusCreated, httpres.GenerateOK(aw))
}

func (route *AccountRoute) getAccount(ctx *gin.Context) {
	var req usecase.GetAccountReqParams
	c := context.Background()

	if err := ctx.ShouldBindUri(&req); err != nil {
		er := errors.New("bad param phone")
		err := fmt.Errorf("%s: r.getAccount: %w", httpres.GenericBadRequest, er)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

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
