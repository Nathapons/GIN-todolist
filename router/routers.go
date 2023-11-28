package router

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter(app *gin.Engine) {
	SetUpBlogRouter(app)
	SetUpSwaggerRouter(app)
	SetUpUserRouter(app)
	SetupAuthRouter(app)
}
