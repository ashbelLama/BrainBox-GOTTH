package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ashbelLama/brainbox/view/admin"
	"github.com/gin-gonic/gin"
)

type content struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AdminGETController(ctx *gin.Context) {
	admin.Home().Render(ctx.Request.Context(), ctx.Writer)
}

func AdminPOSTController(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodPost {
		err := ctx.Request.ParseForm()
		if err != nil {
			http.Error(ctx.Writer, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	username := strings.TrimSpace(ctx.PostForm("username"))
	password := ctx.PostForm("password")

	if verifyLogin(username, password) {
		ctx.Header("HX-Redirect", "/admin/dashboard")
	} else {
		fmt.Println("invalid username or password")
	}

}

func verifyLogin(username, password string) bool {
	if username == "admin" && password == "password" {
		return true
	}
	return false
}
