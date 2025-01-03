package server

import (
	"storage/src/services/server/handler"
	"storage/src/services/server/middleware"

	"github.com/gin-gonic/gin"
)

func v1configs(v1 *gin.RouterGroup) {
	v1.Use(middleware.ApiKeyAuth())

	// v1.POST("/upload", middleware.UcanVerify(), handler.UploadFile)
	// TODO: fix the ucan part with identity contract based verifications
	v1.POST("/upload/identity", handler.UploadIdentityFile)
}
