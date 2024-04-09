package v1

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/internal/utils"
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
		slog.Error("unprocessable data", "error", err)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Success: false,
			Error: utils.ErrorStruct{
				Code:    "40001",
				Message: "Unprocessable data",
			},
		})
		return
	}

	account, err := route.usecase.CreateAccount(c, params)

	if err != nil {
		slog.Error("fail to create account", "error", err)
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Success: false,
			Error: utils.ErrorStruct{
				Code:    "50001",
				Message: "Fail to store data",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, SuccessResponse{
		Success: true,
		Data:    account,
	})
}

func (route *AccountRoute) getAccount(ctx *gin.Context) {
	var req usecase.GetAccountReqParams
	c := context.Background()
	if err := ctx.ShouldBindUri(&req); err != nil {
		slog.Error("bad request", "param", ctx.Param("phone"), "error", err)
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Success: false,
			Error: utils.ErrorStruct{
				Code:    "40000",
				Message: "Unprocessable data",
			},
		})
		return
	}

	account, err := route.usecase.GetAccount(c, req.Phone)

	if err != nil {
		slog.Error("Unable to get account", "error", err)
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
			Success: false,
			Error: utils.ErrorStruct{
				Code:    "40400",
				Message: "Account not found",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, account)
}
