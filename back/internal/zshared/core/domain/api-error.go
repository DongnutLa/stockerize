package shared_domain

import (
	"fmt"
)

type ErrorGroup string

var (
	GeneralErrors    ErrorGroup = "GEN"
	AuthErrors       ErrorGroup = "AUT"
	UserErrors       ErrorGroup = "USR"
	NewsletterErrors ErrorGroup = "NWS"
	AdminErrors      ErrorGroup = "ADM"
	FileErrors       ErrorGroup = "FIL"
	TopicErrors      ErrorGroup = "TOP"
)

type ApiError struct {
	Message        string      `json:"message"`
	Code           string      `json:"code"`
	HttpStatusCode int         `json:"statusCode,omitempty"`
	Detail         interface{} `json:"detail,omitempty"`
}

func NewApiError(
	message string,
	group ErrorGroup,
	httpCode int,
	consecutive int,
) *ApiError {
	return &ApiError{
		Message:        message,
		Code:           buildErrorCode(group, consecutive),
		HttpStatusCode: httpCode,
	}
}

func buildErrorCode(group ErrorGroup, consecutive int) string {
	return fmt.Sprintf(`%s-%s%0*d`, "ERR", group, 3, consecutive)
}

func (e *ApiError) SetDetail(d interface{}) *ApiError {
	return &ApiError{
		Message:        e.Message,
		Code:           e.Code,
		HttpStatusCode: e.HttpStatusCode,
		Detail:         d,
	}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf(`%s: %s`, e.Code, e.Message)
}
