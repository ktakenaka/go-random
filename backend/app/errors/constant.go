package errors

import (
	"errors"
	"net/http"
)

// Application errors
var (
	ErrExample    = errors.New("ERR_EXAMPLE")
	ErrValidation = errors.New("ERR_VALIDATION")
	ErrUnknown    = errors.New("ERR_UNKNOWN")
)
var errorStatusMap = map[error]int{
	ErrExample:    http.StatusBadRequest,
	ErrValidation: http.StatusUnprocessableEntity,
	ErrUnknown:    http.StatusInternalServerError,
}

const (
	// use English for logging
	enLang = "en"
)
