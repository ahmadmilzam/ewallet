package handler

import (
	"context"
	"net/http"

	apiv1 "github.com/ahmadmilzam/ewallet/handler/apiv1"
	"github.com/ahmadmilzam/ewallet/handler/middleware"
	"github.com/ahmadmilzam/ewallet/handler/server"
	"github.com/gin-gonic/gin"
)

// Dependencies contains dependencies required for server and handlers.
// type Dependencies struct{}

// func Serve(ctx context.Context, addr string, deps Dependencies) error {
func Serve(ctx context.Context, addr string) error {
	srv := server.New()
	if _, err := srv.Register(Routes()); err != nil {
		return err
	}

	return srv.Serve(ctx, addr)
}

// Routes returns an array of server.Route with configured middleware for each.
func Routes() []server.Route {
	instrumentStatsD := middleware.InstrumentStatsD()

	return []server.Route{
		{
			Method: http.MethodGet,
			Path:   "v1/ping",
			Handler: []gin.HandlerFunc{
				instrumentStatsD,
				apiv1.Ping(),
			},
		},
		{
			Method: http.MethodPost,
			Path:   "v1/accounts",
			Handler: []gin.HandlerFunc{
				instrumentStatsD,
				apiv1.CreateAccount(),
			},
		},
		{
			Method: http.MethodGet,
			Path:   "v1/accounts/:id",
			Handler: []gin.HandlerFunc{
				instrumentStatsD,
				apiv1.GetAccount(),
			},
		},
	}
}
