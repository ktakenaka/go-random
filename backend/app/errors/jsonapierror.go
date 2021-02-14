package errors

// JSONAPIError - ref: https://jsonapi.org/examples/#error-objects-basics
/*
{
  "errors": [
    {
      "status": "422",
      "source": { "pointer": "/data/attributes/firstName" },
      "title":  "Invalid Attribute",
      "detail": "First name must contain at least three characters."
    }
  ]
}
*/
type JSONAPIError struct {
	status int    // 基本は err と http.StatusXXX の連想配列を使うようにする。発生箇所での上書きも可能にする
	code   string // titleとの使い分けに悩むな、なしでも良いかも. それか別で errとcodeの連想配列定義するか
	source string // こいつの使い方な...
	title  string // http.StatusText(status) で出てきたやつ使えば良いかも
	detail string // defaultでは err と i18n の組を使う（発生箇所で上書きできるようにしよかと思ったけど、i18n考えると工夫が必要）
}

// WithSource - source
func (e *JSONAPIError) WithSource(source string) {
	e.source = source
}

// Status - status
func (e *JSONAPIError) Status() int {
	return e.status
}

// Code - code
func (e *JSONAPIError) Code() string {
	return e.code
}

// Source - source
func (e *JSONAPIError) Source() string {
	return e.source
}

// Title - title
func (e *JSONAPIError) Title() string {
	return e.title
}

// Detail - detail
func (e *JSONAPIError) Detail() string {
	return e.detail
}
