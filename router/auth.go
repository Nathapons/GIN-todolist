package router

import (
	v1 "main/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(app *gin.Engine) {
	router := app.Group("api/v1")
	router.POST("/login", v1.Login)
}