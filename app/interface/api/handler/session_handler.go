package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/interface/api/middleware"
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

	csrfTkn, err := jwtutil.GenerateCSRFToken()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	claims := jwtutil.AuthClaims{
		UserID:    user.ID, //TODO: make user_id hash
		IssueTime: time.Now(),
		CSRFToken: csrfTkn,
	}

	aTkn, rTkn, csrfTkn, err := jwtutil.GenerateToken(&claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	middleware.SetAccessCookie(ctx, aTkn)
	middleware.SetRefreshCookie(ctx, rTkn)

	ctx.JSON(http.StatusOK, gin.H{"data": csrfTkn})
}
