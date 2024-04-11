package api

import (
	"errors"
	"net/http"

	v1 "github.com/ahmadmilzam/ewallet/internal/api/v1"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, u usecase.AccountUsecaseInterface) {
	// Options -.
	// gin.SetMode(gin.DebugMode)
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())
	router.HandleMethodNotAllowed = true
	// K8s probe for kubernetes health checks -.
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "The server is up and running")
	})

	// Handling a page not found endpoint -.
	router.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusNotFound,
			httpres.GenerateErrResponse(errors.New("NOT_FOUND"), "Endpoint not found"),
		)
	})

	// Routers -.
	rgroupv1 := router.Group("/v1")
	{
		v1.NewAccountRoute(rgroupv1, u)
	}
}
