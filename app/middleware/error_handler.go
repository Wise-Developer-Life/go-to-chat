package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-to-chat/app/exception"
	"go-to-chat/app/utility"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		err := context.Errors.Last()
		if err == nil {
			return
		}

		var customErr exception.BaseError
		if ok := errors.As(err, &customErr); ok {
			utility.SendErrorResponse(
				context,
				customErr.HttpStatusCode(),
				customErr.Code(),
				customErr.Error(),
			)
		} else {
			utility.SendErrorResponse(
				context,
				http.StatusInternalServerError,
				"internal_server_error",
				err.Error(),
			)
		}
	}
}
