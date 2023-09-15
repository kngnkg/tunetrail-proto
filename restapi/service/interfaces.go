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

type JWTer interface {
	ParseIdToken(ctx context.Context, idToken string) (*model.AuthInfo, error)
}

type UserRepository interface {
	WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error
	UserExistsByUserName(ctx context.Context, db store.Queryer, userName string) (bool, error)
	GetUserByUserName(ctx context.Context, db store.Queryer, userName string) (*model.User, error)
	GetUserByUserId(ctx context.Context, db store.Queryer, id model.UserID) (*model.User, error)
	RegisterUser(ctx context.Context, db store.Execer, u *model.User) error
	UpdateUser(ctx context.Context, db store.Execer, u *model.User) error
	DeleteUserByUserName(ctx context.Context, db store.Execer, userName string) error
	AddFollow(ctx context.Context, db store.Execer, userId, follweeUserId model.UserID) error
	DeleteFollow(ctx context.Context, db store.Execer, userId, follweeUserId model.UserID) error
	GetUserByUserNameWithFollowInfo(ctx context.Context, db store.Queryer, userName string, signedInUserId model.UserID) (*model.User, error)
	GetFolloweesByUserId(ctx context.Context, db store.Queryer, signedInUserId model.UserID) ([]*model.User, error)
	GetFollowersByUserId(ctx context.Context, db store.Queryer, signedInUserId model.UserID) ([]*model.User, error)
}

type PostRepository interface {
	WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error
	AddPost(ctx context.Context, db store.Queryer, p *model.Post) (string, error)
	GetPostById(ctx context.Context, db store.Queryer, postId string) (*model.Post, error)
	GetFolloweesByUserId(ctx context.Context, db store.Queryer, signedInUserId model.UserID) ([]*model.User, error)
	GetPostsByUserIdsNext(ctx context.Context, db store.Queryer, userId []model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
	GetReplies(ctx context.Context, db store.Queryer, parentPostId string, pagenation *model.Pagenation) (*model.Timeline, error)
}

type HealthRepository interface {
	Ping(ctx context.Context, db store.Queryer) error
}
