package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/ktakenaka/go-random/app/interface/api/presenter"
)

const (
	metaKey  = "meta"
	errorKey = "error"
	dataKey  = "data"
)

type ResponseFormatter struct{}

func NewResponseFormatter() ResponseFormatter {
	return ResponseFormatter{}
}

func (m *ResponseFormatter) PanicRecovery(ctx *gin.Context) {
	defer func() {
		if p := recover(); p != nil {
			log.Print(p)
			meta := presenter.MetaCodePair[presenter.CodePanic]
			res := presenter.Response{
				Meta: meta,
			}
			ctx.JSON(http.StatusInternalServerError, res)
		}
	}()

	ctx.Next()
}

func (m *ResponseFormatter) Format(ctx *gin.Context) {
	ctx.Next()

	var res presenter.Response
	meta := getMetaResponse(ctx)

	errs := getErrorResponse(ctx)
	if len(errs) > 0 {
		res = presenter.Response{
			Meta:   meta,
			Errors: errs,
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

	ctx.JSON(http.StatusOK, res)
}

func SetMetaResponse(ctx *gin.Context, code presenter.MetaCode) {
	meta := presenter.MetaCodePair[code]
	ctx.Set(metaKey, meta)
}

func SetDataResponse(ctx *gin.Context, data interface{}) {
	ctx.Set(dataKey, data)
}

func SetErrorResponse(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	// TODO: Wrap errors at the place to happen => logging => easy to find the place
	log.Println(err)

	if ve, ok := err.(validator.ValidationErrors); ok {
		errs := make([]presenter.ResponseError, len(ve))

		for i, v := range ve {
			source := presenter.ResponseErrorSource{
				Pointer: v.Field(),
			}
			errs[i] = presenter.ResponseError{
				Source: source,
				Detail: v.Tag(),
			}
		}
		SetMetaResponse(ctx, presenter.CodeUnprocessableEntity)
		ctx.Set(errorKey, errs)
	} else {
		errPrs := presenter.ResponseError{
			Detail: err.Error(),
		}
		SetMetaResponse(ctx, presenter.CodeInternalServerError)
		ctx.Set(errorKey, []presenter.ResponseError{errPrs})
	}
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
