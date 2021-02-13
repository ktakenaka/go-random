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

	appErr "github.com/ktakenaka/go-random/backend/app/errors"
)

var authValue = struct {
	UserID, OfficeID uint32
}{1, 10}

var (
	authKey = "auth"
	errKey  = "error"
)

func main() {
	ctx := &gin.Context{
		Request: &http.Request{
			Method:     "POST",
			URL:        &url.URL{Path: "/example"},
			Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
			Header:     http.Header{"Accept-Language": []string{"ja"}},
			RemoteAddr: "127.0.0.1",
		},
	}
	ctx.Set(authKey, authValue)

	handlerFunc(ctx)
	middlewareErrFunc(ctx)
}

func handlerFunc(ctx *gin.Context) {
	err := functionWithError("hello")
	if err != nil {
		ctx.Set(errKey, err)
	}
}

func middlewareErrFunc(ctx *gin.Context) {
	errInterface, ok := ctx.Get(errKey)
	if !ok {
		// No errors
		return
	}

	err, ok := errInterface.(error)
	if !ok {
		newLogger().WithRequest(ctx.Request).Error("failed to convert interface to error")
		return
	}

	lang := ctx.GetHeader("Accept-Language")
	if errors.Is(err, appErr.ErrExample) {
		var apperr *appErr.AppError
		if ok := errors.As(err, &apperr); !ok {
			newLogger().Error("failed to convert")
		}
		apperr.Build(lang)
		switch {
		case errors.Is(apperr, appErr.ErrExample):
			newLogger().WithRequest(ctx.Request).Info(err)
		default:
			newLogger().WithRequest(ctx.Request).Error(err)
		}
		return
	}

	newLogger().WithRequest(ctx.Request).Error(err)
}

func functionWithError(arg string) error {
	err := appErr.New(appErr.ErrExample).
		WithMsgLog(fmt.Sprintf("a user passes an invalid arg: %s", arg)).
		WithParams(map[string]interface{}{"Name": arg}) //  TODO: ここでField名をi18nで使いたいときどうする？
	return xerrors.Errorf("%w", err)
}

func functionWithValidationError() error {
	return nil
}

func functionWithUnknownError() error {
	return nil
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
