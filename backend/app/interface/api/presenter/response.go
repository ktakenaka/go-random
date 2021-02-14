package presenter

// ResponseMeta is Meta Object of JSON API
// Every response includes Meta even if it's the error response
type ResponseMeta struct {
	Code    MetaCode `json:"code"`
	Message string   `json:"message"`
}

// ResponseError provides additional information about problems encountered while performing an operation.
// It will be included in the array, key will be `errors`
type ResponseError struct {
	Source ResponseErrorSource `json:"source"`
	Detail string              `json:"detail"`
}

// JSONAPIError follows json api
type JSONAPIError struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Source string `json:"source"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// ResponseErrorSource is an object containing references to the source of the error,
// optionally including any of the following members:
type ResponseErrorSource struct {
	Pointer string `json:"pointer,omitempty"`
	Param   string `json:"param,omitempty"`
}

// Response is the Response format of JSON API
type Response struct {
	Meta   ResponseMeta   `json:"meta"`
	Data   interface{}    `json:"data,omitempty"`
	Errors []JSONAPIError `json:"errors,omitempty"`
}

// MetaCode is the type of MetaCode
// First 3 letters must be same as http status
type MetaCode uint

// MetaCode for success cases
const (
	CodeSuccess   MetaCode = 20001
	CodeCreated   MetaCode = 20101
	CodeAccepted  MetaCode = 20201
	CodeNoContent MetaCode = 20401
)

// MetaCode for client errors
const (
	CodeBadRequest          MetaCode = 40001
	CodeUnauthorized        MetaCode = 40101
	CodeForbidden           MetaCode = 40301
	CodeNotFound            MetaCode = 40401
	CodeUnprocessableEntity MetaCode = 42201
)

// MetaCode for server errors
const (
	CodeInternalServerError MetaCode = 50001
	CodePanic               MetaCode = 50002
)

// MetaCodePair expects to be used in handler and formatter
// SuccessCase => put MetaCode in handler
// ErrorCase => in formatter
var MetaCodePair = map[MetaCode]ResponseMeta{
	CodeSuccess: {
		Code:    CodeSuccess,
		Message: "success",
	},
	CodeCreated: {
		Code:    CodeCreated,
		Message: "created",
	},
	CodeAccepted: {
		Code:    CodeAccepted,
		Message: "accepted",
	},
	CodeNoContent: {
		Code:    CodeNoContent,
		Message: "no content",
	},
	CodeBadRequest: {
		Code:    CodeBadRequest,
		Message: "bad request",
	},
	CodeUnauthorized: {
		Code:    CodeUnauthorized,
		Message: "unauthorized",
	},
	CodeForbidden: {
		Code:    CodeForbidden,
		Message: "forbidden",
	},
	CodeNotFound: {
		Code:    CodeNotFound,
		Message: "not found",
	},
	CodeUnprocessableEntity: {
		Code:    CodeUnprocessableEntity,
		Message: "unprocessable entity",
	},
	CodeInternalServerError: {
		Code:    CodeInternalServerError,
		Message: "internal server error",
	},
	CodePanic: {
		Code:    CodePanic,
		Message: "panic occurred",
	},
}
