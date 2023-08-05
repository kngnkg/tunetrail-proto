package store

import (
	"context"
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/testutil"
	"github.com/kngnkg/tunetrail/restapi/testutil/fixture"
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
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)
			err := s.r.RegisterUser(s.ctx, tx, tt.user)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr != nil {
				// 異常系の場合はここでテストを終了する
				return
			}

			// ユーザーが登録されているか確認する
			got, err := s.r.GetUserByUserName(s.ctx, tx, tt.user.Name)
			if err != nil {
				s.T().Fatal(err)
			}
			testutil.AssertUser(s.T(), tt.user, got)
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
			testutil.AssertUser(s.T(), tt.want, got)
		})
	}
}

func (s *UserStoreTestSuite) TestUpdateUser() {
	targetUser := fixture.User(&model.User{
		UserName:  "target",
		CreatedAt: s.fc.Now(),
		UpdatedAt: s.fc.Now(),
	})

	tests := []struct {
		name     string
		user     *model.User // 更新するユーザー情報
		wantUser *model.User // 更新後のユーザー情報
		wantErr  error
	}{
		{
			// ユーザー情報を更新できる
			"ok",
			&model.User{
				Id:       targetUser.Id,
				UserName: "updated",
				Name:     "updated",
				IconUrl:  "https://example.com/updated.png",
				Bio:      "updated",
			},
			&model.User{
				Id:        targetUser.Id,
				UserName:  "updated",
				Name:      "updated",
				IconUrl:   "https://example.com/updated.png",
				Bio:       "updated",
				CreatedAt: targetUser.CreatedAt,
				UpdatedAt: targetUser.UpdatedAt,
			},
			nil,
		},
		{
			// 存在しないidの場合はエラーを返す
			"errIdNotExists",
			fixture.User(&model.User{
				Id: fixture.NewUserId(),
			}),
			nil,
			ErrUserNotFound,
		},
		{
			// 存在するユーザー名の場合はエラーを返す
			"errUserNameExists",
			fixture.User(&model.User{
				UserName: "dummy",
			}),
			nil,
			ErrUserNameAlreadyExists,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tx := initUserStoreTest(s.T(), s.ctx, s.db, s.r, s.dummy)

			// テスト対象のユーザーを作成する
			if err := s.r.RegisterUser(s.ctx, tx, targetUser); err != nil {
				s.T().Fatal(err)
			}

			// "errIdNotExists"の場合は更新するユーザーのidを設定しない
			if tt.name != "errIdNotExists" {
				// 更新するユーザーのidを設定する
				tt.user.Id = targetUser.Id
			}
			// テスト対象のユーザーを更新する
			err := s.r.UpdateUser(s.ctx, tx, tt.user)
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantUser == nil {
				// 更新後のユーザー情報がnilの場合はここでテストを終了する
				return
			}

			// 更新後のユーザー情報を取得する
			got, err := s.r.GetUserByUserName(s.ctx, tx, tt.user.UserName)
			if err != nil {
				s.T().Fatal(err)
			}
			// 更新後のユーザー情報のIDを設定する
			tt.wantUser.Id = got.Id
			assert.Equal(s.T(), tt.wantUser, got)
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
