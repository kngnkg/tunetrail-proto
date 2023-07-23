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
)

func setupForUserHandlerTest(t *testing.T, moqService *UserServiceMock) {
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

	moqService.GetUserByUserNameFunc = func(ctx context.Context, userName string) (*model.User, error) {
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
	}

	moqService.UpdateUserFunc = func(ctx context.Context, u *model.UserUpdateData) error {
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
	}

	moqService.DeleteUserByUserNameFunc = func(ctx context.Context, userName string) error {
		switch userName {
		case "notFound":
			return service.ErrUserNotFound
		case "dummy":
			return nil
		default:
			return errors.New("unexpected error")
		}
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
			"testdata/user/register/ok_request.json.golden",
			http.StatusOK,
			"testdata/user/register/ok_response.json.golden",
		},
		// フィールドの値が不正な場合
		{
			"emptyUserName",
			"testdata/user/register/empty_user_name_request.json.golden",
			http.StatusBadRequest,
			"testdata/user/register/empty_user_name_response.json.golden",
		},
		// ユーザー名が既に存在する場合
		{
			"alreadyExistsUserName",
			"testdata/user/register/user_name_exists_request.json.golden",
			http.StatusBadRequest,
			"testdata/user/register/user_name_exists_response.json.golden",
		},
		// メールアドレスが既に存在する場合
		{
			"alreadyExistsEmail",
			"testdata/user/register/email_exists_request.json.golden",
			http.StatusBadRequest,
			"testdata/user/register/email_exists_response.json.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moqService := &UserServiceMock{}
			setupForUserHandlerTest(t, moqService)
			uh := &UserHandler{
				Service: moqService,
			}

			url := testutil.RunTestServer(t, "POST", "/user", uh.RegisterUser)
			reqBody := testutil.LoadFile(t, tt.reqFile)
			resp := testutil.SendRequest(t, "POST", url, reqBody)
			// 期待するレスポンスボディのファイルをロードする
			wantResp := testutil.LoadFile(t, tt.wantRespFile)
			testutil.AssertResponse(t, resp, tt.wantStatus, wantResp)
		})
	}
}

func TestGetUserByUserName(t *testing.T) {
	t.Parallel()

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
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("userName: %s", tt.userName)
			moqService := &UserServiceMock{}
			setupForUserHandlerTest(t, moqService)
			uh := &UserHandler{
				Service: moqService,
			}

			url := testutil.RunTestServer(t, "GET", "/user/:user_name", uh.GetUserByUserName)
			url = strings.Replace(url, ":user_name", tt.userName, 1)
			resp := testutil.SendGetRequest(t, url)
			// 期待するレスポンスボディのファイルをロードする
			wantResp := testutil.LoadFile(t, tt.wantRespFile)
			testutil.AssertResponse(t, resp, tt.wantStatus, wantResp)
		})
	}
}

func TestUpdateUser(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			moqService := &UserServiceMock{}
			setupForUserHandlerTest(t, moqService)
			uh := &UserHandler{
				Service: moqService,
			}

			url := testutil.RunTestServer(t, "PUT", "/user", uh.UpdateUser)
			reqBody := testutil.LoadFile(t, tt.reqFile)
			resp := testutil.SendRequest(t, "PUT", url, reqBody)
			// 期待するレスポンスボディのファイルをロードする
			wantResp := testutil.LoadFile(t, tt.wantRespFile)
			testutil.AssertResponse(t, resp, tt.wantStatus, wantResp)
		})
	}
}

func TestDeleteUserByUserName(t *testing.T) {
	t.Parallel()

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
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("userName: %s", tt.userName)
			moqService := &UserServiceMock{}
			setupForUserHandlerTest(t, moqService)
			uh := &UserHandler{
				Service: moqService,
			}

			url := testutil.RunTestServer(t, "DELETE", "/user/:user_name", uh.DeleteUserByUserName)
			url = strings.Replace(url, ":user_name", tt.userName, 1)
			resp := testutil.SendRequest(t, "DELETE", url, nil)
			wantResp := testutil.LoadFile(t, tt.wantRespFile)
			testutil.AssertResponse(t, resp, tt.wantStatus, wantResp)
		})
	}
}
