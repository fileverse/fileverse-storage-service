package server

import (
	"storage/src/services/server/middleware"

	"github.com/gin-gonic/gin"
)

func v1configs(v1 *gin.RouterGroup) {
	v1.Use(middleware.ApiKeyAuth())

}
