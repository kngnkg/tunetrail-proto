package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/kngnkg/tunetrail/restapi/clock"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/service"
	"github.com/kngnkg/tunetrail/restapi/testutil"
	"github.com/kngnkg/tunetrail/restapi/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	ah *AuthHandler // テスト対象のハンドラ
}

func TestAuthHandlerTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(AuthHandlerTestSuite))
}

func (s *AuthHandlerTestSuite) SetupTest() {
	testutil.SetGinTestMode(s.T())

	// バリデーションの初期化
	validate.InitValidation()

	asm := setupAuthServiceMock(s.T())
	s.ah = &AuthHandler{
		Service:       asm,
		AllowedDomain: "example.com",
	}
}

func (s *AuthHandlerTestSuite) TestRegisterUser() {
	tests := []struct {
		name         string
		reqFile      string // リクエストのファイルパス
		wantStatus   int    // ステータスコード
		wantRespFile string // レスポンスのファイルパス
	}{
		// 正常系
		{
			"ok",
			"testdata/auth/register/ok_request.json.golden",
			http.StatusOK,
			"testdata/auth/register/ok_response.json.golden",
		},
		// フィールドの値が不正な場合
		{
			"emptyUserName",
			"testdata/auth/register/empty_user_name_request.json.golden",
			http.StatusBadRequest,
			"testdata/auth/register/empty_user_name_response.json.golden",
		},
		// ユーザー名が既に存在する場合
		{
			"alreadyExistsUserName",
			"testdata/auth/register/user_name_exists_request.json.golden",
			http.StatusConflict,
			"testdata/auth/register/user_name_exists_response.json.golden",
		},
		// メールアドレスが既に存在する場合
		{
			"alreadyExistsEmail",
			"testdata/auth/register/email_exists_request.json.golden",
			http.StatusConflict,
			"testdata/auth/register/email_exists_response.json.golden",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			url := testutil.RunTestServer(s.T(), "POST", "/user/register", s.ah.RegisterUser)
			reqBody := testutil.LoadFile(s.T(), tt.reqFile)
			resp := testutil.SendRequest(s.T(), "POST", url, reqBody)

			assert.Equal(s.T(), tt.wantStatus, resp.StatusCode)

			if tt.wantRespFile == "" {
				return
			}
			wantResp := testutil.LoadFile(s.T(), tt.wantRespFile)
			testutil.AssertResponseBody(s.T(), resp, wantResp)
		})
	}
}

func (s *AuthHandlerTestSuite) TestConfirmEmail() {
	tests := []struct {
		name         string
		reqFile      string // リクエストのファイルパス
		wantStatus   int    // ステータスコード
		wantRespFile string // レスポンスのファイルパス
	}{
		{
			"success",
			"testdata/auth/confirm_email/ok_request.json.golden",
			http.StatusNoContent,
			"",
		},
		// リクエストの値が不正な場合
		{
			"invalid parameter",
			"testdata/auth/confirm_email/invalid_parameter_request.json.golden",
			http.StatusBadRequest,
			"testdata/auth/confirm_email/invalid_parameter_response.json.golden",
		},
		// ユーザー名が存在しない場合
		{
			"UserName not found",
			"testdata/auth/confirm_email/username_not_found_request.json.golden",
			http.StatusNotFound,
			"testdata/auth/confirm_email/username_not_found_response.json.golden",
		},
		// コードが不正な場合
		{
			"invalid code",
			"testdata/auth/confirm_email/invalid_code_request.json.golden",
			http.StatusBadRequest,
			"testdata/auth/confirm_email/invalid_code_response.json.golden",
		},
		// コードが期限切れの場合
		{
			"expired code",
			"testdata/auth/confirm_email/expired_code_request.json.golden",
			http.StatusBadRequest,
			"testdata/auth/confirm_email/expired_code_response.json.golden",
		},
		// 既に確認済みの場合
		{
			"already confirmed",
			"testdata/auth/confirm_email/already_confirmed_request.json.golden",
			http.StatusConflict,
			"testdata/auth/confirm_email/already_confirmed_response.json.golden",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			url := testutil.RunTestServer(s.T(), "PUT", "/user/confirm", s.ah.ConfirmEmail)
			reqBody := testutil.LoadFile(s.T(), tt.reqFile)
			resp := testutil.SendRequest(s.T(), "PUT", url, reqBody)

			assert.Equal(s.T(), tt.wantStatus, resp.StatusCode)

			if tt.wantRespFile == "" {
				return
			}
			wantResp := testutil.LoadFile(s.T(), tt.wantRespFile)
			testutil.AssertResponseBody(s.T(), resp, wantResp)
		})
	}
}

