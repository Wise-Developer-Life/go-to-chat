package utility

import "golang.org/x/crypto/bcrypt"

type PasswordUtil interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type passwordUtilImpl struct{}

func NewPasswordUtil() PasswordUtil {
	return &passwordUtilImpl{}
}

func (util *passwordUtilImpl) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (util *passwordUtilImpl) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
