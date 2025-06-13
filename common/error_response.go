package common

import "net/http"

type errorFormat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

type errorResponse struct {
	Success bool        `json:"success"`
	Error   errorFormat `json:"error"`
}

func NewErrorResponse(code int, message string, detail string) *errorResponse {
	err := errorFormat{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
	return &errorResponse{
		Success: false,
		Error:   err,
	}
}

func NewSimpleErrorResponse(message string) *errorResponse {
	err := errorFormat{
		Code:    http.StatusBadRequest,
		Message: "Error",
		Detail: message,
	}
	return &errorResponse{
		Success: false,
		Error:   err,
	}
}
