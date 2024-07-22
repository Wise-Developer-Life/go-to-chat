package userService

import (
	"go-to-chat/app/exception"
	"go-to-chat/app/model"
	userRepository "go-to-chat/app/repository"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type CreateUserBody struct {
	Name     string
	Email    string
	Password string
}

type UpdateUserBody struct {
	Name     string
	Password string
}

func CreateUser(body *CreateUserBody) (*model.User, error) {
	hashedPassword, err := hashPassword(body.Password)

	if err != nil {
		return nil, err
	}

	userModel := model.User{
		Name:            body.Name,
		Email:           body.Email,
		EncodedPassword: hashedPassword,
	}
	newUser, err := userRepository.CreateUser(&userModel)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func GetUser(userId int) (*model.User, error) {
	userModel, err := userRepository.GetUserById(userId)

	if err != nil {
		return nil, &exception.ResourceNotFoundError{
			ResourceName: "user",
			ResourceID:   strconv.Itoa(userId),
		}
	}

	return userModel, nil
}

func UpdateUser(userId int, body *UpdateUserBody) (*model.User, error) {
	existedUser, err := userRepository.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(body.Password)
	if err != nil {
		return nil, err
	}

	existedUser.Name = body.Name
	existedUser.EncodedPassword = hashedPassword

	updatedUser, err := userRepository.UpdateUser(existedUser)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
