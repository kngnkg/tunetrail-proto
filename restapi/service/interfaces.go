package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type Auth interface {
	SignUp(ctx context.Context, userId model.UserID, email, password string) error
	ConfirmSignUp(ctx context.Context, userId model.UserID, code string) error
	SignIn(ctx context.Context, userIdentifier, password string) (*model.Tokens, error)
	RefreshToken(ctx context.Context, userIdentifier, refreshToken string) (*model.Tokens, error)
}

type UserRepository interface {
	// WithTransactionはトランザクションを実行する
	WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error
	// UserExistsByUserNameはユーザー名が既に存在するかどうかを返す
	UserExistsByUserName(ctx context.Context, db store.Queryer, userName string) (bool, error)
	// GetUserByUserNameはユーザー名からユーザーを取得する
	GetUserByUserName(ctx context.Context, db store.Queryer, userName string) (*model.User, error)
	// RegisterUserはユーザーを登録する
	// RegisterUserInfo
	RegisterUser(ctx context.Context, db store.Execer, u *model.User) error
	// UpdateUserはユーザーを更新する
	UpdateUser(ctx context.Context, db store.Execer, u *model.User) error
	// DeleteUserByUserNameはユーザー名からユーザーを削除する
	DeleteUserByUserName(ctx context.Context, db store.Execer, userName string) error
}

type HealthRepository interface {
	Ping(ctx context.Context, db store.Queryer) error
}
