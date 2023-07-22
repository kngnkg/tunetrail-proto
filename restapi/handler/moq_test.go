// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"github.com/kngnkg/tunetrail/restapi/model"
	"sync"
)

// Ensure, that HealthServiceMock does implement HealthService.
// If this is not the case, regenerate this file with moq.
var _ HealthService = &HealthServiceMock{}

// HealthServiceMock is a mock implementation of HealthService.
//
//	func TestSomethingThatUsesHealthService(t *testing.T) {
//
//		// make and configure a mocked HealthService
//		mockedHealthService := &HealthServiceMock{
//			HealthCheckFunc: func(ctx context.Context) (*model.Health, error) {
//				panic("mock out the HealthCheck method")
//			},
//		}
//
//		// use mockedHealthService in code that requires HealthService
//		// and then make assertions.
//
//	}
type HealthServiceMock struct {
	// HealthCheckFunc mocks the HealthCheck method.
	HealthCheckFunc func(ctx context.Context) (*model.Health, error)

	// calls tracks calls to the methods.
	calls struct {
		// HealthCheck holds details about calls to the HealthCheck method.
		HealthCheck []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockHealthCheck sync.RWMutex
}

// HealthCheck calls HealthCheckFunc.
func (mock *HealthServiceMock) HealthCheck(ctx context.Context) (*model.Health, error) {
	if mock.HealthCheckFunc == nil {
		panic("HealthServiceMock.HealthCheckFunc: method is nil but HealthService.HealthCheck was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockHealthCheck.Lock()
	mock.calls.HealthCheck = append(mock.calls.HealthCheck, callInfo)
	mock.lockHealthCheck.Unlock()
	return mock.HealthCheckFunc(ctx)
}

// HealthCheckCalls gets all the calls that were made to HealthCheck.
// Check the length with:
//
//	len(mockedHealthService.HealthCheckCalls())
func (mock *HealthServiceMock) HealthCheckCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockHealthCheck.RLock()
	calls = mock.calls.HealthCheck
	mock.lockHealthCheck.RUnlock()
	return calls
}

// Ensure, that UserServiceMock does implement UserService.
// If this is not the case, regenerate this file with moq.
var _ UserService = &UserServiceMock{}

// UserServiceMock is a mock implementation of UserService.
//
//	func TestSomethingThatUsesUserService(t *testing.T) {
//
//		// make and configure a mocked UserService
//		mockedUserService := &UserServiceMock{
//			DeleteUserByUserNameFunc: func(ctx context.Context, userName string) error {
//				panic("mock out the DeleteUserByUserName method")
//			},
//			GetUserByUserNameFunc: func(ctx context.Context, userName string) (*model.User, error) {
//				panic("mock out the GetUserByUserName method")
//			},
//			RegisterUserFunc: func(ctx context.Context, details *model.UserRegistrationData) (*model.User, error) {
//				panic("mock out the RegisterUser method")
//			},
//			UpdateUserFunc: func(ctx context.Context, u *model.UserUpdateData) error {
//				panic("mock out the UpdateUser method")
//			},
//		}
//
//		// use mockedUserService in code that requires UserService
//		// and then make assertions.
//
//	}
type UserServiceMock struct {
	// DeleteUserByUserNameFunc mocks the DeleteUserByUserName method.
	DeleteUserByUserNameFunc func(ctx context.Context, userName string) error

	// GetUserByUserNameFunc mocks the GetUserByUserName method.
	GetUserByUserNameFunc func(ctx context.Context, userName string) (*model.User, error)

	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(ctx context.Context, details *model.UserRegistrationData) (*model.User, error)

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(ctx context.Context, u *model.UserUpdateData) error

	// calls tracks calls to the methods.
	calls struct {
		// DeleteUserByUserName holds details about calls to the DeleteUserByUserName method.
		DeleteUserByUserName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserName is the userName argument value.
			UserName string
		}
		// GetUserByUserName holds details about calls to the GetUserByUserName method.
		GetUserByUserName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserName is the userName argument value.
			UserName string
		}
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Details is the details argument value.
			Details *model.UserRegistrationData
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// U is the u argument value.
			U *model.UserUpdateData
		}
	}
	lockDeleteUserByUserName sync.RWMutex
	lockGetUserByUserName    sync.RWMutex
	lockRegisterUser         sync.RWMutex
	lockUpdateUser           sync.RWMutex
}

