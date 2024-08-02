package user

import (
	"fmt"
	"go-to-chat/app/exception"
	"go-to-chat/app/model"
	"go-to-chat/app/utility"
	"mime/multipart"
	"path"
	"path/filepath"
	"strconv"
)

const imageFileRootDir = "./data/images"

type CreateUserBody struct {
	Name     string
	Email    string
	Password string
}

type UpdateUserBody struct {
	Name string
}

type UserService interface {
	CreateUser(body *CreateUserBody) (*model.User, error)
	GetUser(userId int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(userId int, body *UpdateUserBody) (*model.User, error)
	UploadProfileImage(userId int, file *multipart.FileHeader) error
	GetProfileImage(userId int, fileName string) (string, error)
}

type userServiceImpl struct {
	Repository   UserRepository
	PasswordUtil utility.PasswordUtil
}

func NewUserService(repository UserRepository) UserService {
	return &userServiceImpl{
		Repository:   repository,
		PasswordUtil: utility.NewPasswordUtil(),
	}
}

func (u *userServiceImpl) CreateUser(body *CreateUserBody) (*model.User, error) {
	userWithEmail, err := u.Repository.GetUserByEmail(body.Email)
	if userWithEmail != nil {
		return nil,
			exception.NewResourceConflictError(
				"user",
				fmt.Sprintf("User with email %s already exists", body.Email),
			)
	}

	hashedPassword, err := u.PasswordUtil.HashPassword(body.Password)
	if err != nil {
		return nil, err
	}

	userModel := model.User{
		Name:            body.Name,
		Email:           body.Email,
		EncodedPassword: hashedPassword,
	}
	newUser, err := u.Repository.CreateUser(&userModel)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *userServiceImpl) GetUser(userId int) (*model.User, error) {
	userModel, err := u.Repository.GetUserById(userId)

	if err != nil {
		return nil, &exception.ResourceNotFoundError{
			ResourceName: "user",
			ResourceID:   strconv.Itoa(userId),
		}
	}

	return userModel, nil
}

func (u *userServiceImpl) GetUserByEmail(email string) (*model.User, error) {
	userModel, err := u.Repository.GetUserByEmail(email)

	if err != nil {
		return nil, &exception.ResourceNotFoundError{
			ResourceName: "user",
			ResourceID:   email,
		}
	}

	return userModel, nil
}

func (u *userServiceImpl) UpdateUser(userId int, body *UpdateUserBody) (*model.User, error) {
	existedUser, err := u.Repository.GetUserById(userId)

	if err != nil {
		return nil, exception.NewResourceNotFoundError("user", strconv.Itoa(userId))
	}

	existedUser.Name = body.Name

	updatedUser, err := u.Repository.UpdateUser(existedUser)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func generateProfileImagePath(userId int, fileName string) string {
	fileExt := filepath.Ext(fileName)
	return path.Join(imageFileRootDir, strconv.Itoa(userId), "profile_img"+fileExt)
}

func (u *userServiceImpl) UploadProfileImage(userId int, file *multipart.FileHeader) error {
	existedUser, err := u.Repository.GetUserById(userId)

	if err != nil {
		return exception.NewResourceNotFoundError("user", strconv.Itoa(userId))
	}

	filePath := generateProfileImagePath(userId, file.Filename)
	_, err = utility.SaveFileLocally(file, filePath)

	if err != nil {
		return err
	}

	existedUser.ProfileUrl = fmt.Sprintf("http://localhost:8082/api/v1/user/%d/profile-image?file=%s", userId, filepath.Base(filePath))
	_, err = u.Repository.UpdateUser(existedUser)

	if err != nil {
		return err
	}

	return nil
}

func (u *userServiceImpl) GetProfileImage(userId int, fileName string) (string, error) {
	existedUser, err := u.Repository.GetUserById(userId)

	if err != nil {
		return "", exception.NewResourceNotFoundError("user", strconv.Itoa(userId))
	}

	if existedUser.ProfileUrl == "" {
		return "", exception.NewResourceNotFoundError("profile image", strconv.Itoa(userId))
	}

	return generateProfileImagePath(userId, fileName), nil
}
