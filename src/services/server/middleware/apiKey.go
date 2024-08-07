package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {
	validAPIKey := os.Getenv("VALID_API_KEY")
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")

		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key required"})
			c.Abort()
			return
		}

		if apiKey != validAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
