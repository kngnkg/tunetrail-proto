// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
	"sync"
)

// Ensure, that HealthRepositoryMock does implement HealthRepository.
// If this is not the case, regenerate this file with moq.
var _ HealthRepository = &HealthRepositoryMock{}

// HealthRepositoryMock is a mock implementation of HealthRepository.
//
//	func TestSomethingThatUsesHealthRepository(t *testing.T) {
//
//		// make and configure a mocked HealthRepository
//		mockedHealthRepository := &HealthRepositoryMock{
//			PingFunc: func(ctx context.Context, db store.Queryer) error {
//				panic("mock out the Ping method")
//			},
//		}
//
//		// use mockedHealthRepository in code that requires HealthRepository
//		// and then make assertions.
//
//	}
type HealthRepositoryMock struct {
	// PingFunc mocks the Ping method.
	PingFunc func(ctx context.Context, db store.Queryer) error

	// calls tracks calls to the methods.
	calls struct {
		// Ping holds details about calls to the Ping method.
		Ping []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
		}
	}
	lockPing sync.RWMutex
}

// Ping calls PingFunc.
func (mock *HealthRepositoryMock) Ping(ctx context.Context, db store.Queryer) error {
	if mock.PingFunc == nil {
		panic("HealthRepositoryMock.PingFunc: method is nil but HealthRepository.Ping was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Queryer
	}{
		Ctx: ctx,
		Db:  db,
	}
	mock.lockPing.Lock()
	mock.calls.Ping = append(mock.calls.Ping, callInfo)
	mock.lockPing.Unlock()
	return mock.PingFunc(ctx, db)
}

// PingCalls gets all the calls that were made to Ping.
// Check the length with:
//
//	len(mockedHealthRepository.PingCalls())
func (mock *HealthRepositoryMock) PingCalls() []struct {
	Ctx context.Context
	Db  store.Queryer
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Queryer
	}
	mock.lockPing.RLock()
	calls = mock.calls.Ping
	mock.lockPing.RUnlock()
	return calls
}

// Ensure, that UserRepositoryMock does implement UserRepository.
// If this is not the case, regenerate this file with moq.
var _ UserRepository = &UserRepositoryMock{}

// UserRepositoryMock is a mock implementation of UserRepository.
//
//	func TestSomethingThatUsesUserRepository(t *testing.T) {
//
//		// make and configure a mocked UserRepository
//		mockedUserRepository := &UserRepositoryMock{
//			DeleteUserByUserNameFunc: func(ctx context.Context, db store.Queryer, userName string) error {
//				panic("mock out the DeleteUserByUserName method")
//			},
//			GetUserByUserNameFunc: func(ctx context.Context, db store.Queryer, userName string) (*model.User, error) {
//				panic("mock out the GetUserByUserName method")
//			},
//			RegisterUserFunc: func(ctx context.Context, db store.Queryer, u *model.User) error {
//				panic("mock out the RegisterUser method")
//			},
//			UpdateUserFunc: func(ctx context.Context, db store.Queryer, u *model.User) error {
//				panic("mock out the UpdateUser method")
//			},
//			UserExistsByUserNameFunc: func(ctx context.Context, db store.Queryer, userName string) (bool, error) {
//				panic("mock out the UserExistsByUserName method")
//			},
//			WithTransactionFunc: func(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error {
//				panic("mock out the WithTransaction method")
//			},
//		}
//
//		// use mockedUserRepository in code that requires UserRepository
//		// and then make assertions.
//
//	}
type UserRepositoryMock struct {
	// DeleteUserByUserNameFunc mocks the DeleteUserByUserName method.
	DeleteUserByUserNameFunc func(ctx context.Context, db store.Queryer, userName string) error

	// GetUserByUserNameFunc mocks the GetUserByUserName method.
	GetUserByUserNameFunc func(ctx context.Context, db store.Queryer, userName string) (*model.User, error)

	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(ctx context.Context, db store.Queryer, u *model.User) error

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(ctx context.Context, db store.Queryer, u *model.User) error

	// UserExistsByUserNameFunc mocks the UserExistsByUserName method.
	UserExistsByUserNameFunc func(ctx context.Context, db store.Queryer, userName string) (bool, error)

	// WithTransactionFunc mocks the WithTransaction method.
	WithTransactionFunc func(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error

	// calls tracks calls to the methods.
	calls struct {
		// DeleteUserByUserName holds details about calls to the DeleteUserByUserName method.
		DeleteUserByUserName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
			// UserName is the userName argument value.
			UserName string
		}
		// GetUserByUserName holds details about calls to the GetUserByUserName method.
		GetUserByUserName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
			// UserName is the userName argument value.
			UserName string
		}
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
			// U is the u argument value.
			U *model.User
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
			// U is the u argument value.
			U *model.User
		}
		// UserExistsByUserName holds details about calls to the UserExistsByUserName method.
		UserExistsByUserName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
			// UserName is the userName argument value.
			UserName string
		}
		// WithTransaction holds details about calls to the WithTransaction method.
		WithTransaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Beginner
			// F is the f argument value.
			F func(tx *sqlx.Tx) error
		}
	}
	lockDeleteUserByUserName sync.RWMutex
	lockGetUserByUserName    sync.RWMutex
	lockRegisterUser         sync.RWMutex
	lockUpdateUser           sync.RWMutex
	lockUserExistsByUserName sync.RWMutex
	lockWithTransaction      sync.RWMutex
}

