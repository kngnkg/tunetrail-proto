package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kwtryo/tunetrail/api/model"
	"github.com/kwtryo/tunetrail/api/store"
)

var (
	// ユーザー名が既に存在する
	ErrUserNameAlreadyExists = errors.New("service: user name already exists")
	// メールアドレスが既に存在する
	ErrEmailAlreadyExists = errors.New("service: email already exists")
	// ユーザーが存在しない
	ErrUserNotFound = errors.New("service: user not found")
)

type UserRepository interface {
	// WithTransactionはトランザクションを実行する
	WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error
	// RegisterUserはユーザーを登録する
	RegisterUser(ctx context.Context, db store.Queryer, u *model.User) error
	// UserExistsByUserNameはユーザー名が既に存在するかどうかを返す
	UserExistsByUserName(ctx context.Context, db store.Queryer, userName string) (bool, error)
	// UserExistsByEmailはメールアドレスが既に存在するかどうかを返す
	UserExistsByEmail(ctx context.Context, db store.Queryer, email string) (bool, error)
	// GetUserByUserNameはユーザー名からユーザーを取得する
	GetUserByUserName(ctx context.Context, db store.Queryer, userName string) (*model.User, error)
	// UpdateUserはユーザーを更新する
	UpdateUser(ctx context.Context, db store.Queryer, u *model.User) error
	// DeleteUserByUserNameはユーザー名からユーザーを削除する
	DeleteUserByUserName(ctx context.Context, db store.Queryer, userName string) error
}

type UserService struct {
	DB   store.Beginner
	Repo UserRepository
}

// RegisterUserはユーザーを登録する
func (us *UserService) RegisterUser(
	ctx context.Context, userName, name, password, email string,
) (*model.User, error) {
	registeredUser := &model.User{}
	// トランザクション開始
	err := us.Repo.WithTransaction(ctx, us.DB, func(tx *sqlx.Tx) error {
		// ユーザー名が既に存在するかどうかを確認
		exists, err := us.Repo.UserExistsByUserName(ctx, tx, userName)
		if err != nil {
			return err
		}
		if exists {
			return ErrUserNameAlreadyExists
		}
		// メールアドレスが既に存在するかどうかを確認
		exists, err = us.Repo.UserExistsByEmail(ctx, tx, email)
		if err != nil {
			return err
		}
		if exists {
			return ErrEmailAlreadyExists
		}

		u := &model.User{
			UserName: userName,
			Name:     name,
			Password: password,
			Email:    email,
			IconUrl:  "",
			Bio:      "",
		}
		if err = us.Repo.RegisterUser(ctx, tx, u); err != nil {
			return err
		}
		registeredUser = u
		return nil
	})
	if err != nil {
		return nil, err
	}

	return registeredUser, nil
}

// GetUserByUserNameはユーザー名からユーザーを取得する
func (us *UserService) GetUserByUserName(ctx context.Context, userName string) (*model.User, error) {
	var u *model.User
	err := us.Repo.WithTransaction(ctx, us.DB, func(tx *sqlx.Tx) error {
		got, err := us.Repo.GetUserByUserName(ctx, tx, userName)
		if err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				return fmt.Errorf("%w: userName=%v: %w", ErrUserNotFound, userName, err)
			}
			return err
		}
		u = got
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateUserはユーザーを更新する
func (us *UserService) UpdateUser(ctx context.Context, u *model.User) error {
	err := us.Repo.WithTransaction(ctx, us.DB, func(tx *sqlx.Tx) error {
		// ユーザー名が既に存在するかどうかを確認
		exists, err := us.Repo.UserExistsByUserName(ctx, tx, u.UserName)
		if err != nil {
			return err
		}
		if exists {
			return ErrUserNameAlreadyExists
		}
		// メールアドレスが既に存在するかどうかを確認
		exists, err = us.Repo.UserExistsByEmail(ctx, tx, u.Email)
		if err != nil {
			return err
		}
		if exists {
			return ErrEmailAlreadyExists
		}

		// ユーザーを更新する
		if err := us.Repo.UpdateUser(ctx, tx, u); err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				return fmt.Errorf("%w: userName=%v: %w", ErrUserNotFound, u.UserName, err)
			}
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserByUserNameはユーザー名からユーザーを削除する
func (us *UserService) DeleteUserByUserName(ctx context.Context, userName string) error {
	err := us.Repo.WithTransaction(ctx, us.DB, func(tx *sqlx.Tx) error {
		if err := us.Repo.DeleteUserByUserName(ctx, tx, userName); err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				return fmt.Errorf("%w: userName=%v: %w", ErrUserNotFound, userName, err)
			}
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
