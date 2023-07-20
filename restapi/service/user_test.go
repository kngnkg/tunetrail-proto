package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/auth"
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

var (
	VALID_USER_ID  = "1"
	VALID_PASSWORD = "password"

	DUMMY_1_USER_ID  = "dummy1id"
	DUMMY_1_USERNAME = "dummy1"
	DUMMY_1_EMAIL    = "dummy1@example.com"

	DUMMY_2_USER_ID  = "dummy2id"
	DUMMY_2_USERNAME = "dummy2"
	DUMMY_2_EMAIL    = "dummy2@example.com"
)

func TestUserServiceTestSuite(t *testing.T) {
	t.Parallel()
	moqDB := &BeginnerMock{}
	moqRepo := &UserRepositoryMock{}
	moqAuth := &AuthMock{}
	fc := &clock.FixedClocker{}

	// テスト用のダミーユーザーを作成
	// これらのユーザーだけがDBに存在することを想定している
	dummyUsers := []*model.User{
		fixture.User(&model.User{
			Id:       DUMMY_1_USER_ID,
			UserName: DUMMY_1_USERNAME,
			Name:     "dummy1",
			Password: VALID_PASSWORD,
			Email:    DUMMY_1_EMAIL,
			IconUrl:  "https://example.com/icon.png",
			Bio:      "dummy1",
			// タイムスタンプを固定する
			CreatedAt: fc.Now(),
			UpdatedAt: fc.Now(),
		}),
		fixture.User(&model.User{
			Id:       DUMMY_2_USER_ID,
			UserName: DUMMY_2_USERNAME,
			Name:     "dummy2",
			Password: VALID_PASSWORD,
			Email:    DUMMY_2_EMAIL,
			IconUrl:  "https://example.com/icon.png",
			Bio:      "dummy2",
			// タイムスタンプを固定する
			CreatedAt: fc.Now(),
			UpdatedAt: fc.Now(),
		}),
	}

	// 各種モック関数の設定
	moqDB.BeginTxxFunc = func(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
		return &sqlx.Tx{}, nil
	}

	moqRepo.WithTransactionFunc = func(ctx context.Context, db store.Beginner, fn func(tx *sqlx.Tx) error) error {
		tx, err := db.BeginTxx(ctx, nil)
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}
		err = fn(tx)
		if err != nil {
			return fmt.Errorf("failed to execute mock transaction: %w", err)
		}
		return nil
	}

	moqRepo.RegisterUserFunc = func(ctx context.Context, db store.Queryer, u *model.User) error {
		// ダミーの値を設定
		u.Id = "1"
		u.CreatedAt = fc.Now()
		u.UpdatedAt = fc.Now()
		return nil
	}

	moqRepo.UserExistsByUserNameFunc = func(ctx context.Context, db store.Queryer, userName string) (bool, error) {
		if userName == DUMMY_1_USERNAME || userName == DUMMY_2_USERNAME {
			return true, nil
		}
		return false, nil
	}

	moqRepo.UserExistsByEmailFunc = func(ctx context.Context, db store.Queryer, email string) (bool, error) {
		if email == DUMMY_1_EMAIL || email == DUMMY_2_EMAIL {
			return true, nil
		}
		return false, nil
	}

	moqRepo.GetUserByUserNameFunc = func(ctx context.Context, db store.Queryer, userName string) (*model.User, error) {
		if userName == DUMMY_1_USERNAME {
			return dummyUsers[0], nil
		}
		if userName == DUMMY_2_USERNAME {
			return dummyUsers[1], nil
		}
		return nil, store.ErrUserNotFound
	}

	// ダミーユーザー1を更新する場合を想定
	moqRepo.UpdateUserFunc = func(ctx context.Context, db store.Queryer, u *model.User) error {
		if u.Id != DUMMY_1_USER_ID && u.Id != DUMMY_2_USER_ID {
			return store.ErrUserNotFound
		}
		if u.UserName == DUMMY_2_USERNAME {
			return store.ErrUserNameAlreadyExists
		}
		if u.Email == DUMMY_2_EMAIL {
			return store.ErrEmailAlreadyExists
		}
		u.UpdatedAt = fc.Now()
		return nil
	}

	moqRepo.DeleteUserByUserNameFunc = func(ctx context.Context, db store.Queryer, userName string) error {
		if userName == DUMMY_1_USERNAME || userName == DUMMY_2_USERNAME {
			return nil
		}
		return store.ErrUserNotFound
	}

	moqAuth.SignUpFunc = func(ctx context.Context, email, password string) (string, error) {
		if password != VALID_PASSWORD {
			return "", auth.ErrInvalidPassword
		}
		if email == DUMMY_1_EMAIL || email == DUMMY_2_EMAIL {
			return "", auth.ErrEmailAlreadyExists
		}
		return VALID_USER_ID, nil
	}

	suite.Run(t, &UserServiceTestSuite{
		us: &UserService{
			DB:   moqDB,
			Repo: moqRepo,
			Auth: moqAuth,
		},
		fc:         fc,
		dummyUsers: dummyUsers,
	})
}

