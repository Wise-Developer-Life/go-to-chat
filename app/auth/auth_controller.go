package auth

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/exception"
	"go-to-chat/app/user"
	"go-to-chat/app/utility"
	"net/http"
)

// FIXME: use service provider pattern
var authService = NewAuthService(user.NewUserService(user.NewUserRepository()))

func Login(c *gin.Context) {
	requestBody, err := utility.ValidateRequestJson[LoginRequest](c)

	if err != nil {
		utility.NotifyError(c, exception.NewValidationError(err.Error()))
		return
	}

	accessToken, refreshToken, err := authService.Login(
		&LoginBody{
			Email:    requestBody.Email,
			Password: requestBody.Password,
		})

	if err != nil {
		utility.NotifyError(c, exception.NewAuthError())
		return
	}

	utility.SendSuccessResponse(c,
		http.StatusOK,
		"success",
		NewLoginResponse(accessToken, refreshToken),
	)
}

func Logout(c *gin.Context) {
	err := authService.Logout()

	if err != nil {
		utility.NotifyError(c, err)
		return
	}

	utility.SendSuccessResponse(c, http.StatusOK, "success", nil)
}

func RefreshToken(c *gin.Context) {
	token := c.GetHeader("Authorization")

	newToken, err := authService.RefreshAccessToken(token)

	if err != nil {
		utility.NotifyError(c, err)
		return
	}

	utility.SendSuccessResponse(c,
		http.StatusOK,
		"success",
		NewRefreshTokenResponse(newToken),
	)
}
