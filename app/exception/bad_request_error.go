package exception

import "net/http"

type BadRequestError struct {
	message string
}

func (e *BadRequestError) Error() string {
	return e.message
}

func (e *BadRequestError) Code() string {
	return "bad_request"
}

func (e *BadRequestError) HttpStatusCode() int {
	return http.StatusBadRequest
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		message: message,
	}
}
