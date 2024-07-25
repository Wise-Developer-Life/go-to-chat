package utility

import (
	"github.com/gin-gonic/gin"
)

// ValidateRequestJson TODO: use validator instead of ShouldBindJSON
func ValidateRequestJson[T any](c *gin.Context) (*T, error) {
	var jsonBody T

	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		return nil, err
	}

	return &jsonBody, nil
}
