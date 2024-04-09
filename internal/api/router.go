package api

import (
	"net/http"

	v1 "github.com/ahmadmilzam/ewallet/internal/api/v1"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/internal/utils"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, u usecase.AccountUsecaseInterface) {
	// Options -.
	// gin.SetMode(gin.DebugMode)
	// handler.Use(gin.Logger())
	// handler.Use(gin.Recovery())

	// K8s probe for kubernetes health checks -.
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "The server is up and running")
	})

	// Handling a page not found endpoint -.
	handler.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse{
			Success: false,
			Error: utils.ErrorStruct{
				Code:    "40400",
				Message: "The requested path is not found!",
			},
		})
	})

	// Routers -.
	rgroupv1 := handler.Group("/v1")
	{
		v1.NewAccountRoute(rgroupv1, u)
	}
}
