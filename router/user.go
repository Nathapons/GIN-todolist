package router

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpUserRouter(app *gin.Engine) {
	router := app.Group("api/v1")

	router.GET("/user", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUser)
	router.POST("/user", controllers.CreateUser)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
	router.GET("/user/activate/:id", controllers.ActiveUser)
}
