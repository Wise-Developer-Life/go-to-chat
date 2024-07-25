package auth

import (
	"github.com/stretchr/testify/assert"
	"go-to-chat/app/exception"
	"os"
	"testing"
	"time"
)

var jwtService JwtService

func TestMain(m *testing.M) {
	jwtSecret = []byte("test_secret")
	jwtService = NewJwtService()
	code := m.Run()
	os.Exit(code)
}

func TestJwtService(t *testing.T) {
	jwtService = NewJwtService()
	assert.NotNil(t, jwtService)

	t.Run("Generate Token", func(t *testing.T) {
		token, err := jwtService.GenerateToken("test@email.com", time.Duration(20))
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Validate Token", func(t *testing.T) {
		t.Run("Valid Token", func(t *testing.T) {
			expiration := time.Second * 20
			token, err := jwtService.GenerateToken("test@email.com", expiration)
			claims, err := jwtService.ValidateToken(token)
			assert.NoError(t, err)
			assert.Equal(t, "test@email.com", claims.Email)
			assert.WithinDuration(
				t,
				claims.ExpiresAt.Time,
				time.Now().Add(expiration),
				time.Second,
			)
		})

		t.Run("Expired Token", func(t *testing.T) {
			expiration := time.Duration(0)
			token, err := jwtService.GenerateToken("test@email.com", expiration)
			claims, err := jwtService.ValidateToken(token)
			assert.ErrorIs(t, err, &exception.AuthError{})
			assert.Empty(t, claims)
		})

		t.Run("Invalid Token", func(t *testing.T) {
			token, err := jwtService.ValidateToken("invalid.token.input")
			assert.ErrorIs(t, err, &exception.AuthError{})
			assert.Empty(t, token)
		})
	})
}
