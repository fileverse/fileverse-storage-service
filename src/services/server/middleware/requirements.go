package middleware

import (
	"storage/src/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestRequirements() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		beforeRequest(ctx)
		ctx.Next()
		afterRequest(ctx)
	}
}

func beforeRequest(c *gin.Context) {
	contractAddress := c.GetHeader("Contract")
	invokerAddress := c.GetHeader("Invoker")
	chainID := c.GetHeader("Chain")
	namespace := c.GetHeader("Namespace")

	requestID := uuid.New().String()
	c.Set("requestId", requestID)
	c.Set("isAuthenticated", false)
	c.Set("invoker", invokerAddress)
	c.Set("contract", contractAddress)
	c.Set("chain", chainID)
	c.Set("namespace", namespace)

	// set logger
	requestUrl := c.Request.URL.Path
	requestMethod := c.Request.Method
	logger := logger.Logger.With("url", requestUrl, "method", requestMethod)
	c.Set("logger", logger)
}

func afterRequest(ctx *gin.Context) {
	// Empty for now
}
