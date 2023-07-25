package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

type Auth interface {
	SignUp(ctx context.Context, email, password string) (string, error)
	ConfirmSignUp(ctx context.Context, cognitoUserName, code string) error
}

type UserRepository interface {
	// WithTransactionはトランザクションを実行する
	WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error
	// RegisterUserはユーザーを登録する
	// RegisterUserInfo
	RegisterUser(ctx context.Context, db store.Queryer, u *model.User) error
	// UserExistsByUserNameはユーザー名が既に存在するかどうかを返す
	UserExistsByUserName(ctx context.Context, db store.Queryer, userName string) (bool, error)
	// GetUserByUserNameはユーザー名からユーザーを取得する
	GetUserByUserName(ctx context.Context, db store.Queryer, userName string) (*model.User, error)
	// UpdateUserはユーザーを更新する
	UpdateUser(ctx context.Context, db store.Queryer, u *model.User) error
	// DeleteUserByUserNameはユーザー名からユーザーを削除する
	DeleteUserByUserName(ctx context.Context, db store.Queryer, userName string) error
}

type HealthRepository interface {
	Ping(ctx context.Context, db store.Queryer) error
}
