package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Account struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func CreateAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
	}
}

func GetAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
	}
}