// DeleteUserByUserName calls DeleteUserByUserNameFunc.
func (mock *UserServiceMock) DeleteUserByUserName(ctx context.Context, userName string) error {
	if mock.DeleteUserByUserNameFunc == nil {
		panic("UserServiceMock.DeleteUserByUserNameFunc: method is nil but UserService.DeleteUserByUserName was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		UserName string
	}{
		Ctx:      ctx,
		UserName: userName,
	}
	mock.lockDeleteUserByUserName.Lock()
	mock.calls.DeleteUserByUserName = append(mock.calls.DeleteUserByUserName, callInfo)
	mock.lockDeleteUserByUserName.Unlock()
	return mock.DeleteUserByUserNameFunc(ctx, userName)
}

// DeleteUserByUserNameCalls gets all the calls that were made to DeleteUserByUserName.
// Check the length with:
//
//	len(mockedUserService.DeleteUserByUserNameCalls())
func (mock *UserServiceMock) DeleteUserByUserNameCalls() []struct {
	Ctx      context.Context
	UserName string
} {
	var calls []struct {
		Ctx      context.Context
		UserName string
	}
	mock.lockDeleteUserByUserName.RLock()
	calls = mock.calls.DeleteUserByUserName
	mock.lockDeleteUserByUserName.RUnlock()
	return calls
}

// GetUserByUserName calls GetUserByUserNameFunc.
func (mock *UserServiceMock) GetUserByUserName(ctx context.Context, userName string) (*model.User, error) {
	if mock.GetUserByUserNameFunc == nil {
		panic("UserServiceMock.GetUserByUserNameFunc: method is nil but UserService.GetUserByUserName was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		UserName string
	}{
		Ctx:      ctx,
		UserName: userName,
	}
	mock.lockGetUserByUserName.Lock()
	mock.calls.GetUserByUserName = append(mock.calls.GetUserByUserName, callInfo)
	mock.lockGetUserByUserName.Unlock()
	return mock.GetUserByUserNameFunc(ctx, userName)
}

// GetUserByUserNameCalls gets all the calls that were made to GetUserByUserName.
// Check the length with:
//
//	len(mockedUserService.GetUserByUserNameCalls())
func (mock *UserServiceMock) GetUserByUserNameCalls() []struct {
	Ctx      context.Context
	UserName string
} {
	var calls []struct {
		Ctx      context.Context
		UserName string
	}
	mock.lockGetUserByUserName.RLock()
	calls = mock.calls.GetUserByUserName
	mock.lockGetUserByUserName.RUnlock()
	return calls
}

// RegisterUser calls RegisterUserFunc.
func (mock *UserServiceMock) RegisterUser(ctx context.Context, details *model.UserRegistrationData) (*model.User, error) {
	if mock.RegisterUserFunc == nil {
		panic("UserServiceMock.RegisterUserFunc: method is nil but UserService.RegisterUser was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Details *model.UserRegistrationData
	}{
		Ctx:     ctx,
		Details: details,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(ctx, details)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedUserService.RegisterUserCalls())
func (mock *UserServiceMock) RegisterUserCalls() []struct {
	Ctx     context.Context
	Details *model.UserRegistrationData
} {
	var calls []struct {
		Ctx     context.Context
		Details *model.UserRegistrationData
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *UserServiceMock) UpdateUser(ctx context.Context, u *model.UserUpdateData) error {
	if mock.UpdateUserFunc == nil {
		panic("UserServiceMock.UpdateUserFunc: method is nil but UserService.UpdateUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		U   *model.UserUpdateData
	}{
		Ctx: ctx,
		U:   u,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(ctx, u)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//
//	len(mockedUserService.UpdateUserCalls())
func (mock *UserServiceMock) UpdateUserCalls() []struct {
	Ctx context.Context
	U   *model.UserUpdateData
} {
	var calls []struct {
		Ctx context.Context
		U   *model.UserUpdateData
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}
