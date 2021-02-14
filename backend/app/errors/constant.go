package errors

import (
	"errors"
	"net/http"
)

// Application errors
var (
	ErrExample    = errors.New("ERR_EXAMPLE")
	ErrValidation = errors.New("ERR_VALIDATION")
)
var errorStatusMap = map[error]int{
	ErrExample: http.StatusBadRequest,
}

const (
	// use English for logging
	enLang = "en"
)
