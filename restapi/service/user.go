package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type UserService struct {
	DB   store.DBConnection
	Repo UserRepository
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
func (us *UserService) UpdateUser(ctx context.Context, data *model.UserUpdateData) error {
	err := us.Repo.WithTransaction(ctx, us.DB, func(tx *sqlx.Tx) error {
		// ユーザー名が既に存在するかどうかを確認
		exists, err := us.Repo.UserExistsByUserName(ctx, tx, data.UserName)
		if err != nil {
			return err
		}
		if exists {
			return ErrUserNameAlreadyExists
		}
		// メールアドレスが既に存在するかどうかを確認

		// メールアドレス、パスワードを更新する処理(Cognito)

		user := &model.User{
			Id:       data.Id,
			UserName: data.UserName,
			Name:     data.Name,
			IconUrl:  data.IconUrl,
			Bio:      data.Bio,
		}
		// ユーザーを更新する
		if err := us.Repo.UpdateUser(ctx, tx, user); err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				return fmt.Errorf("%w: userName=%v: %w", ErrUserNotFound, data.UserName, err)
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

// FollowUserはユーザーをフォローする
func (us *UserService) FollowUser(ctx context.Context, userName, follweeUserName string) error {
	err := us.Repo.WithTransaction(ctx, us.DB, func(tx *sqlx.Tx) error {
		if err := us.Repo.AddFollow(ctx, tx, userName, follweeUserName); err != nil {
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
