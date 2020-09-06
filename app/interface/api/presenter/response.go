package presenter

// ResponseMeta is Meta Object of JSON API
// Every response includes Meta even if it's the error response
type ResponseMeta struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

// ResponseError provides additional information about problems encountered while performing an operation.
// It will be included in the array, key will be `errors`
type ResponseError struct {
	Source ResponseErrorSource `json:"source"`
	Detail string              `json:"detail"`
}

// ResponseErrorSource is an object containing references to the source of the error,
// optionally including any of the following members:
type ResponseErrorSource struct {
	Pointer string `json:"pointer,omitempty"`
	Param   string `json:"param,omitempty"`
}

// Response is the Response format of JSON API
type Response struct {
	Meta   ResponseMeta    `json:"meta"`
	Data   interface{}     `json:"data,omitempty"`
	Errors []ResponseError `json:"errors,omitempty"`
}
