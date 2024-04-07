package v1

import (
	"net/http"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
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

	if err := ctx.ShouldBindJSON(&params); err != nil {
		logger.ErrAttr(err)
		ctx.JSON(entity.GetStatusCode(err), entity.ErrorCodeResponse(err))
		return
	}

	blog, err := route.usecase.CreateAccount(ctx, params)

	if err != nil {
		ctx.JSON(entity.GetStatusCode(err), entity.ErrorCodeResponse(err))
	}

	ctx.JSON(http.StatusCreated, blog)
}

func (route *AccountRoute) getAccount(ctx *gin.Context) {
	var req usecase.GetAccountReqParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.ErrAttr(err)
		ctx.JSON(entity.GetStatusCode(err), entity.ErrorCodeResponse(err))
		return
	}

	blog, err := route.usecase.GetAccount(ctx, req.Phone)

	if err != nil {
		ctx.JSON(entity.GetStatusCode(err), entity.ErrorCodeResponse(err))
	}

	ctx.JSON(http.StatusCreated, blog)
}
