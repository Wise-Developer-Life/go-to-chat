package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go-to-chat/app/exception"
	"time"
)

// FIXME: change this to a more secure secret
var jwtSecret = []byte("secret")

type JwtClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(email string, expiration time.Duration) (string, error)
	ValidateToken(tokenString string) (*JwtClaims, error)
}

type jwtServiceImpl struct{}

func NewJwtService() JwtService {
	return &jwtServiceImpl{}
}

func (jwtService *jwtServiceImpl) GenerateToken(email string, expiration time.Duration) (string, error) {
	claims := &JwtClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Subject:   email,
			Issuer:    "go-to-chat",
			Audience:  jwt.ClaimStrings{"go-to-chat"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtService *jwtServiceImpl) ValidateToken(tokenString string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
		jwt.WithIssuer("go-to-chat"),
		jwt.WithAudience("go-to-chat"),
		jwt.WithExpirationRequired(),
	)

	if err != nil || !token.Valid {
		return nil, exception.NewAuthError()
	}

	return claims, nil
}
