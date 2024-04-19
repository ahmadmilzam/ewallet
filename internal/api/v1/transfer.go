package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/gin-gonic/gin"
)

type TransferRoute struct {
	usecase usecase.AppUsecaseInterface
}

func NewTransferRoute(handler *gin.RouterGroup, u usecase.AppUsecaseInterface) {
	route := &TransferRoute{u}
	h := handler.Group("/transfers")
	{
		h.POST("/", route.createTransfer)
	}
}

func (route *TransferRoute) createTransfer(ctx *gin.Context) {
	c := context.Background()

	params := &usecase.TransferReqParams{}

	if err := ctx.ShouldBindJSON(params); err != nil {
		err = fmt.Errorf("%s: createTransfer fail to parse request: %w", httpres.GenericBadRequest, err)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

	isValid, err := params.Validate()
	if !isValid {
		msg := "Invalid request data"
		ctx.Set("msg", msg)
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	transferResponse, err := route.usecase.CreateTransfer(c, params)

	if err != nil {
		var msg string
		sc := httpres.GetStatusCode(err)
		if sc >= 500 {
			msg = "Internal server error"
		} else {
			msg = "Fail to create a transfer"
		}

		ctx.Set("msg", msg)
		ctx.Set("err", err)
		ctx.JSON(
			sc,
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	ctx.JSON(http.StatusOK, httpres.GenerateOK(transferResponse))
}
