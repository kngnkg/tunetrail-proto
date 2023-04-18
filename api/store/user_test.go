package store

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kwtryo/tunetrail/api/clock"
	"github.com/kwtryo/tunetrail/api/model"
	"github.com/kwtryo/tunetrail/api/testutil"
	"github.com/kwtryo/tunetrail/api/testutil/fixture"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserStoreTestSuite struct {
	suite.Suite
	r     *Repository         // テスト対象のリポジトリ
	ctx   context.Context     // テスト用のコンテキスト
	db    *sqlx.DB            // テスト用のDB
	fc    *clock.FixedClocker // テスト用の時刻を固定する
	dummy *model.User         // テスト用のダミーユーザー
}

func TestUserStoreTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, &UserStoreTestSuite{
		r: &Repository{
			Clocker: &clock.FixedClocker{},
		},
		fc: &clock.FixedClocker{},
	})
}

func (s *UserStoreTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.db = testutil.OpenDbForTest(s.T(), s.ctx)

	// テスト用のダミーユーザー
	s.dummy = fixture.User(&model.User{
		UserName: "dummy",
		Email:    "dummy@example.com",
		// タイムスタンプを固定する
		CreatedAt: s.fc.Now(),
		UpdatedAt: s.fc.Now(),
	})
}

// トランザクションを開始し、テスト用のテーブルを準備する
func initUserStoreTest(
	t *testing.T, ctx context.Context, db *sqlx.DB, r *Repository, dummy *model.User,
) *sqlx.Tx {
	// トランザクションを開始する
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = tx.Rollback() })

	// テーブルを初期化する
	testutil.DeleteUserAll(t, ctx, tx)
	// テスト用のダミーユーザーを登録する
	if err := r.RegisterUser(ctx, tx, dummy); err != nil {
		t.Fatal(err)
	}
	return tx
}

func (s *UserStoreTestSuite) TestRegisterUser() {
	tests := []struct {
		name    string
		user    *model.User
		wantErr error
	}{
		{
			"ok",
			fixture.User(&model.User{
				// タイムスタンプを固定する
				CreatedAt: s.fc.Now(),
				UpdatedAt: s.fc.Now(),
			}),
			nil,
		},
		{
			// ユーザー名が既に存在する場合はエラーになる
			"errUserNameAlreadyExists",
			fixture.User(&model.User{
				// ダミーユーザーと同じユーザー名を設定する
				UserName:  "dummy",
				CreatedAt: s.fc.Now(),
				UpdatedAt: s.fc.Now(),
			}),
			ErrUserNameAlreadyExists,
		},
		{
			// メールアドレスが既に存在する場合はエラーになる
			"errEmailAlreadyExists",
			fixture.User(&model.User{
				// ダミーユーザーと同じメールアドレスを設定する
				Email:     "dummy@example.com",
				CreatedAt: s.fc.Now(),
				UpdatedAt: s.fc.Now(),
			}),
			ErrEmailAlreadyExists,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)
			err := s.r.RegisterUser(s.ctx, tx, tt.user)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr == nil {
				// ユーザーが登録されているか確認する
				got, err := s.r.GetUserByUserName(s.ctx, tx, tt.user.Name)
				if err != nil {
					s.T().Fatal(err)
				}
				assert.Equal(s.T(), tt.user, got)
			}
		})
	}
}

func (s *UserStoreTestSuite) TestUserExistsByUserName() {
	tests := []struct {
		name     string
		userName string
		want     bool
		wantErr  error
	}{
		{
			// ユーザー名が存在する場合はtrueを返す
			"okExists",
			// ダミーユーザーのユーザー名を設定する
			"dummy",
			true,
			nil,
		},
		{
			// ユーザー名が存在しない場合はfalseを返す
			"okNotExists",
			// 存在しないユーザー名を設定する
			"notExists",
			false,
			nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)
			got, err := s.r.UserExistsByUserName(s.ctx, tx, tt.userName)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.UserExistsByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(s.T(), tt.want, got)
		})
	}
}

func (s *UserStoreTestSuite) TestUserExistsByEmail() {
	tests := []struct {
		name    string
		email   string
		want    bool
		wantErr error
	}{
		{
			// メールアドレスが存在する場合はtrueを返す
			"okExists",
			// ダミーユーザーのメールアドレスを設定する
			"dummy@example.com",
			true,
			nil,
		},
		{
			// メールアドレスが存在しない場合はfalseを返す
			"okNotExists",
			// 存在しないメールアドレスを設定する
			"notExists@example.com",
			false,
			nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)
			got, err := s.r.UserExistsByEmail(s.ctx, tx, tt.email)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.UserExistsByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(s.T(), tt.want, got)
		})
	}
}

func (s *UserStoreTestSuite) TestGetUserByUserName() {
	tests := []struct {
		name     string
		userName string
		want     *model.User
		wantErr  error
	}{
		{
			// ユーザー名が存在する場合は該当するユーザーを返す
			"ok",
			// ダミーユーザーのユーザー名を設定する
			"dummy",
			s.dummy,
			nil,
		},
		{
			// ユーザー名が存在しない場合はnilを返す
			"errNotExists",
			// 存在しないユーザー名を設定する
			"notExists",
			nil,
			ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)
			got, err := s.r.GetUserByUserName(s.ctx, tx, tt.userName)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.GetUserByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(s.T(), tt.want, got)
		})
	}
}

func (s *UserStoreTestSuite) TestGetUserByEmail() {
	tests := []struct {
		name    string
		email   string
		want    *model.User
		wantErr error
	}{
		{
			// メールアドレスが存在する場合は該当するユーザーを返す
			"ok",
			// ダミーユーザーのメールアドレスを設定する
			"dummy@example.com",
			s.dummy,
			nil,
		},
		{
			// メールアドレスが存在しない場合はnilを返す
			"errNotExists",
			// 存在しないメールアドレスを設定する
			"notExists@example.com",
			nil,
			ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)
			got, err := s.r.GetUserByEmail(s.ctx, tx, tt.email)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(s.T(), tt.want, got)
		})
	}
}

func (s *UserStoreTestSuite) TestDeleteUserByUserName() {
	tests := []struct {
		name        string
		userName    string
		dummyExists bool // ダミーユーザーが存在するかどうか
		wantErr     error
	}{
		{
			"ok",
			// ダミーユーザーのユーザー名を設定する
			"dummy",
			false,
			nil,
		},
		{
			// ユーザー名が存在しない場合はエラーを返す
			"errNotExists",
			// 存在しないユーザー名を設定する
			"notExists",
			true,
			ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)
			err := s.r.DeleteUserByUserName(s.ctx, tx, tt.userName)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.UserExistsByUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
			// ダミーユーザーが削除されているか確認する
			exists, err := s.r.UserExistsByUserName(s.ctx, tx, "dummy")
			if err != nil {
				s.T().Fatal(err)
			}
			assert.Equal(s.T(), tt.dummyExists, exists)
		})
	}
}
