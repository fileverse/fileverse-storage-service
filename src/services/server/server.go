package server

import (
	"net/http"
	"storage/src/services/server/middleware"

	"github.com/gin-gonic/gin"
)

func Build() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.CorsHandler())

	router.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Endpoints
	v1 := router.Group("api/v1")
	v1configs(v1)

	return router

}
