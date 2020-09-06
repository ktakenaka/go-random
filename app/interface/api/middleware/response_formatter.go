package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ktakenaka/go-random/app/interface/api/presenter"
)

const (
	metaKey  = "meta"
	errorKey = "error"
	dataKey  = "data"
)

// the method to set Error
// the method to set Meta
// the method to set data

/*
{
  "meta": {
    "code": 20001,
    "message": "success"
  },
  "data": {
    "id": 1,
    "title": "sample"
  }
}
*/

type ResponseFormatter struct{}

func NewResponseFormatter() ResponseFormatter {
	return ResponseFormatter{}
}

func (m *ResponseFormatter) Format(ctx *gin.Context) {
	ctx.Next()

	var res presenter.Response

	meta := getMetaResponse(ctx)
	errors := getErrorResponse(ctx)
	if len(errors) > 0 {
		res = presenter.Response{
			Meta:   meta,
			Errors: errors,
		}
		// TODO: change status from context
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := getDataResponse(ctx)
	res = presenter.Response{
		Meta: meta,
		Data: data,
	}

	// TODO: change status from context
	ctx.JSON(http.StatusOK, res)
}

func SetMetaResponse(ctx *gin.Context, meta presenter.ResponseMeta) {
	ctx.Set(metaKey, meta)
}

func SetDataResponse(ctx *gin.Context, data interface{}) {
	ctx.Set(dataKey, data)
}

func SetErrorResponse(ctx *gin.Context, err presenter.ResponseError) {
	errors, ok := ctx.Get(errorKey)
	var errs []presenter.ResponseError

	if !ok {
		errs = make([]presenter.ResponseError, 1)
		errs[0] = err
	} else {
		errs = append(errors.([]presenter.ResponseError), err)
	}

	ctx.Set(errorKey, errs)
}

func getMetaResponse(ctx *gin.Context) presenter.ResponseMeta {
	meta, ok := ctx.Get(metaKey)
	if !ok {
		return presenter.ResponseMeta{}
	}

	return meta.(presenter.ResponseMeta)
}

func getDataResponse(ctx *gin.Context) interface{} {
	data, ok := ctx.Get(dataKey)
	if !ok {
		return ""
	}

	return data
}

func getErrorResponse(ctx *gin.Context) []presenter.ResponseError {
	var errors []presenter.ResponseError

	errCtx, ok := ctx.Get(errorKey)
	if !ok {
		return errors
	}

	for _, err := range errCtx.([]interface{}) {
		errors = append(errors, err.(presenter.ResponseError))
	}

	return errors
}
