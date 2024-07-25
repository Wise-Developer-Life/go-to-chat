package exception

import "net/http"

type AuthError struct{}

func NewAuthError() *AuthError {
	return &AuthError{}
}

func (e *AuthError) Error() string {
	return "Authentication failed"
}

func (e *AuthError) Code() string {
	return "auth_fail"
}

func (e *AuthError) HttpStatusCode() int {
	return http.StatusUnauthorized
}
