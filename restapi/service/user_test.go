package service

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
	"github.com/kngnkg/tunetrail/restapi/testutil/fixture"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	us         *UserService        // テスト対象のサービス
	fc         *clock.FixedClocker // テスト用の時刻を固定する
	dummyUsers []*model.User       // テスト用のダミーユーザー
}

func TestUserServiceTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserServiceTestSuite))
}

func (s *UserServiceTestSuite) SetupTest() {
	fc := &clock.FixedClocker{}

	// テスト用のダミーユーザーを作成
	// これらのユーザーだけがDBに存在することを想定している
	dummyUsers := fixture.CreateUsers(2)

	moqRepo := s.setupRepoMock()

	s.us = &UserService{
		DB:   &DBConnectionMock{},
		Repo: moqRepo,
	}
	s.fc = fc
	s.dummyUsers = dummyUsers
}

func (s *UserServiceTestSuite) TestGetUserByUserName() {
	type args struct {
		ctx            context.Context
		userName       string
		signedInUserId model.UserID
	}

	tests := []struct {
		name     string
		args     args
		wantUser *model.User
		wantErr  error
	}{
		{
			"ok",
			args{
				ctx:            context.Background(),
				userName:       s.dummyUsers[0].UserName,
				signedInUserId: s.dummyUsers[0].Id,
			},
			s.dummyUsers[0],
			nil,
		},
		{
			// 存在しないユーザー名の場合
			"errUserNotFound",
			args{
				ctx:      context.Background(),
				userName: "notfound",
			},
			nil,
			ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.us.GetUserByUserName(tt.args.ctx, tt.args.userName, tt.args.signedInUserId)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("UserService.GetUserByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr == nil {
				assert.Equal(s.T(), tt.wantUser, got)
			}
		})
	}
}

func (s *UserServiceTestSuite) TestUpdateUser() {
	type args struct {
		ctx  context.Context
		data *model.UserUpdateData
	}

	// ダミーユーザー1を更新することを想定
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			"ok",
			args{
				ctx: context.Background(),
				data: &model.UserUpdateData{
					Id:       s.dummyUsers[0].Id,
					UserName: "update",
					Name:     "test",
					IconUrl:  "https://example.com/icon.png",
					Bio:      "test",
					Password: "test",
					Email:    "email@example.com",
				},
			},
			nil,
		},
		// IDが存在しない場合
		{
			"errUserNotFound",
			args{
				ctx: context.Background(),
				data: &model.UserUpdateData{
					Id:       fixture.NewUserId(),
					UserName: "update",
					Name:     "test",
					IconUrl:  "https://example.com/icon.png",
					Bio:      "test",
					Password: "test",
					Email:    "email@example.com",
				},
			},
			ErrUserNotFound,
		},
		{
			// ユーザー名が重複している場合
			"errUserNameAlreadyExists",
			args{
				ctx: context.Background(),
				data: &model.UserUpdateData{
					Id:       s.dummyUsers[0].Id,
					UserName: s.dummyUsers[1].UserName, // ダミーユーザー2と同じユーザー名
					Name:     "test",
					IconUrl:  "https://example.com/icon.png",
					Bio:      "test",
					Password: "test",
					Email:    "email@example.com",
				},
			},
			ErrUserNameAlreadyExists,
		},
		// {
		// 	// メールアドレスが重複している場合
		// 	"errEmailNameAlreadyExists",
		// 	args{
		// 		ctx: context.Background(),
		// 		data: &model.UserUpdateData{
		// 			Id:       DUMMY_1_USER_ID,
		// 			UserName: DUMMY_2_USERNAME,
		// 			Name:     "test",
		// 			IconUrl:  "https://example.com/icon.png",
		// 			Bio:      "test",
		// 			Password: VALID_PASSWORD,
		// 			Email:    DUMMY_2_EMAIL,
		// 		},
		// 	},
		// 	ErrEmailAlreadyExists,
		// },
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := s.us.UpdateUser(tt.args.ctx, tt.args.data)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("UserService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *UserServiceTestSuite) TestDeleteUserByUserName() {
	type args struct {
		ctx      context.Context
		userName string
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			"ok",
			args{
				ctx:      context.Background(),
				userName: s.dummyUsers[0].UserName,
			},
			nil,
		},
		{
			// 存在しないユーザー名の場合
			"errUserNotFound",
			args{
				ctx:      context.Background(),
				userName: "notfound",
			},
			ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := s.us.DeleteUserByUserName(tt.args.ctx, tt.args.userName)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("UserService.DeleteUserByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *UserServiceTestSuite) setupRepoMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		WithTransactionFunc: func(ctx context.Context, db store.Beginner, fn func(tx *sqlx.Tx) error) error {
			tx := &sqlx.Tx{}
			err := fn(tx)
			if err != nil {
				return fmt.Errorf("failed to execute mock transaction: %w", err)
			}
			return nil
		},
		UserExistsByUserNameFunc: func(ctx context.Context, db store.Queryer, userName string) (bool, error) {
			for _, u := range s.dummyUsers {
				if userName == u.UserName {
					return true, nil
				}
			}
			return false, nil
		},
		GetUserByUserNameFunc: func(ctx context.Context, db store.Queryer, userName string) (*model.User, error) {
			for _, u := range s.dummyUsers {
				if userName == u.UserName {
					return u, nil
				}
			}
			return nil, store.ErrUserNotFound
		},
		GetUserByUserNameWithFollowInfoFunc: func(ctx context.Context, db store.Queryer, userName string, signedInUserId model.UserID) (*model.User, error) {
			for _, u := range s.dummyUsers {
				if userName == u.UserName {
					return u, nil
				}
			}
			return nil, store.ErrUserNotFound
		},
		// ダミーユーザー1を更新する場合を想定
		UpdateUserFunc: func(ctx context.Context, db store.Execer, u *model.User) error {
			if u.Id != s.dummyUsers[0].Id && u.Id != s.dummyUsers[1].Id {
				return store.ErrUserNotFound
			}
			if u.UserName == s.dummyUsers[1].UserName {
				return store.ErrUserNameAlreadyExists
			}
			u.UpdatedAt = s.fc.Now()
			return nil
		},
		DeleteUserByUserNameFunc: func(ctx context.Context, db store.Execer, userName string) error {
			for _, u := range s.dummyUsers {
				if userName == u.UserName {
					return nil
				}
			}
			return store.ErrUserNotFound
		},
	}
}
