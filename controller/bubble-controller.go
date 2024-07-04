package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ashbelLama/brainbox/model"
	"github.com/ashbelLama/brainbox/service"
	"github.com/gin-gonic/gin"
)

func GetBubbleController(ctx *gin.Context) {
	jsonData, err := service.GetBubble()
	if err != nil {
		log.Fatal(err)
	}
	ctx.String(http.StatusOK, string(jsonData))
}

func PostBubbleController(ctx *gin.Context) {
	var p model.Bubble
	err := json.NewDecoder(ctx.Request.Body).Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	if service.PostBubble(p) {
		ctx.String(http.StatusOK, "Stored successfully")
		return
	}
	ctx.String(http.StatusBadRequest, "Bad request")
}
