package router

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(app *gin.Engine) {
	router := app.Group("api/v1")
	router.POST("/login", controllers.Login)
}