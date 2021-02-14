package errors

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"

	validator "github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"

	"github.com/ktakenaka/go-random/backend/app/errors/translator"
)

// ValidationError wraps error
type ValidationError struct {
	JSONAPIError
	err validator.FieldError
}

// ValidationErrors bunch of ValidationError
type ValidationErrors []*ValidationError

// NewValidationErrors constructor
func NewValidationErrors(ve validator.ValidationErrors) ValidationErrors {
	errs := make([]*ValidationError, len(ve))
	for i := range ve {
		errs[i] = newValidationError(ve[i])
	}
	return errs
}

// Build - builder (status, code, title, detail)
func (e ValidationErrors) Build(lang string) {
	for i := range e {
		e[i].build(lang)
	}
}

func (e ValidationErrors) Error() string {
	buff := bytes.NewBufferString("")
	for i := range e {
		buff.WriteString(e[i].buildDetail(enLang))
		buff.WriteString("\n")
	}
	return strings.TrimSpace(buff.String())
}

func newValidationError(err validator.FieldError) *ValidationError {
	return &ValidationError{err: err}
}

func (e *ValidationError) build(lang string) {
	e.status = http.StatusUnprocessableEntity
	e.code = ErrValidation.Error()
	e.source = e.buildPointer()
	e.title = http.StatusText(e.status)
	e.detail = e.buildDetail(lang)
}

// Error - errors.Error() for logging, not for user messages
func (e *ValidationError) Error() string {
	// This condition is after Build()
	if e.detail != "" {
		return fmt.Sprintf("title: %s, detail: %s", e.title, e.buildDetail(enLang))
	}
	return e.err.Error()
}

// Is - errors.Is
func (e *ValidationError) Is(err error) bool {
	return errors.Is(e.err, err)
}

// Unwrap - errors.Unwrap
func (e *ValidationError) Unwrap() error {
	return e.err
}

func (e *ValidationError) buildDetail(lang string) string {
	field, err := translator.LocalizeField(lang, e.err.StructNamespace())
	if err != nil {
		field = e.err.Field()
	}
	arg := translator.Arg{
		Lang:  lang,
		MsgID: e.err.Tag(),
		Data: map[string]interface{}{
			"Value": e.err.Value(),
			"Param": e.err.Param(),
			"Field": field,
		},
	}
	msg, err := translator.Localize(arg)
	if err != nil {
		return e.err.Error()
	}
	return msg
}

func (e *ValidationError) buildPointer() string {
	str := strings.ReplaceAll(e.err.StructNamespace(), ".", "/")
	return strcase.ToSnake(str)
}
