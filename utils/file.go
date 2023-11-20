package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context) string {
	file, _ := c.FormFile("file")
	newUUID := uuid.New().String()
	dir, _ := os.Getwd()

	filePath := fmt.Sprintf("%s/upload/%s/%s", dir, strings.Replace(newUUID, "-", "", -1), file.Filename)
	c.SaveUploadedFile(file, filePath)

	return filePath
}

func RemoveFile(filePath string) {
	os.Remove(filePath)
}
