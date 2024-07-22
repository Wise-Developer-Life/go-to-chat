package userRepository

import (
	"errors"
	"go-to-chat/app/model"
	"go-to-chat/database"
	"gorm.io/gorm"
)

func CreateUser(user *model.User) (*model.User, error) {
	result := database.GetDBInstance().Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserById(id int) (*model.User, error) {
	var user model.User
	result := database.GetDBInstance().First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func UpdateUser(user *model.User) (*model.User, error) {
	// Perform the update
	updateResult := database.GetDBInstance().Updates(user)
	if updateResult.Error != nil {
		return nil, updateResult.Error
	}

	return user, nil
}

//// UpdateUser Update a user
//func (repository *UserRepositoryImpl) UpdateUser(user *model.User) (*model.User, exception) {
//	return nil
//}
