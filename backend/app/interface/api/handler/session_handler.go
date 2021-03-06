package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/backend/app/interface/api/middleware"
	"github.com/ktakenaka/go-random/backend/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/backend/app/registry"
)

type SessionHandler struct{}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{}
}

func (hdl *SessionHandler) CreateWithGoogle(ctx *gin.Context) {
	uc := registry.InitializeSignInUsecase()

	var err error
	defer func() {
		middleware.SetErrorResponse(ctx, err)
	}()

	var req presenter.GoogleSessionRequest
	if err = ctx.ShouldBind(&req); err != nil {
		return
	}
	aTkn, rTkn, csrfTkn, err := uc.Execute(req.Code, req.Nonce)
	if err != nil {
		return
	}

	middleware.SetAccessCookie(ctx, aTkn)
	middleware.SetRefreshCookie(ctx, rTkn)
	middleware.SetMetaResponse(ctx, presenter.CodeCreated)
	middleware.SetDataResponse(
		ctx,
		presenter.SessionResponse{
			CSRFToken: csrfTkn,
		},
	)
}
