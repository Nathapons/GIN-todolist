package router

import (
	v1 "main/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetUpUserRouter(app *gin.Engine) {
	router := app.Group("api/v1")

	router.GET("/user", v1.GetUsers)
	router.GET("/user/:id", v1.GetUser)
	router.POST("/user", v1.CreateUser)
	router.PUT("/user/:id", v1.UpdateUser)
	router.DELETE("/user/:id", v1.DeleteUser)
	router.GET("/user/activate/:id", v1.ActiveUser)
}
