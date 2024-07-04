package controller

import (
	"github.com/ashbelLama/brainbox/view/home"
	"github.com/gin-gonic/gin"
)

func HomeController(ctx *gin.Context) {
	home.Home().Render(ctx.Request.Context(), ctx.Writer)
}
