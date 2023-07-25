package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/auth"
	"github.com/kngnkg/tunetrail/restapi/model"
)

type AuthService struct {
	DB   *sqlx.DB
	Repo UserRepository
	Auth Auth
}

// type Token struct{}

// RegisterUserはユーザーを登録する
func (as *AuthService) RegisterUser(ctx context.Context, data *model.UserRegistrationData) (*model.User, error) {
	id, err := as.Auth.SignUp(ctx, data.Email, data.Password)
	if err != nil {
		if errors.Is(err, auth.ErrEmailAlreadyExists) {
			return nil, fmt.Errorf("%w: %w", ErrEmailAlreadyExists, err)
		}
		return nil, err
	}

	registeredUser := &model.User{}
	// トランザクション開始
	err = as.Repo.WithTransaction(ctx, as.DB, func(tx *sqlx.Tx) error {
		// ユーザー名が既に存在するかどうかを確認
		exists, err := as.Repo.UserExistsByUserName(ctx, tx, data.UserName)
		if err != nil {
			return err
		}
		if exists {
			return ErrUserNameAlreadyExists
		}

		u := &model.User{
			Id:       id,
			UserName: data.UserName,
			Name:     data.Name,
			IconUrl:  "",
			Bio:      "",
		}
		if err = as.Repo.RegisterUser(ctx, tx, u); err != nil {
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
