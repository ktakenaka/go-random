package errors

import (
	"fmt"
)

type APIErrorBuilder interface {
	error
	fmt.Stringer
	WithTitle(title string) APIErrorBuilder
	WithDetail(detail string) APIErrorBuilder
	WithSource(pointer string) APIErrorBuilder
	WithLang(lang string) APIErrorBuilder
	Build() APIError
}

type APIError interface {
	error
	Is(err error) bool
	Unwrap() error

	Status() int
	Code() string
	Source() string
	Detail() string
}

// APIErrorDetails 詳細エラー
type APIErrorDetails interface {
	fmt.Stringer
	Resource() string
	Name() string
	Reason() string
}
