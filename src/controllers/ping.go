package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "server is running successfully",
		"status":  200,
	})
}