func (s *UserServiceTestSuite) SetupTest() {}

func (s *UserServiceTestSuite) TestRegisterUser() {
	type args struct {
		ctx      context.Context
		userName string
		name     string
		password string
		email    string
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
				ctx:      context.Background(),
				userName: "test",
				name:     "test",
				password: VALID_PASSWORD,
				email:    "test@example.com",
			},
			&model.User{
				Id:        VALID_USER_ID,
				UserName:  "test",
				Name:      "test",
				Password:  VALID_PASSWORD,
				Email:     "test@example.com",
				IconUrl:   "",
				Bio:       "",
				CreatedAt: s.fc.Now(),
				UpdatedAt: s.fc.Now(),
			},
			nil,
		},
		{
			// ユーザー名が重複している場合
			"errUserNameAlreadyExists",
			args{
				ctx:      context.Background(),
				userName: DUMMY_1_USERNAME, // ダミーユーザーのユーザー名と重複させる
				name:     "test",
				password: VALID_PASSWORD,
				email:    "test@example.com",
			},
			nil,
			ErrUserNameAlreadyExists,
		},
		{
			// メールアドレスが重複している場合
			"errEmailAlreadyExists",
			args{
				ctx:      context.Background(),
				userName: "test",
				name:     "test",
				password: VALID_PASSWORD,
				email:    DUMMY_1_EMAIL, // ダミーユーザーのメールアドレスと重複させる
			},
			nil,
			ErrEmailAlreadyExists,
		},
		// パスワードが不正な場合
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.us.RegisterUser(
				tt.args.ctx, tt.args.userName, tt.args.name, tt.args.password, tt.args.email,
			)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("UserService.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr == nil {
				assert.Equal(s.T(), tt.wantUser, got)
			}
		})
	}
}

func (s *UserServiceTestSuite) TestGetUserByUserName() {
	type args struct {
		ctx      context.Context
		userName string
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
				ctx:      context.Background(),
				userName: DUMMY_1_USERNAME,
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
			got, err := s.us.GetUserByUserName(tt.args.ctx, tt.args.userName)
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
		user *model.User
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
				user: fixture.User(&model.User{
					Id:        DUMMY_1_USER_ID,
					CreatedAt: s.fc.Now(),
					UpdatedAt: s.fc.Now(),
				}),
			},
			nil,
		},
		// IDが存在しない場合
		{
			"errUserNotFound",
			args{
				ctx: context.Background(),
				user: fixture.User(&model.User{
					Id:        "0",
					CreatedAt: s.fc.Now(),
					UpdatedAt: s.fc.Now(),
				}),
			},
			ErrUserNotFound,
		},
		{
			// ユーザー名が重複している場合
			"errUserNameAlreadyExists",
			args{
				ctx: context.Background(),
				user: fixture.User(&model.User{
					Id:       DUMMY_1_USER_ID,
					UserName: DUMMY_1_USERNAME,
				}),
			},
			ErrUserNameAlreadyExists,
		},
		{
			// メールアドレスが重複している場合
			"errEmailNameAlreadyExists",
			args{
				ctx: context.Background(),
				user: fixture.User(&model.User{
					Id:    DUMMY_1_USER_ID,
					Email: DUMMY_2_EMAIL,
				}),
			},
			ErrEmailAlreadyExists,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := s.us.UpdateUser(tt.args.ctx, tt.args.user)
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
				userName: DUMMY_1_USERNAME,
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
