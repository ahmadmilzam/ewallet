package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/gin-gonic/gin"
)

type WalletRoute struct {
	usecase usecase.AppUsecaseInterface
}

func NewWalletRoute(handler *gin.RouterGroup, u usecase.AppUsecaseInterface) {
	route := &WalletRoute{u}
	h := handler.Group("/wallets")
	{
		// h.POST("/", route.createWallet)
		h.GET("/:id", route.getWallet)
	}
}

func (route *WalletRoute) getWallet(ctx *gin.Context) {
	c := context.Background()
	id := ctx.Param("id")

	if err := ctx.ShouldBindUri(&id); err != nil {
		err = fmt.Errorf("%s: getWallet: %w", httpres.GenericBadRequest, err)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			httpres.GetStatusCode(err),
			httpres.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

	account, err := route.usecase.GetWallets(c, id)

	if err != nil {
		var msg string
		sc := httpres.GetStatusCode(err)
		if sc >= 500 {
			msg = "Internal server error"
		} else {
			msg = "Wallet not found"
		}

		ctx.Set("msg", "Unable to get wallet")
		ctx.Set("err", err)
		ctx.JSON(
			sc,
			httpres.GenerateErrResponse(err, msg),
		)
		return
	}

	ctx.JSON(http.StatusOK, httpres.GenerateOkResponse(account))
}
