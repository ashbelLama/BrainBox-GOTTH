package controller

import (
	"github.com/ashbelLama/brainbox/view/layout"
	"github.com/gin-gonic/gin"
)

func LayoutController(ctx *gin.Context) {
	layout.Base().Render(ctx.Request.Context(), ctx.Writer)
}
