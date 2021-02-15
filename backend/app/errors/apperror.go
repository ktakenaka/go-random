package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ktakenaka/go-random/backend/app/errors/translator"
)

// AppError wraps error
type AppError struct {
	JSONAPIError

	err    error                  // expected to set original error
	params map[string]interface{} // Expected userInput params (Not to be translated)
	fields map[string]interface{} // Expected app managed values (To be translated)
	msgLog string                 // a logging message (not for a user message)
}

// NewAppError constructor
func NewAppError(err error) *AppError {
	return &AppError{err: err}
}

// WithParams - params
func (e *AppError) WithParams(params map[string]interface{}) *AppError {
	e.params = params
	return e
}

// WithFields - fields
func (e *AppError) WithFields(fields map[string]interface{}) *AppError {
	e.fields = fields
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

	e.code = e.err.Error()
	e.title = http.StatusText(e.status)
	e.detail = e.buildDetail(lang)
}

// Error - errors.Error() for logging, not for user messages
func (e *AppError) Error() string {
	// This condition is after Build()
	if e.detail != "" {
		return fmt.Sprintf(
			"title: %s, detail: %s, log: %s",
			e.title,
			e.buildDetail(enLang),
			e.msgLog,
		)
	}
	return e.err.Error()
}

// Is - errors.Is
func (e *AppError) Is(err error) bool {
	return errors.Is(e.err, err)
}

// Unwrap - errors.Unwrap
func (e *AppError) Unwrap() error {
	return e.err
}

// detail is for a user message. Shouldn't be a raw error
func (e *AppError) buildDetail(lang string) string {
	data := e.params
	for k, v := range e.fields {
		if vstr, ok := v.(string); ok {
			translated, err := translator.LocalizeField(lang, vstr)
			if err != nil {
				data[k] = vstr
				continue
			}
			data[k] = translated
		}
	}

	arg := translator.Arg{
		Lang:  lang,
		MsgID: e.err.Error(),
		Data:  data,
	}

	msg, err := translator.Localize(arg)
	if err != nil {
		// FIXME:
		// when just forgetting to define i18n messages for the error, it shows "unexpected error".
		return ErrUnknown.Error()
	}
	return msg
}
