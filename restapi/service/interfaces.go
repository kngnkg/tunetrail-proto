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

type Transactioner interface {
	WithTransaction(ctx context.Context, db store.Beginner, f func(tx *sqlx.Tx) error) error
}

type LikeRepository interface {
	Transactioner
	AddLike(ctx context.Context, db store.Execer, userId model.UserID, postId string) error
	DeleteLike(ctx context.Context, db store.Execer, userId model.UserID, postId string) error
}

type FollowRepository interface {
	Transactioner
	AddFollow(ctx context.Context, db store.Execer, userId, follweeUserId model.UserID) error
	DeleteFollow(ctx context.Context, db store.Execer, userId, follweeUserId model.UserID) error
	GetFolloweesByUserId(ctx context.Context, db store.Queryer, signedInUserId model.UserID) ([]*model.User, error)
	GetFollowersByUserId(ctx context.Context, db store.Queryer, signedInUserId model.UserID) ([]*model.User, error)
}

type UserRepository interface {
	Transactioner
	FollowRepository
	RegisterUser(ctx context.Context, db store.Execer, u *model.User) error
	UpdateUser(ctx context.Context, db store.Execer, u *model.User) error
	DeleteUserByUserName(ctx context.Context, db store.Execer, userName string) error
	GetUserByUserId(ctx context.Context, db store.Queryer, id model.UserID) (*model.User, error)
	GetUserByUserName(ctx context.Context, db store.Queryer, userName string) (*model.User, error)
	GetUserByUserNameWithFollowInfo(ctx context.Context, db store.Queryer, userName string, signedInUserId model.UserID) (*model.User, error)
	UserExistsByUserName(ctx context.Context, db store.Queryer, userName string) (bool, error)
}

type PostRepository interface {
	Transactioner
	FollowRepository
	AddPost(ctx context.Context, db store.Queryer, p *model.Post) (string, error)
	AddReplyRelation(ctx context.Context, db store.Execer, postId, parentId string) error
	DeletePost(ctx context.Context, db store.Execer, postId string) error
	GetPostById(ctx context.Context, db store.Queryer, postId string, signedInUserId model.UserID) (*model.Post, error)
	GetPostsByUserId(ctx context.Context, db store.Queryer, userId model.UserID, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
	GetPostsByUserIds(ctx context.Context, db store.Queryer, userIds []model.UserID, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
	GetLikedPostsByUserId(ctx context.Context, db store.Queryer, userId model.UserID, signedInUserId model.UserID, pagenation *model.Pagenation) (*model.Timeline, error)
	GetReplies(ctx context.Context, db store.Queryer, parentPostId string, pagenation *model.Pagenation) (*model.Timeline, error)
}

type HealthRepository interface {
	Ping(ctx context.Context, db store.Queryer) error
}
