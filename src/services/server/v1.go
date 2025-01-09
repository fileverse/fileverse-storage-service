package server

import (
	"storage/src/services/server/handler"
	"storage/src/services/server/middleware"

	"github.com/gin-gonic/gin"
)

func v1configs(v1 *gin.RouterGroup) {
	// v1.Use(middleware.ApiKeyAuth())

	// v1.POST("/upload", middleware.UcanVerify(), handler.UploadFile)
	v1.POST("/upload/identity", middleware.IdentityUcanVerify(), handler.UploadIdentityFile)
}
