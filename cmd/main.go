package main

import (
	"github.com/ashbelLama/brainbox/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/public", "./public")

	router.GET("/", func(ctx *gin.Context) {
		controller.LayoutController(ctx)
	})
	router.GET("/admin", func(ctx *gin.Context) {
		controller.AdminController(ctx)
	})

	router.Run("localhost:3000")
}
