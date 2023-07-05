package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kwtryo/tunetrail/api/clock"
	"github.com/kwtryo/tunetrail/api/model"
	"github.com/kwtryo/tunetrail/api/store"
	"github.com/kwtryo/tunetrail/api/testutil/fixture"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	us    *UserService        // テスト対象のサービス
	fc    *clock.FixedClocker // テスト用の時刻を固定する
	dummy *model.User         // テスト用のダミーユーザー
}

func TestUserServiceTestSuite(t *testing.T) {
	t.Parallel()
	moqDB := &BeginnerMock{}
	moqRepo := &UserRepositoryMock{}
	fc := &clock.FixedClocker{}

	// 各種モック関数の設定
	moqDB.BeginTxxFunc =
		func(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
			return &sqlx.Tx{}, nil
		}

	moqRepo.WithTransactionFunc =
		func(ctx context.Context, db store.Beginner, fn func(tx *sqlx.Tx) error) error {
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

	moqRepo.RegisterUserFunc =
		func(ctx context.Context, db store.Queryer, u *model.User) error {
			// ダミーの値を設定
			u.Id = 1
			u.CreatedAt = fc.Now()
			u.UpdatedAt = fc.Now()
			return nil
		}

	moqRepo.UserExistsByUserNameFunc =
		func(ctx context.Context, db store.Queryer, userName string) (bool, error) {
			if userName == "dummy" {
				return true, nil
			}
			return false, nil
		}

	moqRepo.UserExistsByEmailFunc =
		func(ctx context.Context, db store.Queryer, email string) (bool, error) {
			if email == "dummy@example.com" {
				return true, nil
			}
			return false, nil
		}

	moqRepo.GetUserByUserNameFunc =
		func(ctx context.Context, db store.Queryer, userName string) (*model.User, error) {
			if userName == "dummy" {
				u := fixture.User(&model.User{
					Id:       1,
					UserName: "dummy",
					Name:     "dummy",
					Password: "dummy",
					Email:    "dummy@example.com",
					IconUrl:  "https://example.com/icon.png",
					Bio:      "dummy",
					// タイムスタンプを固定する
					CreatedAt: fc.Now(),
					UpdatedAt: fc.Now(),
				})
				return u, nil
			}
			return nil, store.ErrUserNotFound
		}

	moqRepo.UpdateUserFunc =
		func(ctx context.Context, db store.Queryer, u *model.User) error {
			if u.Id != 1 {
				return store.ErrUserNotFound
			}
			if u.UserName == "dummy" {
				return store.ErrUserNameAlreadyExists
			}
			if u.Email == "dummy@example.com" {
				return store.ErrEmailAlreadyExists
			}
			u.UpdatedAt = fc.Now()
			return nil
		}

	moqRepo.DeleteUserByUserNameFunc =
		func(ctx context.Context, db store.Queryer, userName string) error {
			if userName == "dummy" {
				return nil
			}
			return store.ErrUserNotFound
		}

	suite.Run(t, &UserServiceTestSuite{
		us: &UserService{
			DB:   moqDB,
			Repo: moqRepo,
		},
		fc: fc,
	})
}

func (s *UserServiceTestSuite) SetupTest() {
	// テスト用のダミーユーザー
	s.dummy = fixture.User(&model.User{
		UserName: "dummy",
		Name:     "dummy",
		Password: "dummy",
		Email:    "dummy@example.com",
		IconUrl:  "https://example.com/icon.png",
		Bio:      "dummy",
		// タイムスタンプを固定する
		CreatedAt: s.fc.Now(),
		UpdatedAt: s.fc.Now(),
	})
}

func (s *UserServiceTestSuite) TestRegisterUser() {
	type args struct {
		ctx      context.Context
		userName string
		name     string
		password string
		email    string
		iconUrl  string
		Bio      string
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
				password: "test",
				email:    "test@example.com",
			},
			&model.User{
				Id:        1,
				UserName:  "test",
				Name:      "test",
				Password:  "test",
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
				userName: "dummy", // ダミーユーザーのユーザー名と重複させる
				name:     "test",
				password: "test",
				email:    "test@example.com",
			},
			nil,
			ErrUserNameAlreadyExists,
		},
		{
			// メールアドレスが重複している場合
			"errEmailNameAlreadyExists",
			args{
				ctx:      context.Background(),
				userName: "test",
				name:     "test",
				password: "test",
				email:    "dummy@example.com", // ダミーユーザーのメールアドレスと重複させる
			},
			nil,
			ErrEmailAlreadyExists,
		},
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
				userName: "dummy",
			},
			&model.User{
				Id:        1,
				UserName:  "dummy",
				Name:      "dummy",
				Password:  "dummy",
				Email:     "dummy@example.com",
				IconUrl:   "https://example.com/icon.png",
				Bio:       "dummy",
				CreatedAt: s.fc.Now(),
				UpdatedAt: s.fc.Now(),
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
					Id:        1,
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
					Id:        999,
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
					Id:       1,
					UserName: "dummy", // ダミーユーザーのユーザー名と重複させる
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
					Id:    1,
					Email: "dummy@example.com", // ダミーユーザーのメールアドレスと重複させる
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
				userName: "dummy",
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
