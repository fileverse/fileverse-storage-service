package middleware

import (
	"storage/src/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RequestRequirements() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		beforeRequest(ctx)
		ctx.Next()
		afterRequest(ctx)
	}
}

func beforeRequest(c *gin.Context) {
	requestUrl := c.Request.URL.Path
	requestMethod := c.Request.Method

	logger := logger.Logger.WithGroup("request").With("url", requestUrl, "method", requestMethod)
	c.Set("logger", logger)
}

func afterRequest(ctx *gin.Context) {
	// Empty for now
}
