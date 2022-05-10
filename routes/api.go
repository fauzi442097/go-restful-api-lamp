package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "sukses",
		})
	})

	return route
}
