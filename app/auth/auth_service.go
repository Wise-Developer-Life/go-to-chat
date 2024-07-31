package auth

import (
	"go-to-chat/app/exception"
	"go-to-chat/app/user"
	"go-to-chat/app/utility"
	"time"
)

type LoginBody struct {
	Email    string
	Password string
}

// FIXME: change duration
var accessTokenDuration = time.Hour * 1
var refreshTokenDuration = time.Hour * 24 * 3

type AuthService interface {
	Login(body *LoginBody) (string, string, error)
	Logout() error
	RefreshAccessToken(refreshToken string) (string, error)
}

type AuthServiceImpl struct {
	UserService  user.UserService
	JwtService   JwtService
	PasswordUtil utility.PasswordUtil
}

func NewAuthService(userService user.UserService) AuthService {
	return &AuthServiceImpl{
		UserService:  userService,
		JwtService:   NewJwtService(),
		PasswordUtil: utility.NewPasswordUtil(),
	}
}

func (authService *AuthServiceImpl) Login(body *LoginBody) (string, string, error) {
	userByEmail, err := authService.UserService.GetUserByEmail(body.Email)

	if err != nil {
		return "", "", &exception.ResourceNotFoundError{ResourceName: "userByEmail", ResourceID: body.Email}
	}

	if !authService.PasswordUtil.CheckPasswordHash(body.Password, userByEmail.EncodedPassword) {
		return "", "", exception.NewAuthError()
	}

	accessToken, err := authService.JwtService.GenerateToken(userByEmail.Email, accessTokenDuration)
	refreshToken, err := authService.JwtService.GenerateToken(userByEmail.Email, refreshTokenDuration)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}

func (authService *AuthServiceImpl) Logout() error {
	return nil
}

func (authService *AuthServiceImpl) RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := authService.JwtService.ValidateToken(refreshToken)

	if err != nil {
		return "", exception.NewAuthError()
	}

	accessToken, err := authService.JwtService.GenerateToken(claims.Email, accessTokenDuration)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
