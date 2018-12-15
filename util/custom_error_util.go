package util

import "net/http"

type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// BadRequest : Returns a 400 error encapsulated as *HTTPError
func BadRequest(message string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

// InternalServerError : Returns a 500 error encapsulated as *HTTPError
func InternalServerError(message string) *HTTPError {
	return &HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
