package router

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpUserRouter(app *gin.Engine) {
	router := app.Group("api/v1")

	router.GET("/user", controllers.GetUsers)
	router.POST("/user", controllers.CreateUser)
}
