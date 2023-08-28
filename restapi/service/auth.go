package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/auth"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type AuthService struct {
	DB    store.DBConnection
	Repo  UserRepository
	Auth  Auth
	JWTer JWTer
}

// RegisterUserはユーザーを登録する
func (as *AuthService) RegisterUser(ctx context.Context, data *model.UserRegistrationData) (*model.User, error) {
	var maxAttempt = 3
	var userId model.UserID

	for i := 0; i < maxAttempt; i++ {
		userId = model.NewUserId() // ユーザーIDを生成

		if err := as.Auth.SignUp(ctx, userId, data.Email, data.Password); err != nil {
			if errors.Is(err, auth.ErrUserIdAlreadyExists) {
				continue // ユーザーIDが既に存在する場合はリトライ
			}
			if errors.Is(err, auth.ErrEmailAlreadyExists) {
				return nil, fmt.Errorf("%w: %w", ErrEmailAlreadyExists, err)
			}
			return nil, err
		}

		if i == maxAttempt-1 {
			return nil, fmt.Errorf("maximum retry attempts exceeded")
		}

		break
	}

	registeredUser := &model.User{}
	// トランザクション開始
	err := as.Repo.WithTransaction(ctx, as.DB, func(tx *sqlx.Tx) error {
		// ユーザー名が既に存在するかどうかを確認
		exists, err := as.Repo.UserExistsByUserName(ctx, tx, data.UserName)
		if err != nil {
			return err
		}
		if exists {
			return ErrUserNameAlreadyExists
		}

		u := &model.User{
			Id:       userId,
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
		// TODO: Authのユーザー登録をロールバックする
		return nil, err
	}

	return registeredUser, nil
}

func (as *AuthService) ConfirmEmail(ctx context.Context, userName, code string) error {
	// TODO: 後でリファクタリングする
	user, err := as.Repo.GetUserByUserName(ctx, as.DB, userName)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			return fmt.Errorf("%w: userName=%v: %w", ErrUserNotFound, userName, err)
		}
		return err
	}

	if err := as.Auth.ConfirmSignUp(ctx, user.Id, code); err != nil {
		if errors.Is(err, auth.ErrCodeMismatch) {
			return fmt.Errorf("%w: %w", ErrCodeMismatch, err)
		}
		if errors.Is(err, auth.ErrCodeExpired) {
			return fmt.Errorf("%w: %w", ErrCodeExpired, err)
		}
		if errors.Is(err, auth.ErrEmailAlreadyExists) {
			return fmt.Errorf("%w: %w", ErrEmailAlreadyExists, err)
		}
		return err
	}

	return nil
}

func (as *AuthService) SignIn(ctx context.Context, data *model.UserSignInData) (*model.Tokens, error) {
	// メールアドレスでサインインする場合
	if data.Email != "" {
		tokens, err := as.Auth.SignIn(ctx, data.Email, data.Password)
		if err != nil {
			if errors.Is(err, auth.ErrUserNotFound) {
				return nil, fmt.Errorf("%w: %w", ErrUserNotFound, err)
			}
			if errors.Is(err, auth.ErrNotAuthorized) {
				return nil, fmt.Errorf("%w: %w", ErrNotAuthorized, err)
			}
			return nil, err
		}
		return tokens, nil
	}

	// ユーザー名でサインインする場合
	// TODO: 後でリファクタリングする
	user, err := as.Repo.GetUserByUserName(ctx, as.DB, data.UserName)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			return nil, fmt.Errorf("%w: userName=%v: %w", ErrUserNotFound, data.UserName, err)
		}
		return nil, err
	}

	tokens, err := as.Auth.SignIn(ctx, string(user.Id), data.Password)
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, fmt.Errorf("%w: %w", ErrUserNotFound, err)
		}
		if errors.Is(err, auth.ErrNotAuthorized) {
			return nil, fmt.Errorf("%w: %w", ErrNotAuthorized, err)
		}
		return nil, err
	}
	return tokens, nil
}

func (as *AuthService) GetSignedInUser(ctx context.Context, idToken string) (*model.User, error) {
	parsed, err := as.JWTer.ParseIdToken(ctx, idToken)
	if err != nil {
		// TODO: エラー処理
		return nil, err
	}

	return as.Repo.GetUserByUserId(ctx, as.DB, parsed.Id)
}

func (as *AuthService) RefreshToken(ctx context.Context, idToken, refreshToken string) (*model.Tokens, error) {
	ai, err := as.JWTer.ParseIdToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	tokens, err := as.Auth.RefreshToken(ctx, string(ai.Id), refreshToken)
	if err != nil {
		if errors.Is(err, auth.ErrNotAuthorized) {
			return nil, fmt.Errorf("%w: %w", ErrNotAuthorized, err)
		}
		return nil, err
	}
	return tokens, nil
}
