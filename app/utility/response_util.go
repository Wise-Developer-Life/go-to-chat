package utility

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/dto/response"
)

func SendSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	res := response.ApiResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	c.JSON(statusCode, res)
}

func SendErrorResponse(c *gin.Context, statusCode int, errorCode string, message string) {
	res := response.ApiResponse{
		Status:  errorCode,
		Message: message,
	}
	c.JSON(statusCode, res)
}
