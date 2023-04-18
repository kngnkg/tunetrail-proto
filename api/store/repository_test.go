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

type RepositoryTestSuite struct {
	suite.Suite
	r     *Repository         // テスト対象のリポジトリ
	ctx   context.Context     // テスト用のコンテキスト
	db    *sqlx.DB            // テスト用のDB
	fc    *clock.FixedClocker // テスト用の時刻を固定する
	dummy *model.User         // テスト用のダミーユーザー
}

func TestRepositoryTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, &UserStoreTestSuite{
		r: &Repository{
			Clocker: &clock.FixedClocker{},
		},
		fc: &clock.FixedClocker{},
	})
}

func (s *RepositoryTestSuite) SetupTest() {
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

func (s *RepositoryTestSuite) TestWithTransaction() {
	tests := []struct {
		name       string
		user       *model.User // テスト用のユーザー
		wantExists bool        // ユーザーが登録されているかどうか
		wantErr    error
	}{
		{
			"ok",
			fixture.User(&model.User{
				UserName: "tranTest",
				Email:    "tranTest@email.com",
				// タイムスタンプを固定する
				CreatedAt: s.fc.Now(),
				UpdatedAt: s.fc.Now(),
			}),
			true,
			nil,
		},
		{
			"rollback",
			fixture.User(&model.User{
				UserName: "tranTest",
				Email:    "tranTest@email.com",
				// タイムスタンプを固定する
				CreatedAt: s.fc.Now(),
				UpdatedAt: s.fc.Now(),
			}),
			false,
			nil,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			// トランザクションを開始する
			err := s.r.WithTransaction(s.ctx, s.db, func(tx *sqlx.Tx) error {
				// テーブルを初期化する
				testutil.DeleteUserAll(s.T(), s.ctx, tx)
				// ユーザーを登録する
				err := s.r.RegisterUser(s.ctx, tx, tt.user)
				if err != nil {
					s.T().Fatal(err)
				}

				if tt.name != "ok" {
					return errors.New("error from test")
				}
				return nil
			})
			if !errors.Is(err, tt.wantErr) {
				s.T().Errorf("Repository.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			// ユーザーが登録されているか確認する
			exists, err := s.r.UserExistsByUserName(s.ctx, s.db, tt.user.UserName)
			if err != nil {
				s.T().Fatal(err)
			}
			assert.Equal(s.T(), tt.wantExists, exists)

			if exists {
				// ユーザーを削除する
				err := s.r.DeleteUserByUserName(s.ctx, s.db, tt.user.UserName)
				if err != nil {
					s.T().Fatal(err)
				}
			}
		})
	}
}
