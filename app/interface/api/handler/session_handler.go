package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/interface/api/middleware"
	"github.com/ktakenaka/go-random/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/app/registry"
)

func AddSessionHandler(g *gin.RouterGroup) {
	g.POST("/google", createSessionWithGoogle)
}

func createSessionWithGoogle(ctx *gin.Context) {
	uc := registry.InitializeSignInUsecase()

	var req presenter.GoogleSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}
	aTkn, rTkn, csrfTkn, err := uc.Execute(req.Code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	middleware.SetAccessCookie(ctx, aTkn)
	middleware.SetRefreshCookie(ctx, rTkn)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": presenter.SessionResponse{CSRFToken: csrfTkn},
		},
	)
}
