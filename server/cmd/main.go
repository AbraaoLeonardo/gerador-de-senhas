package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"server/routes"
	"time"
)

func main() {
	router := gin.Default()

	// ✅ CORS configuration
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Vue dev server
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	routes.RegisterRoutes(router)

	router.Run(":8080")
}
