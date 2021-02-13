package errors

import (
	"errors"
	"net/http"
)

// Application errors
var (
	ErrExample = errors.New("ERR_EXAMPLE")
)
var errorStatusMap = map[error]int{
	ErrExample: http.StatusBadRequest,
}

type jsonAPICode string

const (
	codeExample   jsonAPICode = "4001"
	codeUnhandled jsonAPICode = "5999"
)

var errorJSONAPICodeMap = map[error]jsonAPICode{
	ErrExample: codeExample,
}
