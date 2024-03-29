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
//			GetFolloweesFunc: func(ctx context.Context, userId model.UserID) ([]*model.User, error) {
//				panic("mock out the GetFollowees method")
//			},
//			GetFollowersFunc: func(ctx context.Context, userId model.UserID) ([]*model.User, error) {
//				panic("mock out the GetFollowers method")
//			},
//			GetSignedInUserFunc: func(ctx context.Context, userId model.UserID) (*model.User, error) {
//				panic("mock out the GetSignedInUser method")
//			},
//			GetUserByUserNameFunc: func(ctx context.Context, userName string, signedInUserId model.UserID) (*model.User, error) {
//				panic("mock out the GetUserByUserName method")
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

	// GetFolloweesFunc mocks the GetFollowees method.
	GetFolloweesFunc func(ctx context.Context, userId model.UserID) ([]*model.User, error)

	// GetFollowersFunc mocks the GetFollowers method.
	GetFollowersFunc func(ctx context.Context, userId model.UserID) ([]*model.User, error)

	// GetSignedInUserFunc mocks the GetSignedInUser method.
	GetSignedInUserFunc func(ctx context.Context, userId model.UserID) (*model.User, error)

	// GetUserByUserNameFunc mocks the GetUserByUserName method.
	GetUserByUserNameFunc func(ctx context.Context, userName string, signedInUserId model.UserID) (*model.User, error)

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
		// GetFollowees holds details about calls to the GetFollowees method.
		GetFollowees []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserId is the userId argument value.
			UserId model.UserID
		}
		// GetFollowers holds details about calls to the GetFollowers method.
		GetFollowers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserId is the userId argument value.
			UserId model.UserID
		}
		// GetSignedInUser holds details about calls to the GetSignedInUser method.
		GetSignedInUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserId is the userId argument value.
			UserId model.UserID
		}
		// GetUserByUserName holds details about calls to the GetUserByUserName method.
		GetUserByUserName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserName is the userName argument value.
			UserName string
			// SignedInUserId is the signedInUserId argument value.
			SignedInUserId model.UserID
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
	lockGetFollowees         sync.RWMutex
	lockGetFollowers         sync.RWMutex
	lockGetSignedInUser      sync.RWMutex
	lockGetUserByUserName    sync.RWMutex
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

