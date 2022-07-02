package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct{}

func (a *App) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "app pong",
	})
}
