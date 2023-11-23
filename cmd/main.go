package main

import (
	"fmt"
	"main/config"
	"main/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "main/docs"
)

// @title Blog Service API
// @version 1.0
// @description Blog Service API in Go using GIN framework
// @host 	localhost:8000
// @BasePath /api/v1
func main() {
	envMap, err := godotenv.Read()
	if err != nil {
		panic(fmt.Sprintf("Error : %s", err))
	}

	os.Mkdir("upload", os.ModePerm)

	DB_HOST := envMap["DB_HOST"]
	DB_NAME := envMap["DB_NAME"]
	DB_PORT := envMap["DB_PORT"]
	DB_USER := envMap["DB_USER"]
	DB_PASSWORD := envMap["DB_PASSWORD"]
	PORT := envMap["PORT"]

	app := gin.Default()

	config.SetCors(app)
	config.SetDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	router.SetUpRouter(app)

	app.Run(":" + PORT)
}
