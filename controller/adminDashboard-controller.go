package controller

import (
	"github.com/ashbelLama/brainbox/view/admin"
	"github.com/gin-gonic/gin"
)

func AdminDashboardController(ctx *gin.Context) {
	admin.Dashboard().Render(ctx.Request.Context(), ctx.Writer)
}
