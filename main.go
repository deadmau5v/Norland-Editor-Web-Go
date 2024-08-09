package main

import (
	"norland-config/apis"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Next()
	})

	api := app.Group("api")
	api.GET("/default", apis.GetDefaultGameData)

	app.Run(":8080")
}
