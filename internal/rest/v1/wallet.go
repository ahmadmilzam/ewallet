package v1

/*
type WalletRoute struct {
	usecase usecase.AppUsecaseInterface
}

func NewWalletRoute(handler *gin.RouterGroup, u usecase.AppUsecaseInterface) {
	route := &WalletRoute{u}
	h := handler.Group("/wallets")
	{
		// h.POST("/", route.createWallet)
		h.GET("/:id", route.getWallet)
		h.GET("/phone/:phone", route.getWalletsByPhone)
	}
}

func (route *WalletRoute) getWallet(ctx *gin.Context) {
	c := context.Background()
	id := ctx.Param("id")

	if err := ctx.ShouldBindUri(&id); err != nil {
		err = fmt.Errorf("%s: getWallet: %w", appErr.GenericBadRequest, err)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			appErr.GetStatusCode(err),
			appErr.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

	account, err := route.usecase.GetWallet(c, id)

	if err != nil {
		var msg string
		sc := appErr.GetStatusCode(err)
		if sc >= 500 {
			msg = "Internal server error"
		} else {
			msg = "Wallet not found"
		}

		ctx.Set("msg", "Unable to get wallet")
		ctx.Set("err", err)
		ctx.JSON(
			sc,
			appErr.GenerateErrResponse(err, msg),
		)
		return
	}

	ctx.JSON(http.StatusOK, appErr.GenerateOkResponse(account))
}

func (route *WalletRoute) getWalletsByPhone(ctx *gin.Context) {
	c := context.Background()
	phone := ctx.Param("phone")

	if err := ctx.ShouldBindUri(&phone); err != nil {
		err = fmt.Errorf("%s: getWalletsByPhone: %w", appErr.GenericBadRequest, err)

		ctx.Set("msg", "Fail to parse request data")
		ctx.Set("err", err)
		ctx.JSON(
			appErr.GetStatusCode(err),
			appErr.GenerateErrResponse(err, "Fail to parse request"),
		)
		return
	}

	account, err := route.usecase.GetWallets(c, phone)

	if err != nil {
		var msg string
		sc := appErr.GetStatusCode(err)
		if sc >= 500 {
			msg = "Internal server error"
		} else {
			msg = "Wallet not found"
		}

		ctx.Set("msg", "Unable to get wallet")
		ctx.Set("err", err)
		ctx.JSON(
			sc,
			appErr.GenerateErrResponse(err, msg),
		)
		return
	}

	ctx.JSON(http.StatusOK, appErr.GenerateOkResponse(account))
}
*/
