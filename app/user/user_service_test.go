package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-to-chat/app/exception"
	"go-to-chat/app/model"
	"go-to-chat/mocks/go-to-chat/app/user"
	"go-to-chat/mocks/go-to-chat/app/utility"
	"testing"
)

func setUp(t *testing.T) (
	*assert.Assertions,
	*user.MockUserRepository,
	*utility.MockPasswordUtil,
	UserService,
) {
	mockedUserRepository := user.NewMockUserRepository(t)
	mockedPasswordUtility := utility.NewMockPasswordUtil(t)
	userService := &userServiceImpl{
		Repository:   mockedUserRepository,
		PasswordUtil: mockedPasswordUtility,
	}

	assert := assert.New(t)
	return assert, mockedUserRepository, mockedPasswordUtility, userService
}

func TestUserService(t *testing.T) {
	t.Run("Service Instance", func(t *testing.T) {
		assert := assert.New(t)
		userService := NewUserService(nil)
		assert.NotNil(userService)
	})

	t.Run("Get User Test", func(t *testing.T) {
		t.Run("User found", func(t *testing.T) {
			assert, mockedUserRepository, _, userService := setUp(t)

			mockedUser := &model.User{
				ID:              1,
				Name:            "Test User",
				Email:           "test@gmail.com",
				EncodedPassword: "password",
			}

			mockedUserRepository.EXPECT().GetUserById(1).Return(mockedUser, nil)
			user, err := userService.GetUser(1)

			assert.NoError(err)
			assert.Equal(mockedUser, user)
		})

		t.Run("User not found", func(t *testing.T) {
			assert, mockedUserRepository, _, userService := setUp(t)
			mockedUserRepository.EXPECT().GetUserById(1).Return(nil, errors.New("user not found"))

			user, err := userService.GetUser(1)

			var errNotFound *exception.ResourceNotFoundError
			assert.ErrorAs(err, &errNotFound)
			assert.Nil(user)
		})
	})
	t.Run("Create User", func(t *testing.T) {
		t.Run("Create User Success", func(t *testing.T) {
			assert, mockedUserRepository, mockedPasswordUtility, userService := setUp(t)

			createUserBody := &CreateUserBody{
				Name:     "Test User",
				Email:    "test@gmail.com",
				Password: "password",
			}

			mockedPasswordUtility.EXPECT().HashPassword("password").Return("password_hash", nil)
			userModelBeforeCreate := &model.User{
				Name:            "Test User",
				Email:           "test@gmail.com",
				EncodedPassword: "password_hash",
			}

			newUser := *userModelBeforeCreate
			newUser.ID = 1

			mockedUserRepository.EXPECT().GetUserByEmail("test@gmail.com").Return(nil, errors.New("user not found"))
			mockedUserRepository.EXPECT().CreateUser(userModelBeforeCreate).Return(&newUser, nil)

			user, err := userService.CreateUser(createUserBody)
			assert.NoError(err)
			assert.Equal(&newUser, user)

		})

		t.Run("Email is already used", func(t *testing.T) {
			assert, mockedUserRepository, _, userService := setUp(t)

			createUserBody := &CreateUserBody{
				Name:     "Test User",
				Email:    "test@gmail.com",
				Password: "password",
			}

			mockedUserRepository.EXPECT().GetUserByEmail("test@gmail.com").Return(&model.User{}, nil)

			user, err := userService.CreateUser(createUserBody)
			var errConflict *exception.ResourceConflictError
			assert.ErrorAs(err, &errConflict)
			assert.Nil(user)
		})
	})

	t.Run("Update User", func(t *testing.T) {
		t.Run("Update User Success", func(t *testing.T) {
			assert, mockedUserRepository, _, userService := setUp(t)

			updateUserBody := &UpdateUserBody{
				Name: "Test User Update",
			}

			existedUser := &model.User{
				ID:   1,
				Name: "Test User",
			}

			mockedUserRepository.EXPECT().GetUserById(1).Return(existedUser, nil)
			mockedUserRepository.EXPECT().UpdateUser(existedUser).Return(existedUser, nil)

			user, err := userService.UpdateUser(1, updateUserBody)
			assert.NoError(err)
			assert.Equal(&model.User{
				ID:   1,
				Name: "Test User Update",
			}, user)
		})

		t.Run("User not found", func(t *testing.T) {
			assert, mockedUserRepository, _, userService := setUp(t)

			updateUserBody := &UpdateUserBody{
				Name: "Test User",
			}

			mockedUserRepository.EXPECT().GetUserById(1).Return(nil, errors.New("user not found"))

			user, err := userService.UpdateUser(1, updateUserBody)
			var errNotFound *exception.ResourceNotFoundError
			assert.ErrorAs(err, &errNotFound)
			assert.Nil(user)
		})
	})

}
