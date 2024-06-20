package controller

import (
	"github.com/ashbelLama/brainbox/view/css/layout"
	"github.com/gin-gonic/gin"
)

func LayoutController(ctx *gin.Context) {
	layout.Hello().Render(ctx.Request.Context(), ctx.Writer)
}
