package middleware

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func CORSMiddleware(app *gin.Engine, allowedOrigins string) {
	config := cors.Config{
		Origins:        allowedOrigins,
		Methods:        "GET, POST, PUT, DELETE, OPTIONS",
		RequestHeaders: "Origin, Authorization, Content-Type",
		MaxAge:         50 * 60, // 50 minutes
	}
	app.Use(cors.Middleware(config))
}
