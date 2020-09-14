package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/interface/api/middleware"
	"github.com/ktakenaka/go-random/app/interface/api/presenter"
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
	middleware.SetErrorResponse(ctx, err)
}
