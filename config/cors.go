package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetCors(app *gin.Engine) {
	app.Use(cors.Default())
}
