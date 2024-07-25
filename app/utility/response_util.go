package utility

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NotifyError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}

func SendSuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SendErrorResponse(c *gin.Context, statusCode int, errorCode string, message string) {
	res := ApiResponse{
		Status:  errorCode,
		Message: message,
	}
	c.JSON(statusCode, res)
}
