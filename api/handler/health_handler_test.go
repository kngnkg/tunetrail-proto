package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/kwtryo/tunetrail/api/model"
	"github.com/kwtryo/tunetrail/api/testutil"
)

func TestHealthHandler_HealthCheck(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		wantStatus   int    // ステータスコード
		wantRespFile string // レスポンス

	}{
		{
			"ok",
			http.StatusOK,
			"testdata/health/ok_response.json.golden",
		},
		{
			"internalServerError",
			http.StatusInternalServerError,
			"testdata/health/server_err_response.json.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moqService := &HealthServiceMock{}
			moqService.HealthCheckFunc =
				func(ctx context.Context) (*model.Health, error) {
					if tt.name == "internalServerError" {
						h := &model.Health{
							Health:   model.StatusOrange,
							Database: model.StatusRed,
						}
						return h, errors.New("error from mock")
					}
					h := &model.Health{
						Health:   model.StatusGreen,
						Database: model.StatusGreen,
					}
					return h, nil
				}

			hh := &HealthHandler{
				Service: moqService,
			}

			url := testutil.RunTestServer(t, "GET", hh.HealthCheck)
			resp := testutil.SendGetRequest(t, url)
			// 期待するレスポンスボディのファイルをロードする
			wantResp := testutil.LoadFile(t, tt.wantRespFile)
			testutil.AssertResponse(t, resp, tt.wantStatus, wantResp)
		})
	}
}
