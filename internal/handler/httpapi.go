package handler

import (
	"context"
	"net/http"

	apiv1 "github.com/ahmadmilzam/ewallet/internal/handler/apiv1"
	"github.com/ahmadmilzam/ewallet/internal/handler/middleware"
	"github.com/ahmadmilzam/ewallet/internal/handler/server"
	"github.com/gin-gonic/gin"
)

// Dependencies contains dependencies required for server and handlers.
type Dependencies struct{}

func Serve(ctx context.Context, addr string, deps Dependencies) error {
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
	}
}
