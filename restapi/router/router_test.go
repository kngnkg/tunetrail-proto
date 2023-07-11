package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kngnkg/tunetrail/restapi/config"
	"github.com/kngnkg/tunetrail/restapi/testutil"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	t.Parallel()

	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name       string
		args       args
		endpoint   string
		wantStatus int
		wantErr    bool
	}{
		{
			"ok",
			args{testutil.CreateConfigForTest(t)},
			"/health",
			http.StatusOK,
			false,
		},
		{
			// 存在しないエンドポイントにアクセスした場合
			"nonexistentEndpoints",
			args{testutil.CreateConfigForTest(t)},
			"/nonexistent",
			http.StatusNotFound,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotr, gotf, err := SetupRouter(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetupRouter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testServer := httptest.NewServer(gotr) // サーバを立てる
			t.Cleanup(func() {
				gotf()
				testServer.Close()
			})

			url := fmt.Sprintf(testServer.URL + tt.endpoint)
			t.Logf("try request to %q", url)
			// サーバーにリクエストを送信
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			defer resp.Body.Close()

			// ステータスコードの確認
			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}
