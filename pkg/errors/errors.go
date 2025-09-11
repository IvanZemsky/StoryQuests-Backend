package commonErrors

import (
	"fmt"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", err.Code, err.Message)
}

func NewError(code, message string) *ErrorResponse {
	return &ErrorResponse{Code: code, Message: message}
}

type BaseErrorResponse struct {
	Error string `json:"error"`
}
