package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsHandler() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowCredentials: false,
		AllowHeaders:     []string{"*"},
		MaxAge:           12 * time.Hour,
	}

	config.AllowAllOrigins = true
	return cors.New(config)
}
