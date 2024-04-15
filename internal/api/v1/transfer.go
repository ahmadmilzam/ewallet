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
	usecase usecase.TransferUsecaseInterface
}

func NewTransferRoute(handler *gin.RouterGroup, u usecase.TransferUsecaseInterface) {
	route := &TransferRoute{u}
	h := handler.Group("/wallets")
	{
		h.POST("/", route.createTransfer)
	}
}

func (route *TransferRoute) createTransfer(ctx *gin.Context) {
	c := context.Background()
	req := &usecase.CreateTransferReqParams{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		err = fmt.Errorf("%s: createTransfer: %w", httpres.GenericBadRequest, err)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

	t, err := route.usecase.CreateTransfer(c, req)

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

	ctx.JSON(http.StatusOK, httpres.GenerateOK(t))
}
