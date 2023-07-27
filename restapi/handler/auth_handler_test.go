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
		Service: asm,
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
	}

	return mock
}
