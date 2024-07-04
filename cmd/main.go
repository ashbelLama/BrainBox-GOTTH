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

	router.GET("/admin", func(ctx *gin.Context) {
		controller.AdminGETController(ctx)
	})
	router.POST("/admin", func(ctx *gin.Context) {
		controller.AdminPOSTController(ctx)
	})
	router.GET("/admin/dashboard", func(ctx *gin.Context) {
		controller.AdminDashboardController(ctx)
	})

	v2 := router.Group("/")
	{
		v2.GET("/", func(ctx *gin.Context) {
			controller.HomeController(ctx)
		})
		v2.GET("/bubble", func(ctx *gin.Context) {
			controller.GetBubbleController(ctx)
		})
		v2.POST("/bubble", func(ctx *gin.Context) {
			controller.PostBubbleController(ctx)
		})
	}

	router.Run("localhost:3000")
}
