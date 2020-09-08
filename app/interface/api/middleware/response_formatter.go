package middleware

import (
	"github.com/go-playground/validator/v10"
	"log"
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
// meta := presenter.ResponseMeta{
// 	Code:    400,
// 	Message: "failure",
// }

// err := presenter.ResponseError{
// 	Detail: err.Error(),
// }

type ResponseFormatter struct{}

func NewResponseFormatter() ResponseFormatter {
	return ResponseFormatter{}
}

func (m *ResponseFormatter) Format(ctx *gin.Context) {
	ctx.Next()

	var res presenter.Response
	var meta presenter.ResponseMeta

	errors := getErrorResponse(ctx)
	if len(errors) > 0 {
		meta = getMetaResponse(ctx)
		res = presenter.Response{
			Meta:   meta,
			Errors: errors,
		}
		// TODO: change status from context
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data := getDataResponse(ctx)
	meta = presenter.ResponseMeta{
		Code:    200,
		Message: "success",
	}
	res = presenter.Response{
		Meta: meta,
		Data: data,
	}

	ctx.JSON(http.StatusOK, res)
}

func SetMetaResponse(ctx *gin.Context, meta presenter.ResponseMeta) {
	ctx.Set(metaKey, meta)
}

func SetDataResponse(ctx *gin.Context, data interface{}) {
	ctx.Set(dataKey, data)
}

func SetError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	// TODO: Wrap errors at the place to happen => logging => easy to find the place
	if ve, ok := err.(validator.ValidationErrors); ok {
		log.Println(ve)
		for k, v := range ve {
			log.Println("===")
			log.Println(k)
			log.Println(v)
			log.Println("===")
		}
	}

	errPrs := presenter.ResponseError{
		Detail: err.Error(),
	}

	meta := presenter.ResponseMeta{
		Code:    500,
		Message: "failure!",
	}

	ctx.Set(metaKey, meta)
	ctx.Set(errorKey, []presenter.ResponseError{errPrs})
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
	errCtx, ok := ctx.Get(errorKey)
	if !ok {
		return []presenter.ResponseError{}
	}

	errs, ok := errCtx.([]presenter.ResponseError)
	if !ok {
		return []presenter.ResponseError{}
	}

	return errs
}
