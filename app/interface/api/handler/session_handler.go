package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/config"
	"github.com/ktakenaka/go-random/app/interface/adaptor/restclient"
	"github.com/ktakenaka/go-random/app/interface/persistence/mysql"
	"github.com/ktakenaka/go-random/app/usecase"
)

func AddSessionHandler(g *gin.RouterGroup) {
	g.POST("/google", createSessionWithGoogle)
}

func createSessionWithGoogle(ctx *gin.Context) {
	cnf := config.GetGoogleOauthConfig()
	gRepo := restclient.NewGoogleRepository(cnf)
	uRepo := mysql.NewUserRepository()
	uc := usecase.NewSignInUsecase(gRepo, uRepo)
	user, err := uc.Execute(ctx.PostForm("code"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"temp": user})
}
