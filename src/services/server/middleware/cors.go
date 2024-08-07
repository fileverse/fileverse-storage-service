package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Api-Key"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	config.AllowAllOrigins = true
	return cors.New(config)
}
