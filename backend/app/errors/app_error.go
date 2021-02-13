package errors

import (
	"errors"
	"net/http"

	"github.com/ktakenaka/go-random/backend/app/errors/translator"
)

// AppError wraps error
type AppError struct {
	JSONAPIError

	err    error
	params map[string]interface{} // TODO: Railsのerror参考にする (定義してあるmessageの置き換え文字との整合性が必要)
	msgLog string
}

// New constructor
func New(err error) *AppError {
	return &AppError{err: err}
}

// WithParams - params
func (e *AppError) WithParams(params map[string]interface{}) *AppError {
	e.params = params
	return e
}

// WithMsgLog - msgLog
func (e *AppError) WithMsgLog(msg string) *AppError {
	e.msgLog = msg
	return e
}

// Build - builder (status, code, title, detail)
func (e *AppError) Build(lang string) {
	if status, ok := errorStatusMap[e.err]; !ok {
		e.status = http.StatusInternalServerError
	} else {
		e.status = status
	}

	if code, ok := errorJSONAPICodeMap[e.err]; !ok {
		e.code = codeUnhandled
	} else {
		e.code = code
	}

	e.title = http.StatusText(e.status)
	e.detail = e.userMessage(lang)
}

// Error - errors.Error()
func (e *AppError) Error() string {
	if e.msgLog != "" {
		return e.msgLog
	}
	// Error() is for logging, that's the reason to use English
	return e.userMessage("en")
}

// Is - errors.Is
func (e *AppError) Is(err error) bool {
	return errors.Is(e.err, err)
}

// Unwrap - errors.Unwrap
func (e *AppError) Unwrap() error {
	return e.err
}

func (e *AppError) userMessage(lang string) string {
	// TODO: handle error
	msg, _ := translator.Localize(lang, e.err.Error(), e.params)
	return msg
}
