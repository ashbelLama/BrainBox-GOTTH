package main

import (
	"github.com/ashbelLama/brainbox/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		controller.LayoutController(ctx)
	})

	router.Run("localhost:3000")
}
