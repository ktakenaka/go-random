package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"

	validator "github.com/go-playground/validator/v10"
	"github.com/ktakenaka/go-random/backend/app/domain/entity"
	appErr "github.com/ktakenaka/go-random/backend/app/errors"
)

var (
	errKey = "error"
)

func main() {
	fmt.Println("--- pattern 1 AppError ---")
	ctx1 := buildContext()
	handlerFunc1(ctx1)
	middlewareErrFunc(ctx1)

	fmt.Println("\n--- pattern 2 validation error ---")
	ctx2 := buildContext()
	handlerFunc2(ctx2)
	middlewareErrFunc(ctx2)

	fmt.Println("\n--- pattern 3 unhandled error ---")
	ctx3 := buildContext()
	handlerFunc3(ctx3)
	middlewareErrFunc(ctx3)
}

func buildContext() *gin.Context {
	return &gin.Context{
		Request: &http.Request{
			Method:     "POST",
			URL:        &url.URL{Path: "/example"},
			Body:       ioutil.NopCloser(bytes.NewBufferString("body is here")),
			Header:     http.Header{"Accept-Language": []string{"ja"}},
			RemoteAddr: "127.0.0.1",
		},
	}
}

func handlerFunc1(ctx *gin.Context) {
	err := functionWithAppError("hello")
	if err != nil {
		ctx.Set(errKey, err)
	}
}

func handlerFunc2(ctx *gin.Context) {
	err := functionWithValidationError()
	if err != nil {
		ctx.Set(errKey, err)
	}
}

func handlerFunc3(ctx *gin.Context) {
	err := functionWithUnknownError()
	if err != nil {
		ctx.Set(errKey, err)
	}
}

func functionWithAppError(arg string) error {
	err := appErr.NewAppError(appErr.ErrExample).
		WithMsgLog(fmt.Sprintf("a user passes an invalid arg: %s", arg)).
		WithParams(map[string]interface{}{"Name": arg}).
		WithFields(map[string]interface{}{"Field": "Sample"})
	return xerrors.Errorf("%w", err)
}

func functionWithValidationError() error {
	en := entity.Sample{Title: "a"}
	return en.Validate()
}

func functionWithUnknownError() error {
	err := errors.New("unknow error")
	err = appErr.NewAppError(err).WithMsgLog("not expected place")
	return xerrors.Errorf("%w", err)
}

func middlewareErrFunc(ctx *gin.Context) {
	errInterface, ok := ctx.Get(errKey)
	if !ok {
		// No errors
		return
	}

	lang := ctx.GetHeader("Accept-Language")

	if ve, ok := errInterface.(validator.ValidationErrors); ok {
		vErrs := appErr.NewValidationErrors(ve)
		vErrs.Build(lang)
		newLogger().WithRequest(ctx.Request).Info(vErrs)
		return
	}

	err, ok := errInterface.(error)
	if !ok {
		newLogger().WithRequest(ctx.Request).Error("failed to convert interface to error")
		return
	}

	var apperr *appErr.AppError
	if ok := errors.As(err, &apperr); ok {
		apperr.Build(lang)
		switch {
		case errors.Is(apperr, appErr.ErrExample):
			newLogger().WithRequest(ctx.Request).Info(err)
		default:
			newLogger().WithRequest(ctx.Request).Error(err)
		}
		return
	}

	// Uncought error
	newLogger().WithRequest(ctx.Request).Error(err)
}

// --- logger example ---
type logger struct {
	req *http.Request
}

func newLogger() *logger {
	return &logger{}
}

func (l *logger) WithRequest(req *http.Request) *logger {
	l.req = req
	return l
}

func (l *logger) Info(v interface{}) {
	l.PrintReq()
	fmt.Printf("Info: %+v\n", v)
}

func (l *logger) Warn(v interface{}) {
	l.PrintReq()
	fmt.Printf("Warn: %+v\n", v)
}

func (l *logger) Error(v interface{}) {
	l.PrintReq()
	fmt.Printf("Error: %+v\n", v)
}

func (l *logger) PrintReq() {
	if l.req != nil {
		fmt.Printf("req: %s %s %s %s %s\n", l.req.Method, l.req.Header, l.req.URL, l.req.Body, l.req.RemoteAddr)
	}
}
