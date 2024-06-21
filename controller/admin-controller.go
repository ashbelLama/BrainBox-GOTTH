package controller

import (
	"github.com/ashbelLama/brainbox/view/admin"
	"github.com/gin-gonic/gin"
)

func AdminController(ctx *gin.Context) {
	admin.Home().Render(ctx.Request.Context(), ctx.Writer)
}

// func writeImage(ctx *gin.Context, img *image.Image) {
// 	buffer := new(bytes.Buffer)
// 	if err := jpeg.Encode(buffer, *img, nil); err != nil {
// 		log.Println("unable to encode image")
// 	}

// 	ctx.Header().Set("Content-Type", "image/jpeg")
// }
