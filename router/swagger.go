package router

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpSwaggerRouter(app *gin.Engine) {
	app.GET("api_doc/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
