package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/backend/app/interface/api/middleware"
	"github.com/ktakenaka/go-random/backend/app/interface/api/presenter"
	"github.com/ktakenaka/go-random/backend/pkg/logger"
)

// BaseHandler is expected to be embedded
type BaseHandler struct{}

// SetData is expected to use concrete handlers
func (hdl *BaseHandler) SetData(ctx *gin.Context, data interface{}) {
	middleware.SetDataResponse(ctx, data)
}

// SetMeta is expected to be used concrete handlers
func (hdl *BaseHandler) SetMeta(ctx *gin.Context, code presenter.MetaCode) {
	middleware.SetMetaResponse(ctx, code)
}

// SetError is expected to be used in concrete handlers
func (hdl *BaseHandler) SetError(ctx *gin.Context, err error) {
	if err != nil {
		logger.Error(err)
	}
	middleware.SetErrorResponse(ctx, err)
}

// JWTClaims extracts claims from JWT
func (hdl *BaseHandler) JWTClaims(ctx *gin.Context) middleware.JWTClaims {
	claims, _ := middleware.ExtractClaims(ctx)
	return claims
}
