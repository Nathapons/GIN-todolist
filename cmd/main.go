package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.New()
	app.Run(":8000")
}