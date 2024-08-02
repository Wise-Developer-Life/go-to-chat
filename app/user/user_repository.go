package user

import (
	"errors"
	"go-to-chat/app/model"
	"go-to-chat/database"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
}

type userRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	result := database.GetDBInstance().Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepositoryImpl) GetUserById(id int) (*model.User, error) {
	var user model.User
	result := database.GetDBInstance().First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *userRepositoryImpl) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := database.GetDBInstance().Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (u *userRepositoryImpl) UpdateUser(user *model.User) (*model.User, error) {
	updatedResult := database.GetDBInstance().Updates(user)

	if updatedResult.Error != nil {
		return nil, updatedResult.Error
	}

	return user, nil
}
