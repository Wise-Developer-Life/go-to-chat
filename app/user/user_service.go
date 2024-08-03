package user

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go-to-chat/app/exception"
	"go-to-chat/app/model"
	"go-to-chat/app/utility"
	"mime/multipart"
	"path"
	"path/filepath"
	"strconv"
)

// FIXME: refactor this in future
const (
	awsRegion    = "us-east-1"
	awsServerUrl = "http://localhost:4566"
	s3AppBucket  = "go-to-chat"
	awsAccessKey = "test"
	awsSecretKey = "test"
)

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
	return path.Join("profile_images", strconv.Itoa(userId)+fileExt)
}

func (u *userServiceImpl) UploadProfileImage(userId int, file *multipart.FileHeader) error {
	existedUser, err := u.Repository.GetUserById(userId)

	if err != nil {
		return exception.NewResourceNotFoundError("user", strconv.Itoa(userId))
	}

	// create a new aws session
	awsClient, err := session.NewSession(&aws.Config{
		Region:           aws.String(awsRegion),
		Endpoint:         aws.String(awsServerUrl),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	})

	if err != nil {
		return err
	}

	// open file
	fileContent, err := file.Open()
	if err != nil {
		return err
	}
	defer func(fileContent multipart.File) {
		err := fileContent.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fileContent)

	s3Svc := s3.New(awsClient)
	_, err = s3Svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(s3AppBucket),
	})

	err = s3Svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(s3AppBucket),
	})
	if err != nil {
		return fmt.Errorf("failed to wait for bucket to exist: %v", err)
	}

	uploader := s3manager.NewUploader(awsClient)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3AppBucket),
		Key:    aws.String(generateProfileImagePath(userId, file.Filename)),
		Body:   fileContent,
	})

	if err != nil {
		return err
	}

	existedUser.ProfileUrl = result.Location
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
