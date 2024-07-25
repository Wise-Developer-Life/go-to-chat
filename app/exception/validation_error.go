package exception

import "net/http"

type ValidationError struct {
	message string
}

func (e *ValidationError) Error() string {
	return e.message
}

func (e *ValidationError) Code() string {
	return "validation_error"
}

func (e *ValidationError) HttpStatusCode() int {
	return http.StatusBadRequest
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		message: message,
	}
}
