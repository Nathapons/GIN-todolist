package router

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(app *gin.Engine) {
	router := app.Group("api/v1")

	router.GET("/blog", controllers.GetBlogs)
}