func (s *AuthHandlerTestSuite) TestSignIn() {
	tests := []struct {
		name         string
		reqFile      string // リクエストのファイルパス
		wantStatus   int    // ステータスコード
		wantRespFile string // レスポンスのファイルパス
	}{
		{
			"success with username",
			"testdata/auth/signin/ok_username_request.json.golden",
			http.StatusOK,
			"",
		},
		{
			"success with email",
			"testdata/auth/signin/ok_email_request.json.golden",
			http.StatusOK,
			"",
		},
		// フィールドの値が不正な場合
		{
			"invalid parameter",
			"testdata/auth/signin/invalid_parameter_request.json.golden",
			http.StatusBadRequest,
			"testdata/auth/signin/invalid_parameter_response.json.golden",
		},
		// ユーザー名が存在しない場合
		{
			"UserName not found",
			"testdata/auth/signin/username_not_found_request.json.golden",
			http.StatusNotFound,
			"testdata/auth/signin/username_not_found_response.json.golden",
		},
		// メールアドレスが存在しない場合
		{
			"Email not found",
			"testdata/auth/signin/email_not_found_request.json.golden",
			http.StatusNotFound,
			"testdata/auth/signin/email_not_found_response.json.golden",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			url := testutil.RunTestServer(s.T(), "POST", "/user/signin", s.ah.SignIn)
			reqBody := testutil.LoadFile(s.T(), tt.reqFile)
			resp := testutil.SendRequest(s.T(), "POST", url, reqBody)

			assert.Equal(s.T(), tt.wantStatus, resp.StatusCode)

			// 正常系の場合
			if tt.wantRespFile == "" {
				// レスポンスのヘッダーにCookieが含まれていることを確認
				cookie := resp.Header["Set-Cookie"]
				assert.NotEmpty(s.T(), cookie)
				return
			}

			wantResp := testutil.LoadFile(s.T(), tt.wantRespFile)
			testutil.AssertResponseBody(s.T(), resp, wantResp)
		})
	}
}

func setupAuthServiceMock(t *testing.T) *AuthServiceMock {
	t.Helper()
	fc := &clock.FixedClocker{}

	mock := &AuthServiceMock{
		RegisterUserFunc: func(ctx context.Context, data *model.UserRegistrationData) (*model.User, error) {
			// ユーザー名またはメールアドレスが既に存在する場合はエラーを返す
			if data.UserName == "alreadyExists" {
				t.Log("username already exists")
				return nil, service.ErrUserNameAlreadyExists
			}
			if data.Email == "alreadyExists@example.com" {
				t.Log("email already exists")
				return nil, service.ErrEmailAlreadyExists
			}
			u := &model.User{
				Id:        "1",
				UserName:  data.UserName,
				Name:      data.Name,
				CreatedAt: fc.Now(),
				UpdatedAt: fc.Now(),
			}
			return u, nil
		},
		ConfirmEmailFunc: func(ctx context.Context, userName, code string) error {
			if userName == "notFound" {
				return service.ErrUserNotFound
			}
			if code == "mismatch" {
				return service.ErrCodeMismatch
			}
			if code == "expired" {
				return service.ErrCodeExpired
			}
			if code == "confirmed" {
				return service.ErrEmailAlreadyExists
			}
			return nil
		},
		SignInFunc: func(ctx context.Context, data *model.UserSignInData) (*model.Tokens, error) {
			if data.UserName == "notFound" {
				return nil, service.ErrUserNotFound
			}
			if data.Email == "notFound@example.com" {
				return nil, service.ErrUserNotFound
			}
			tokens := &model.Tokens{
				Id:      "dummy",
				Access:  "dummy",
				Refresh: "dummy",
			}
			return tokens, nil
		},
	}

	return mock
}
