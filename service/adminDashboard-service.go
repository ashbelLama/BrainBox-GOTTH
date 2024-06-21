package service

import (
	"log"

	"github.com/gin-gonic/gin"
)

type E struct {
	Events string
}

func AdminPostService(ctx *gin.Context) {
	// jsonData, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("%s", string(jsonData))
	// data := &E{}
	// ctx.Bind(data) // data is left unchanged because c.Request.Body has been used up.
	// fmt.Println(data)
	// ctx.JSON(http.StatusOK, c)
	log.Println("admin post service")
}
