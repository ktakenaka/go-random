package errors

import (
	"fmt"
)

// ErrorWrapper error wrapper
type ErrorWrapper struct {
	Message       string
	OriginalError error
}

func (e ErrorWrapper) Error() string {
	return e.Message
}

// Wrap wrap
func Wrap(message string, err error) ErrorWrapper {
	wrapper, ok := err.(ErrorWrapper)
	if ok {
		wrapper.Message = fmt.Sprintf("%s -- %s", message, wrapper.Message)
		return wrapper
	}
	return ErrorWrapper{Message: err.Error(), OriginalError: err}
}
