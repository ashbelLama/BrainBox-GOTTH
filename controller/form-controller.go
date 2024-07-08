package controller

import (
	"github.com/ashbelLama/brainbox/view/form"
	"github.com/gin-gonic/gin"
)

func FormController(ctx *gin.Context) {
	form.Form().Render(ctx.Request.Context(), ctx.Writer)
}
