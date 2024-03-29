package testutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/stretchr/testify/assert"
)

func SetGinTestMode(t *testing.T) {
	t.Helper()
	if gin.Mode() != gin.TestMode {
		gin.SetMode(gin.TestMode)
	}
}

// テスト用のサーバーを起動し、URLを返す。
// t: テスト
// reqMethod: 登録したいリクエストメソッド
// endpoint: 登録したいエンドポイント
// handler: 検証したいハンドラ
func RunTestServer(t *testing.T, reqMethod string, endpoint string, handler gin.HandlerFunc) string {
	t.Helper()
	router := gin.New()

	router.Use(func(c *gin.Context) {
		userId := model.UserID("dummy")
		c.Set("userId", userId)
		c.Next()
	})

	// ハンドラの登録
	switch reqMethod {
	case "GET":
		router.GET(endpoint, handler)
	case "POST":
		router.POST(endpoint, handler)
	case "PUT":
		router.PUT(endpoint, handler)
	case "DELETE":
		router.DELETE(endpoint, handler)
	default:
		t.Fatalf("invalid request method")
	}

	testServer := httptest.NewServer(router) // サーバを立てる
	t.Cleanup(func() { testServer.Close() })

	return fmt.Sprintf(testServer.URL + endpoint)
}

// リクエストを送信し、レスポンスを返す。
// t: テスト
// url: リクエストを送信する対象のURL
func SendGetRequest(t *testing.T, url string) *http.Response {
	t.Helper()
	t.Logf("try request to %q", url)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Cleanup(func() { _ = resp.Body.Close() })
	return resp
}

// リクエストを送信し、レスポンスを返す。
// t: テスト
// reqMethod: リクエストメソッド
// url: リクエストを送信する対象のURL
// body: リクエストボディ
func SendRequest(t *testing.T, reqMethod string, url string, body []byte) *http.Response {
	t.Helper()
	t.Logf("try request to %q", url)

	// []byteをio.Readerに変換
	var reader io.Reader
	if body == nil {
		reader = nil
	} else {
		reader = bytes.NewReader(body)
	}

	client := &http.Client{}
	req, err := http.NewRequest(reqMethod, url, reader)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Cleanup(func() { _ = resp.Body.Close() })
	return resp
}

// レスポンスボディを検証する
// t: テスト
// resp: 検証するレスポンス
// body: 期待するレスポンスボディ
func AssertResponseBody(t *testing.T, resp *http.Response, body []byte) {
	t.Helper()

	gb, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if len(gb) == 0 && len(body) == 0 {
		// レスポンスボディが無い場合は確認不要
		return
	}

	// レスポンスボディの確認
	var jw, jg any
	if err := json.Unmarshal(body, &jw); err != nil {
		t.Fatalf("cannot unmarshal want %q: %v", body, err)
	}
	if err := json.Unmarshal(gb, &jg); err != nil {
		t.Fatalf("cannot unmarshal resp %v: %v", resp, err)
	}
	assert.Equal(t, jw, jg)
}

// ファイルパスからファイルをロードする
func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read from %q: %v", path, err)
	}
	return bt
}
