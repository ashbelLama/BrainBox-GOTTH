package main

import (
	"github.com/ashbelLama/brainbox/controller"
	database "github.com/ashbelLama/brainbox/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.InitDB()

	router.Static("/public", "./public")

	v1 := router.Group("/")
	{
		v1.GET("/", func(ctx *gin.Context) {
			controller.HomeController(ctx)
		})
		v1.GET("/bubble", func(ctx *gin.Context) {
			controller.GetBubbleController(ctx)
		})
		v1.POST("/bubble", func(ctx *gin.Context) {
			controller.PostBubbleController(ctx)
		})
		v1.GET("/form", func(ctx *gin.Context) {
			controller.FormController(ctx)
		})
	}

	router.Run("localhost:3000")
}
