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

type AuthServiceTestSuite struct {
	suite.Suite
	as         *AuthService        // テスト対象のサービス
	fc         *clock.FixedClocker // テスト用の時刻を固定する
	dummyUsers []*model.User       // テスト用のダミーユーザー
}

func TestAuthServiceTestSuite(t *testing.T) {
	t.Parallel()
	moqDB := &DBConnectionMock{}
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

	moqRepo.RegisterUserFunc = func(ctx context.Context, db store.Execer, u *model.User) error {
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

	moqAuth.SignUpFunc = func(ctx context.Context, email, password string) (string, error) {
		if password != VALID_PASSWORD {
			return "", auth.ErrInvalidPassword
		}
		if password == DUMMY_1_PASSWORD || password == DUMMY_2_PASSWORD {
			return "", auth.ErrInvalidPassword
		}
		if email == DUMMY_1_EMAIL || email == DUMMY_2_EMAIL {
			return "", auth.ErrEmailAlreadyExists
		}
		return VALID_USER_ID, nil
	}

	suite.Run(t, &AuthServiceTestSuite{
		as: &AuthService{
			DB:   moqDB,
			Repo: moqRepo,
			Auth: moqAuth,
		},
		fc:         fc,
		dummyUsers: dummyUsers,
	})
}

func (s *AuthServiceTestSuite) SetupTest() {}

func (s *AuthServiceTestSuite) TestRegisterUser() {
	type args struct {
		ctx  context.Context
		data *model.UserRegistrationData
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
				ctx: context.Background(),
				data: &model.UserRegistrationData{
					UserName: "test",
					Name:     "test",
					Password: VALID_PASSWORD,
					Email:    "test@example.com",
				},
			},
			&model.User{
				Id:        VALID_USER_ID,
				UserName:  "test",
				Name:      "test",
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
				ctx: context.Background(),
				data: &model.UserRegistrationData{
					UserName: DUMMY_1_USERNAME, // ダミーユーザーのユーザー名と重複させる
					Name:     "test",
					Password: VALID_PASSWORD,
					Email:    "test@example.com",
				},
			},
			nil,
			ErrUserNameAlreadyExists,
		},
		{
			// メールアドレスが重複している場合
			"errEmailAlreadyExists",
			args{
				ctx: context.Background(),
				data: &model.UserRegistrationData{
					UserName: "test",
					Name:     "test",
					Password: VALID_PASSWORD,
					Email:    DUMMY_1_EMAIL, // ダミーユーザーのメールアドレスと重複させる
				},
			},
			nil,
			ErrEmailAlreadyExists,
		},
		// パスワードが不正な場合
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := s.as.RegisterUser(tt.args.ctx, tt.args.data)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("UserService.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr == nil {
				assert.Equal(s.T(), tt.wantUser, got)
			}
		})
	}
}
