package main

import (
	"main/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "main/docs"
)

// @title Blog todolist API
// @version 1.0
// @description Blog todolist API in Go using GIN framework
// @host 	localhost:8000
// @BasePath /api/v1
func main() {
	envMap, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	PORT := envMap["PORT"]

	app := gin.New()

	router.SetUpRouter(app)

	app.Run(":" + PORT)
}
