// Code generated by mockery. DO NOT EDIT.

package user

import (
	model "go-to-chat/app/model"

	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock payload for the UserRepository payload
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: _a0
func (_m *MockUserRepository) CreateUser(_a0 *model.User) (*model.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User) (*model.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with payload explicit version for method 'CreateUser'
type MockUserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - _a0 *model.User
func (_e *MockUserRepository_Expecter) CreateUser(_a0 interface{}) *MockUserRepository_CreateUser_Call {
	return &MockUserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", _a0)}
}

func (_c *MockUserRepository_CreateUser_Call) Run(run func(_a0 *model.User)) *MockUserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) Return(_a0 *model.User, _a1 error) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_CreateUser_Call) RunAndReturn(run func(*model.User) (*model.User, error)) *MockUserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *MockUserRepository) GetUserByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with payload explicit version for method 'GetUserByEmail'
type MockUserRepository_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockUserRepository_Expecter) GetUserByEmail(email interface{}) *MockUserRepository_GetUserByEmail_Call {
	return &MockUserRepository_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", email)}
}

func (_c *MockUserRepository_GetUserByEmail_Call) Run(run func(email string)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) Return(_a0 *model.User, _a1 error) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUserByEmail_Call) RunAndReturn(run func(string) (*model.User, error)) *MockUserRepository_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: id
func (_m *MockUserRepository) GetUserById(id int) (*model.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*model.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetUserById_Call is a *mock.Call that shadows Run/Return methods with payload explicit version for method 'GetUserById'
type MockUserRepository_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - id int
func (_e *MockUserRepository_Expecter) GetUserById(id interface{}) *MockUserRepository_GetUserById_Call {
	return &MockUserRepository_GetUserById_Call{Call: _e.mock.On("GetUserById", id)}
}

func (_c *MockUserRepository_GetUserById_Call) Run(run func(id int)) *MockUserRepository_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockUserRepository_GetUserById_Call) Return(_a0 *model.User, _a1 error) *MockUserRepository_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetUserById_Call) RunAndReturn(run func(int) (*model.User, error)) *MockUserRepository_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *MockUserRepository) UpdateUser(_a0 *model.User) (*model.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User) (*model.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with payload explicit version for method 'UpdateUser'
type MockUserRepository_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - _a0 *model.User
func (_e *MockUserRepository_Expecter) UpdateUser(_a0 interface{}) *MockUserRepository_UpdateUser_Call {
	return &MockUserRepository_UpdateUser_Call{Call: _e.mock.On("UpdateUser", _a0)}
}

func (_c *MockUserRepository_UpdateUser_Call) Run(run func(_a0 *model.User)) *MockUserRepository_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockUserRepository_UpdateUser_Call) Return(_a0 *model.User, _a1 error) *MockUserRepository_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_UpdateUser_Call) RunAndReturn(run func(*model.User) (*model.User, error)) *MockUserRepository_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
