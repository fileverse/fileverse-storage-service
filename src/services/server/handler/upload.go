package handler

import (
	"net/http"
	"storage/src/pkg/logger"
	"storage/src/services/upload"

	"github.com/gin-gonic/gin"
)

// UploadFile is a handler that uploads a file to Pinatas
func UploadFile(c *gin.Context) {
	log := logger.GetContextLogger(c)
	file, err := c.FormFile("file")
	tags, _ := c.GetQueryArray("tags")

	if err != nil {
		log.Error("invalid file", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	canUpload, err := upload.CanUpload(c, file.Size)
	if err != nil || !canUpload {
		log.Error("failed to upload file", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Upload the file to Pinata
	resp, err := upload.UploadFile(c, file, tags)
	if err != nil {
		log.Error("failed to upload file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UploadIdentityFile(c *gin.Context) {
	log := logger.GetContextLogger(c)

	file, err := c.FormFile("file")
	if err != nil {
		log.Error("invalid file", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Upload the file to Pinata
	resp, err := upload.UploadIdentityFile(c, file)
	if err != nil {
		log.Error("failed to upload file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