// DeleteUserByUserName calls DeleteUserByUserNameFunc.
func (mock *UserRepositoryMock) DeleteUserByUserName(ctx context.Context, db store.Queryer, userName string) error {
	if mock.DeleteUserByUserNameFunc == nil {
		panic("UserRepositoryMock.DeleteUserByUserNameFunc: method is nil but UserRepository.DeleteUserByUserName was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Db       store.Queryer
		UserName string
	}{
		Ctx:      ctx,
		Db:       db,
		UserName: userName,
	}
	mock.lockDeleteUserByUserName.Lock()
	mock.calls.DeleteUserByUserName = append(mock.calls.DeleteUserByUserName, callInfo)
	mock.lockDeleteUserByUserName.Unlock()
	return mock.DeleteUserByUserNameFunc(ctx, db, userName)
}

// DeleteUserByUserNameCalls gets all the calls that were made to DeleteUserByUserName.
// Check the length with:
//
//	len(mockedUserRepository.DeleteUserByUserNameCalls())
func (mock *UserRepositoryMock) DeleteUserByUserNameCalls() []struct {
	Ctx      context.Context
	Db       store.Queryer
	UserName string
} {
	var calls []struct {
		Ctx      context.Context
		Db       store.Queryer
		UserName string
	}
	mock.lockDeleteUserByUserName.RLock()
	calls = mock.calls.DeleteUserByUserName
	mock.lockDeleteUserByUserName.RUnlock()
	return calls
}

// GetUserByUserName calls GetUserByUserNameFunc.
func (mock *UserRepositoryMock) GetUserByUserName(ctx context.Context, db store.Queryer, userName string) (*model.User, error) {
	if mock.GetUserByUserNameFunc == nil {
		panic("UserRepositoryMock.GetUserByUserNameFunc: method is nil but UserRepository.GetUserByUserName was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Db       store.Queryer
		UserName string
	}{
		Ctx:      ctx,
		Db:       db,
		UserName: userName,
	}
	mock.lockGetUserByUserName.Lock()
	mock.calls.GetUserByUserName = append(mock.calls.GetUserByUserName, callInfo)
	mock.lockGetUserByUserName.Unlock()
	return mock.GetUserByUserNameFunc(ctx, db, userName)
}

// GetUserByUserNameCalls gets all the calls that were made to GetUserByUserName.
// Check the length with:
//
//	len(mockedUserRepository.GetUserByUserNameCalls())
func (mock *UserRepositoryMock) GetUserByUserNameCalls() []struct {
	Ctx      context.Context
	Db       store.Queryer
	UserName string
} {
	var calls []struct {
		Ctx      context.Context
		Db       store.Queryer
		UserName string
	}
	mock.lockGetUserByUserName.RLock()
	calls = mock.calls.GetUserByUserName
	mock.lockGetUserByUserName.RUnlock()
	return calls
}

// RegisterUser calls RegisterUserFunc.
func (mock *UserRepositoryMock) RegisterUser(ctx context.Context, db store.Queryer, u *model.User) error {
	if mock.RegisterUserFunc == nil {
		panic("UserRepositoryMock.RegisterUserFunc: method is nil but UserRepository.RegisterUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Queryer
		U   *model.User
	}{
		Ctx: ctx,
		Db:  db,
		U:   u,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(ctx, db, u)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedUserRepository.RegisterUserCalls())
func (mock *UserRepositoryMock) RegisterUserCalls() []struct {
	Ctx context.Context
	Db  store.Queryer
	U   *model.User
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Queryer
		U   *model.User
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *UserRepositoryMock) UpdateUser(ctx context.Context, db store.Queryer, u *model.User) error {
	if mock.UpdateUserFunc == nil {
		panic("UserRepositoryMock.UpdateUserFunc: method is nil but UserRepository.UpdateUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Queryer
		U   *model.User
	}{
		Ctx: ctx,
		Db:  db,
		U:   u,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(ctx, db, u)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//
//	len(mockedUserRepository.UpdateUserCalls())
func (mock *UserRepositoryMock) UpdateUserCalls() []struct {
	Ctx context.Context
	Db  store.Queryer
	U   *model.User
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Queryer
		U   *model.User
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}

// UserExistsByUserName calls UserExistsByUserNameFunc.
func (mock *UserRepositoryMock) UserExistsByUserName(ctx context.Context, db store.Queryer, userName string) (bool, error) {
	if mock.UserExistsByUserNameFunc == nil {
		panic("UserRepositoryMock.UserExistsByUserNameFunc: method is nil but UserRepository.UserExistsByUserName was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Db       store.Queryer
		UserName string
	}{
		Ctx:      ctx,
		Db:       db,
		UserName: userName,
	}
	mock.lockUserExistsByUserName.Lock()
	mock.calls.UserExistsByUserName = append(mock.calls.UserExistsByUserName, callInfo)
	mock.lockUserExistsByUserName.Unlock()
	return mock.UserExistsByUserNameFunc(ctx, db, userName)
}

// UserExistsByUserNameCalls gets all the calls that were made to UserExistsByUserName.
// Check the length with:
//
//	len(mockedUserRepository.UserExistsByUserNameCalls())
func (mock *UserRepositoryMock) UserExistsByUserNameCalls() []struct {
	Ctx      context.Context
	Db       store.Queryer
	UserName string
} {
	var calls []struct {
		Ctx      context.Context
		Db       store.Queryer
		UserName string
	}
	mock.lockUserExistsByUserName.RLock()
	calls = mock.calls.UserExistsByUserName
	mock.lockUserExistsByUserName.RUnlock()
	return calls
}

// WithTransaction calls WithTransactionFunc.
func (mock *UserRepositoryMock) WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error {
	if mock.WithTransactionFunc == nil {
		panic("UserRepositoryMock.WithTransactionFunc: method is nil but UserRepository.WithTransaction was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Beginner
		F   func(tx *sqlx.Tx) error
	}{
		Ctx: ctx,
		Db:  db,
		F:   f,
	}
	mock.lockWithTransaction.Lock()
	mock.calls.WithTransaction = append(mock.calls.WithTransaction, callInfo)
	mock.lockWithTransaction.Unlock()
	return mock.WithTransactionFunc(ctx, db, f)
}

// WithTransactionCalls gets all the calls that were made to WithTransaction.
// Check the length with:
//
//	len(mockedUserRepository.WithTransactionCalls())
func (mock *UserRepositoryMock) WithTransactionCalls() []struct {
	Ctx context.Context
	Db  store.Beginner
	F   func(tx *sqlx.Tx) error
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Beginner
		F   func(tx *sqlx.Tx) error
	}
	mock.lockWithTransaction.RLock()
	calls = mock.calls.WithTransaction
	mock.lockWithTransaction.RUnlock()
	return calls
}

// Ensure, that AuthMock does implement Auth.
// If this is not the case, regenerate this file with moq.
var _ Auth = &AuthMock{}

// AuthMock is a mock implementation of Auth.
//
//	func TestSomethingThatUsesAuth(t *testing.T) {
//
//		// make and configure a mocked Auth
//		mockedAuth := &AuthMock{
//			SignUpFunc: func(ctx context.Context, email string, password string) (string, error) {
//				panic("mock out the SignUp method")
//			},
//		}
//
//		// use mockedAuth in code that requires Auth
//		// and then make assertions.
//
//	}
type AuthMock struct {
	// SignUpFunc mocks the SignUp method.
	SignUpFunc func(ctx context.Context, email string, password string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// SignUp holds details about calls to the SignUp method.
		SignUp []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
			// Password is the password argument value.
			Password string
		}
	}
	lockSignUp sync.RWMutex
}

// SignUp calls SignUpFunc.
func (mock *AuthMock) SignUp(ctx context.Context, email string, password string) (string, error) {
	if mock.SignUpFunc == nil {
		panic("AuthMock.SignUpFunc: method is nil but Auth.SignUp was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Email    string
		Password string
	}{
		Ctx:      ctx,
		Email:    email,
		Password: password,
	}
	mock.lockSignUp.Lock()
	mock.calls.SignUp = append(mock.calls.SignUp, callInfo)
	mock.lockSignUp.Unlock()
	return mock.SignUpFunc(ctx, email, password)
}

// SignUpCalls gets all the calls that were made to SignUp.
// Check the length with:
//
//	len(mockedAuth.SignUpCalls())
func (mock *AuthMock) SignUpCalls() []struct {
	Ctx      context.Context
	Email    string
	Password string
} {
	var calls []struct {
		Ctx      context.Context
		Email    string
		Password string
	}
	mock.lockSignUp.RLock()
	calls = mock.calls.SignUp
	mock.lockSignUp.RUnlock()
	return calls
}
