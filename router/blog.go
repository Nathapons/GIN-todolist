package router

import (
	v1 "main/controllers/v1"

	"github.com/gin-gonic/gin"
)

func SetUpBlogRouter(app *gin.Engine) {
	router := app.Group("api/v1")

	router.GET("/blog", v1.GetBlogs)
	router.POST("/blog", v1.CreateBlogs)
	router.PUT("/blog/:id", v1.UpdateBlogs)
	router.DELETE("blog/:id", v1.DeletBlogs)
}
