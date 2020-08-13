package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/registry"
	"github.com/ktakenaka/go-random/helper/jwtutil"
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

	claims := jwtutil.AuthClaims{
		UserID: user.ID, //TODO: make user_id hash
	}

	aTkn, rTkn, csrfTkn, err := jwtutil.GenerateToken(&claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	ctx.SetCookie("aa_jwt_access", aTkn, 86400, "/api", "", false, true)
	ctx.SetCookie("aa_jwt_refresh", rTkn, 86400, "/api", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"data": csrfTkn})
}
