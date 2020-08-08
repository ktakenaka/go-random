package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/config"
)

func AddSessionHandler(g *gin.RouterGroup) {
	g.POST("/sessions/google", createGoogle)
}

func createGoogle(ctx *gin.Context) {
	cnf := config.GetGoogleOauthConfig()
	fmt.Println(cnf)
	ctx.JSON(http.StatusOK, gin.H{"temp": "hoge"})
}
