package middleware

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"

	appErr "github.com/ktakenaka/go-random/backend/app/errors"
	"github.com/ktakenaka/go-random/backend/app/interface/api/presenter"
	log "github.com/ktakenaka/go-random/backend/pkg/logger"
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

	lang := ctx.GetHeader("Accept-Language")
	if ve, ok := err.(validator.ValidationErrors); ok {
		vErrs := appErr.NewValidationErrors(ve)
		vErrs.Build(lang)
		errs := make([]presenter.JSONAPIError, len(vErrs))
		for i := range vErrs {
			errs[i] = presenter.JSONAPIError{
				Status: vErrs[i].Status(),
				Code:   vErrs[i].Code(),
				Source: vErrs[i].Source(),
				Title:  vErrs[i].Title(),
				Detail: vErrs[i].Detail(),
			}
		}
		SetMetaResponse(ctx, presenter.CodeUnprocessableEntity)
		ctx.Set(errorKey, errs)
		log.WithRequest(ctx.Request).Info(vErrs)
		return
	}

	var apperr *appErr.AppError
	if ok := errors.As(err, &apperr); ok {
		apperr.Build(lang)
		switch {
		case errors.Is(apperr, appErr.ErrExample):
			e := presenter.JSONAPIError{
				Status: apperr.Status(),
				Code:   apperr.Code(),
				Source: apperr.Source(),
				Title:  apperr.Title(),
				Detail: apperr.Detail(),
			}
			SetMetaResponse(ctx, presenter.CodeBadRequest)
			ctx.Set(errorKey, []presenter.JSONAPIError{e})
			log.WithRequest(ctx.Request).Warn(apperr)
		default:
			log.WithRequest(ctx.Request).Error(apperr)
		}
		return
	}

	// TODO: handle unhandled errors
	log.WithRequest(ctx.Request).Error(err)
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

func getErrorResponse(ctx *gin.Context) []presenter.JSONAPIError {
	errCtx, ok := ctx.Get(errorKey)
	if !ok {
		return []presenter.JSONAPIError{}
	}

	errs, ok := errCtx.([]presenter.JSONAPIError)
	if !ok {
		return []presenter.JSONAPIError{}
	}

	return errs
}
