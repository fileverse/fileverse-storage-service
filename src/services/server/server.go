package server

import (
	"net/http"
	"storage/src/services/server/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Build() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.Use(getCors())

	router.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Endpoints
	v1 := router.Group("api/v1")
	v1.Use(middleware.ApiKeyAuth())
	groupV1Endpoints(v1)

	return router

}

func getCors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Api-Key"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	config.AllowAllOrigins = true
	return cors.New(config)
}
