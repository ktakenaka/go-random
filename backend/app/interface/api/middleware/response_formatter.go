package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/ktakenaka/go-random/backend/app/interface/api/presenter"
)

const (
	metaKey  = "meta"
	errorKey = "error"
	dataKey  = "data"
)

// ResponseFormatter format response
type ResponseFormatter struct{}

// NewResponseFormatter constructor
func NewResponseFormatter() ResponseFormatter {
	return ResponseFormatter{}
}

// PanicRecovery when panic occurs, it format response
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

// Format the logic to format response
func (m *ResponseFormatter) Format(ctx *gin.Context) {
	ctx.Next()

	// TODO: refactor. It's not good to depend on Handler
	if ok, _ := regexp.MatchString(
		`.*(text/csv|application/octet-stream).*`,
		ctx.Writer.Header().Get("Content-Type"),
	); ok {
		return
	}

	var res presenter.Response
	meta := getMetaResponse(ctx)

	errs := getErrorResponse(ctx)
	if len(errs) > 0 {
		res = presenter.Response{
			Meta:   meta,
			Errors: errs,
		}
	} else {
		data := getDataResponse(ctx)
		res = presenter.Response{
			Meta: meta,
			Data: data,
		}
	}

	code := int(meta.Code)
	httpStatus, _ := strconv.ParseInt(strconv.Itoa(code)[:3], 10, 64)
	ctx.JSON(int(httpStatus), res)
}

// SetMetaResponse set response
func SetMetaResponse(ctx *gin.Context, code presenter.MetaCode) {
	meta := presenter.MetaCodePair[code]
	ctx.Set(metaKey, meta)
}

// SetDataResponse data response
func SetDataResponse(ctx *gin.Context, data interface{}) {
	ctx.Set(dataKey, data)
}

// SetErrorResponse err response
func SetErrorResponse(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	// TODO: Wrap errors at the place to happen => logging => easy to find the place
	log.Println(err)

	// It's difficult to set meta in handler because it's not clear until we accept errors.
	// Especially the errors from packages.
	// That's the reason to judge error type and set meta here
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
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		params := make([]string, len(ctx.Params))
		for i, p := range ctx.Params {
			params[i] = fmt.Sprintf("%v: %v", p.Key, p.Value)
		}

		source := presenter.ResponseErrorSource{
			Param: strings.Join(params, ", "),
		}
		SetMetaResponse(ctx, presenter.CodeNotFound)
		errs := []presenter.ResponseError{
			{
				Source: source,
				Detail: err.Error(),
			},
		}
		ctx.Set(errorKey, errs)
	} else {
		errs := presenter.ResponseError{
			Detail: err.Error(),
		}

		if _, ok := ctx.Get(metaKey); !ok {
			SetMetaResponse(ctx, presenter.CodeInternalServerError)
			ctx.Set(errorKey, []presenter.ResponseError{errs})
		}
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
