package middleware

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/auth"
	"go-to-chat/app/exception"
	"go-to-chat/app/utility"
	"log"
)

const AuthorizationHeader = "Authorization"

type TokenType string

const (
	TokenTypeAccessToken  TokenType = "access"
	TokenTypeRefreshToken TokenType = "refresh"
)

func AuthHandler(tokenType TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(AuthorizationHeader)

		// FIXME: Implement 2 key pattern
		if tokenType == TokenTypeRefreshToken {
			log.Println("Token Type: Refresh Token")
		} else {
			log.Println("Token Type: Access Token")
		}

		jwtService := auth.NewJwtService()
		jwtClaim, err := jwtService.ValidateToken(token)

		if err != nil {
			utility.NotifyError(c, exception.NewAuthError())
			return
		}

		log.Println("Auth Success, Passing email to context: ", jwtClaim.Email)
		c.Set("user-info", jwtClaim.Email)
		c.Next()
	}
}
