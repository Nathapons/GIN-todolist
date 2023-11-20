package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envMap, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	PORT := envMap["PORT"]

	app := gin.New()
	app.Run(":" + PORT)
}
