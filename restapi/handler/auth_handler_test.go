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
)

func setupForAuthHandlerTest(t *testing.T, moqService *AuthServiceMock) {
	t.Helper()

	fc := &clock.FixedClocker{}
	moqService.RegisterUserFunc = func(ctx context.Context, data *model.UserRegistrationData) (*model.User, error) {
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
	}
}

func TestRegisterUser(t *testing.T) {
	t.Parallel()

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
			http.StatusBadRequest,
			"testdata/auth/register/user_name_exists_response.json.golden",
		},
		// メールアドレスが既に存在する場合
		{
			"alreadyExistsEmail",
			"testdata/auth/register/email_exists_request.json.golden",
			http.StatusBadRequest,
			"testdata/auth/register/email_exists_response.json.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moqService := &AuthServiceMock{}
			setupForAuthHandlerTest(t, moqService)
			ah := &AuthHandler{
				Service: moqService,
			}

			url := testutil.RunTestServer(t, "POST", "/user", ah.RegisterUser)
			reqBody := testutil.LoadFile(t, tt.reqFile)
			resp := testutil.SendRequest(t, "POST", url, reqBody)
			// 期待するレスポンスボディのファイルをロードする
			wantResp := testutil.LoadFile(t, tt.wantRespFile)
			testutil.AssertResponse(t, resp, tt.wantStatus, wantResp)
		})
	}
}