// GetFollowees calls GetFolloweesFunc.
func (mock *UserServiceMock) GetFollowees(ctx context.Context, userId model.UserID) ([]*model.User, error) {
	if mock.GetFolloweesFunc == nil {
		panic("UserServiceMock.GetFolloweesFunc: method is nil but UserService.GetFollowees was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserId model.UserID
	}{
		Ctx:    ctx,
		UserId: userId,
	}
	mock.lockGetFollowees.Lock()
	mock.calls.GetFollowees = append(mock.calls.GetFollowees, callInfo)
	mock.lockGetFollowees.Unlock()
	return mock.GetFolloweesFunc(ctx, userId)
}

// GetFolloweesCalls gets all the calls that were made to GetFollowees.
// Check the length with:
//
//	len(mockedUserService.GetFolloweesCalls())
func (mock *UserServiceMock) GetFolloweesCalls() []struct {
	Ctx    context.Context
	UserId model.UserID
} {
	var calls []struct {
		Ctx    context.Context
		UserId model.UserID
	}
	mock.lockGetFollowees.RLock()
	calls = mock.calls.GetFollowees
	mock.lockGetFollowees.RUnlock()
	return calls
}

// GetFollowers calls GetFollowersFunc.
func (mock *UserServiceMock) GetFollowers(ctx context.Context, userId model.UserID) ([]*model.User, error) {
	if mock.GetFollowersFunc == nil {
		panic("UserServiceMock.GetFollowersFunc: method is nil but UserService.GetFollowers was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserId model.UserID
	}{
		Ctx:    ctx,
		UserId: userId,
	}
	mock.lockGetFollowers.Lock()
	mock.calls.GetFollowers = append(mock.calls.GetFollowers, callInfo)
	mock.lockGetFollowers.Unlock()
	return mock.GetFollowersFunc(ctx, userId)
}

// GetFollowersCalls gets all the calls that were made to GetFollowers.
// Check the length with:
//
//	len(mockedUserService.GetFollowersCalls())
func (mock *UserServiceMock) GetFollowersCalls() []struct {
	Ctx    context.Context
	UserId model.UserID
} {
	var calls []struct {
		Ctx    context.Context
		UserId model.UserID
	}
	mock.lockGetFollowers.RLock()
	calls = mock.calls.GetFollowers
	mock.lockGetFollowers.RUnlock()
	return calls
}

// GetSignedInUser calls GetSignedInUserFunc.
func (mock *UserServiceMock) GetSignedInUser(ctx context.Context, userId model.UserID) (*model.User, error) {
	if mock.GetSignedInUserFunc == nil {
		panic("UserServiceMock.GetSignedInUserFunc: method is nil but UserService.GetSignedInUser was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserId model.UserID
	}{
		Ctx:    ctx,
		UserId: userId,
	}
	mock.lockGetSignedInUser.Lock()
	mock.calls.GetSignedInUser = append(mock.calls.GetSignedInUser, callInfo)
	mock.lockGetSignedInUser.Unlock()
	return mock.GetSignedInUserFunc(ctx, userId)
}

// GetSignedInUserCalls gets all the calls that were made to GetSignedInUser.
// Check the length with:
//
//	len(mockedUserService.GetSignedInUserCalls())
func (mock *UserServiceMock) GetSignedInUserCalls() []struct {
	Ctx    context.Context
	UserId model.UserID
} {
	var calls []struct {
		Ctx    context.Context
		UserId model.UserID
	}
	mock.lockGetSignedInUser.RLock()
	calls = mock.calls.GetSignedInUser
	mock.lockGetSignedInUser.RUnlock()
	return calls
}

// GetUserByUserName calls GetUserByUserNameFunc.
func (mock *UserServiceMock) GetUserByUserName(ctx context.Context, userName string, signedInUserId model.UserID) (*model.User, error) {
	if mock.GetUserByUserNameFunc == nil {
		panic("UserServiceMock.GetUserByUserNameFunc: method is nil but UserService.GetUserByUserName was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		UserName       string
		SignedInUserId model.UserID
	}{
		Ctx:            ctx,
		UserName:       userName,
		SignedInUserId: signedInUserId,
	}
	mock.lockGetUserByUserName.Lock()
	mock.calls.GetUserByUserName = append(mock.calls.GetUserByUserName, callInfo)
	mock.lockGetUserByUserName.Unlock()
	return mock.GetUserByUserNameFunc(ctx, userName, signedInUserId)
}

// GetUserByUserNameCalls gets all the calls that were made to GetUserByUserName.
// Check the length with:
//
//	len(mockedUserService.GetUserByUserNameCalls())
func (mock *UserServiceMock) GetUserByUserNameCalls() []struct {
	Ctx            context.Context
	UserName       string
	SignedInUserId model.UserID
} {
	var calls []struct {
		Ctx            context.Context
		UserName       string
		SignedInUserId model.UserID
	}
	mock.lockGetUserByUserName.RLock()
	calls = mock.calls.GetUserByUserName
	mock.lockGetUserByUserName.RUnlock()
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

// Ensure, that AuthServiceMock does implement AuthService.
// If this is not the case, regenerate this file with moq.
var _ AuthService = &AuthServiceMock{}

// AuthServiceMock is a mock implementation of AuthService.
//
//	func TestSomethingThatUsesAuthService(t *testing.T) {
//
//		// make and configure a mocked AuthService
//		mockedAuthService := &AuthServiceMock{
//			ConfirmEmailFunc: func(ctx context.Context, userName string, code string) error {
//				panic("mock out the ConfirmEmail method")
//			},
//			GetSignedInUserFunc: func(ctx context.Context, idToken string) (*model.User, error) {
//				panic("mock out the GetSignedInUser method")
//			},
//			RefreshTokenFunc: func(ctx context.Context, idToken string, refreshToken string) (*model.Tokens, error) {
//				panic("mock out the RefreshToken method")
//			},
//			RegisterUserFunc: func(ctx context.Context, data *model.UserRegistrationData) (*model.User, error) {
//				panic("mock out the RegisterUser method")
//			},
//			SignInFunc: func(ctx context.Context, data *model.UserSignInData) (*model.Tokens, error) {
//				panic("mock out the SignIn method")
//			},
//		}
//
//		// use mockedAuthService in code that requires AuthService
//		// and then make assertions.
//
//	}
type AuthServiceMock struct {
	// ConfirmEmailFunc mocks the ConfirmEmail method.
	ConfirmEmailFunc func(ctx context.Context, userName string, code string) error

	// GetSignedInUserFunc mocks the GetSignedInUser method.
	GetSignedInUserFunc func(ctx context.Context, idToken string) (*model.User, error)

	// RefreshTokenFunc mocks the RefreshToken method.
	RefreshTokenFunc func(ctx context.Context, idToken string, refreshToken string) (*model.Tokens, error)

	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(ctx context.Context, data *model.UserRegistrationData) (*model.User, error)

	// SignInFunc mocks the SignIn method.
	SignInFunc func(ctx context.Context, data *model.UserSignInData) (*model.Tokens, error)

	// calls tracks calls to the methods.
	calls struct {
		// ConfirmEmail holds details about calls to the ConfirmEmail method.
		ConfirmEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserName is the userName argument value.
			UserName string
			// Code is the code argument value.
			Code string
		}
		// GetSignedInUser holds details about calls to the GetSignedInUser method.
		GetSignedInUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// IdToken is the idToken argument value.
			IdToken string
		}
		// RefreshToken holds details about calls to the RefreshToken method.
		RefreshToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// IdToken is the idToken argument value.
			IdToken string
			// RefreshToken is the refreshToken argument value.
			RefreshToken string
		}
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data *model.UserRegistrationData
		}
		// SignIn holds details about calls to the SignIn method.
		SignIn []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Data is the data argument value.
			Data *model.UserSignInData
		}
	}
	lockConfirmEmail    sync.RWMutex
	lockGetSignedInUser sync.RWMutex
	lockRefreshToken    sync.RWMutex
	lockRegisterUser    sync.RWMutex
	lockSignIn          sync.RWMutex
}

