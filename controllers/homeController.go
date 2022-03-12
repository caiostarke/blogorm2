package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
	})
}
