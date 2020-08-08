package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/registry"
)

func AddSessionHandler(g *gin.RouterGroup) {
	g.POST("/google", createSessionWithGoogle)
}

func createSessionWithGoogle(ctx *gin.Context) {
	uc := registry.InitializeSignInUsecase()
	user, err := uc.Execute(ctx.PostForm("code"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"temp": user})
}
