package routes

import (
	"github.com/gin-gonic/gin"
	"server/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	password := router.Group("/password")
	{
		password.GET("/", handlers.GetPassword)
		password.POST("/", handlers.PostPassword)
	}
}