// ConfirmEmail calls ConfirmEmailFunc.
func (mock *AuthServiceMock) ConfirmEmail(ctx context.Context, userName string, code string) error {
	if mock.ConfirmEmailFunc == nil {
		panic("AuthServiceMock.ConfirmEmailFunc: method is nil but AuthService.ConfirmEmail was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		UserName string
		Code     string
	}{
		Ctx:      ctx,
		UserName: userName,
		Code:     code,
	}
	mock.lockConfirmEmail.Lock()
	mock.calls.ConfirmEmail = append(mock.calls.ConfirmEmail, callInfo)
	mock.lockConfirmEmail.Unlock()
	return mock.ConfirmEmailFunc(ctx, userName, code)
}

// ConfirmEmailCalls gets all the calls that were made to ConfirmEmail.
// Check the length with:
//
//	len(mockedAuthService.ConfirmEmailCalls())
func (mock *AuthServiceMock) ConfirmEmailCalls() []struct {
	Ctx      context.Context
	UserName string
	Code     string
} {
	var calls []struct {
		Ctx      context.Context
		UserName string
		Code     string
	}
	mock.lockConfirmEmail.RLock()
	calls = mock.calls.ConfirmEmail
	mock.lockConfirmEmail.RUnlock()
	return calls
}

// GetSignedInUser calls GetSignedInUserFunc.
func (mock *AuthServiceMock) GetSignedInUser(ctx context.Context, idToken string) (*model.User, error) {
	if mock.GetSignedInUserFunc == nil {
		panic("AuthServiceMock.GetSignedInUserFunc: method is nil but AuthService.GetSignedInUser was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		IdToken string
	}{
		Ctx:     ctx,
		IdToken: idToken,
	}
	mock.lockGetSignedInUser.Lock()
	mock.calls.GetSignedInUser = append(mock.calls.GetSignedInUser, callInfo)
	mock.lockGetSignedInUser.Unlock()
	return mock.GetSignedInUserFunc(ctx, idToken)
}

// GetSignedInUserCalls gets all the calls that were made to GetSignedInUser.
// Check the length with:
//
//	len(mockedAuthService.GetSignedInUserCalls())
func (mock *AuthServiceMock) GetSignedInUserCalls() []struct {
	Ctx     context.Context
	IdToken string
} {
	var calls []struct {
		Ctx     context.Context
		IdToken string
	}
	mock.lockGetSignedInUser.RLock()
	calls = mock.calls.GetSignedInUser
	mock.lockGetSignedInUser.RUnlock()
	return calls
}

// RefreshToken calls RefreshTokenFunc.
func (mock *AuthServiceMock) RefreshToken(ctx context.Context, idToken string, refreshToken string) (*model.Tokens, error) {
	if mock.RefreshTokenFunc == nil {
		panic("AuthServiceMock.RefreshTokenFunc: method is nil but AuthService.RefreshToken was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		IdToken      string
		RefreshToken string
	}{
		Ctx:          ctx,
		IdToken:      idToken,
		RefreshToken: refreshToken,
	}
	mock.lockRefreshToken.Lock()
	mock.calls.RefreshToken = append(mock.calls.RefreshToken, callInfo)
	mock.lockRefreshToken.Unlock()
	return mock.RefreshTokenFunc(ctx, idToken, refreshToken)
}

// RefreshTokenCalls gets all the calls that were made to RefreshToken.
// Check the length with:
//
//	len(mockedAuthService.RefreshTokenCalls())
func (mock *AuthServiceMock) RefreshTokenCalls() []struct {
	Ctx          context.Context
	IdToken      string
	RefreshToken string
} {
	var calls []struct {
		Ctx          context.Context
		IdToken      string
		RefreshToken string
	}
	mock.lockRefreshToken.RLock()
	calls = mock.calls.RefreshToken
	mock.lockRefreshToken.RUnlock()
	return calls
}

// RegisterUser calls RegisterUserFunc.
func (mock *AuthServiceMock) RegisterUser(ctx context.Context, data *model.UserRegistrationData) (*model.User, error) {
	if mock.RegisterUserFunc == nil {
		panic("AuthServiceMock.RegisterUserFunc: method is nil but AuthService.RegisterUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data *model.UserRegistrationData
	}{
		Ctx:  ctx,
		Data: data,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(ctx, data)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedAuthService.RegisterUserCalls())
func (mock *AuthServiceMock) RegisterUserCalls() []struct {
	Ctx  context.Context
	Data *model.UserRegistrationData
} {
	var calls []struct {
		Ctx  context.Context
		Data *model.UserRegistrationData
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}

// SignIn calls SignInFunc.
func (mock *AuthServiceMock) SignIn(ctx context.Context, data *model.UserSignInData) (*model.Tokens, error) {
	if mock.SignInFunc == nil {
		panic("AuthServiceMock.SignInFunc: method is nil but AuthService.SignIn was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Data *model.UserSignInData
	}{
		Ctx:  ctx,
		Data: data,
	}
	mock.lockSignIn.Lock()
	mock.calls.SignIn = append(mock.calls.SignIn, callInfo)
	mock.lockSignIn.Unlock()
	return mock.SignInFunc(ctx, data)
}

// SignInCalls gets all the calls that were made to SignIn.
// Check the length with:
//
//	len(mockedAuthService.SignInCalls())
func (mock *AuthServiceMock) SignInCalls() []struct {
	Ctx  context.Context
	Data *model.UserSignInData
} {
	var calls []struct {
		Ctx  context.Context
		Data *model.UserSignInData
	}
	mock.lockSignIn.RLock()
	calls = mock.calls.SignIn
	mock.lockSignIn.RUnlock()
	return calls
}
