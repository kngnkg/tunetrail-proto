package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/service"
	"github.com/kngnkg/tunetrail/restapi/testutil"
	"github.com/kngnkg/tunetrail/restapi/validate"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	uh *UserHandler // テスト対象のハンドラ
}

func TestUserHandlerTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) SetupTest() {
	testutil.SetGinTestMode(s.T())

	usm := setupUserServiceMock(s.T())
	s.uh = &UserHandler{
		Service: usm,
	}
}

func (s *UserHandlerTestSuite) TestGetUserByUserName() {
	tests := []struct {
		name         string
		userName     string // ユーザー名
		wantStatus   int    // ステータスコード
		wantRespFile string // レスポンスのファイルパス
	}{
		// 正常系
		{
			"ok",
			"dummy",
			http.StatusOK,
			"testdata/user/get_by_username/ok_response.json.golden",
		},
		// ユーザー名が存在しない場合
		{
			"notFound",
			"notFound",
			http.StatusBadRequest,
			"testdata/user/get_by_username/not_found_response.json.golden",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			url := testutil.RunTestServer(s.T(), "GET", "/user/:user_name", s.uh.GetUserByUserName)
			url = strings.Replace(url, ":user_name", tt.userName, 1)
			resp := testutil.SendGetRequest(s.T(), url)
			// 期待するレスポンスボディのファイルをロードする
			wantResp := testutil.LoadFile(s.T(), tt.wantRespFile)
			testutil.AssertResponse(s.T(), resp, tt.wantStatus, wantResp)
		})
	}
}

func (s *UserHandlerTestSuite) TestUpdateUser() {
	// バリデーションの初期化
	validate.InitValidation()
	tests := []struct {
		name         string
		reqFile      string // リクエストのファイルパス
		wantStatus   int    // ステータスコード
		wantRespFile string // レスポンスのファイルパス
	}{
		// 正常系
		{
			"ok",
			"testdata/user/update/ok_request.json.golden",
			http.StatusOK,
			"testdata/user/update/ok_response.json.golden",
		},
		// フィールドの値が不正な場合
		{
			"emptyUserName",
			"testdata/user/update/empty_user_name_request.json.golden",
			http.StatusBadRequest,
			"testdata/user/update/empty_user_name_response.json.golden",
		},
		// ユーザーが存在しない場合
		{
			"notFound",
			"testdata/user/update/not_found_request.json.golden",
			http.StatusBadRequest,
			"testdata/user/update/not_found_response.json.golden",
		},
		// ユーザー名が既に存在する場合
		{
			"alreadyExistsUserName",
			"testdata/user/update/user_name_exists_request.json.golden",
			http.StatusBadRequest,
			"testdata/user/update/user_name_exists_response.json.golden",
		},
		// メールアドレスが既に存在する場合
		{
			"alreadyExistsEmail",
			"testdata/user/update/email_exists_request.json.golden",
			http.StatusBadRequest,
			"testdata/user/update/email_exists_response.json.golden",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			url := testutil.RunTestServer(s.T(), "PUT", "/user", s.uh.UpdateUser)
			reqBody := testutil.LoadFile(s.T(), tt.reqFile)
			resp := testutil.SendRequest(s.T(), "PUT", url, reqBody)
			// 期待するレスポンスボディのファイルをロードする
			wantResp := testutil.LoadFile(s.T(), tt.wantRespFile)
			testutil.AssertResponse(s.T(), resp, tt.wantStatus, wantResp)
		})
	}
}

func (s *UserHandlerTestSuite) TestDeleteUserByUserName() {
	tests := []struct {
		name         string
		userName     string // ユーザー名
		wantStatus   int    // ステータスコード
		wantRespFile string // レスポンスのファイルパス
	}{
		// 正常系
		{
			"ok",
			"dummy",
			http.StatusOK,
			"testdata/user/delete_by_username/ok_response.json.golden",
		},
		// ユーザー名が存在しない場合
		{
			"notFound",
			"notFound",
			http.StatusBadRequest,
			"testdata/user/delete_by_username/not_found_response.json.golden",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			url := testutil.RunTestServer(s.T(), "DELETE", "/user/:user_name", s.uh.DeleteUserByUserName)
			url = strings.Replace(url, ":user_name", tt.userName, 1)
			resp := testutil.SendRequest(s.T(), "DELETE", url, nil)
			wantResp := testutil.LoadFile(s.T(), tt.wantRespFile)
			testutil.AssertResponse(s.T(), resp, tt.wantStatus, wantResp)
		})
	}
}

func setupUserServiceMock(t *testing.T) *UserServiceMock {
	t.Helper()

	fc := &clock.FixedClocker{}
	mock := &UserServiceMock{
		GetUserByUserNameFunc: func(ctx context.Context, userName string) (*model.User, error) {
			u := &model.User{
				Id:        "1",
				UserName:  "dummy",
				Name:      "dummy",
				IconUrl:   "https://example.com/icon.png",
				Bio:       "dummy",
				CreatedAt: fc.Now(),
				UpdatedAt: fc.Now(),
			}
			if userName == "notFound" {
				return nil, service.ErrUserNotFound
			}
			return u, nil
		},
		UpdateUserFunc: func(ctx context.Context, u *model.UserUpdateData) error {
			if u.Id != "1" {
				return service.ErrUserNotFound
			}
			if u.UserName == "exists" {
				return service.ErrUserNameAlreadyExists
			}
			if u.Email == "exists@example.com" {
				return service.ErrEmailAlreadyExists
			}
			return nil
		},
		DeleteUserByUserNameFunc: func(ctx context.Context, userName string) error {
			switch userName {
			case "notFound":
				return service.ErrUserNotFound
			case "dummy":
				return nil
			default:
				return errors.New("unexpected error")
			}
		},
	}

	return mock
}
