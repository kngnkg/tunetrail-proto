package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/kngnkg/tunetrail/restapi/model"
	"github.com/kngnkg/tunetrail/restapi/store"
)

func TestHealthService_HealthCheck(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Health
		wantErr error
	}{
		{
			"ok",
			args{context.Background()},
			&model.Health{
				Health:   model.StatusGreen,
				Database: model.StatusGreen,
			},
			nil,
		},
		{
			// DBとの疎通が取れない場合
			"errCannotCommunicateWithDB",
			args{context.Background()},
			&model.Health{
				Health:   model.StatusOrange,
				Database: model.StatusRed,
			},
			store.ErrCannotCommunicateWithDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moqDB := &QueryerMock{}
			moqRepo := &HealthRepositoryMock{}
			moqRepo.PingFunc = func(ctx context.Context, db store.Queryer) error {
				if tt.name == "errCannotCommunicateWithDB" {
					return store.ErrCannotCommunicateWithDB
				}
				return nil
			}

			hs := &HealthService{
				DB:   moqDB,
				Repo: moqRepo,
			}
			got, err := hs.HealthCheck(tt.args.ctx)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("want err: %v but got: %v", tt.wantErr, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HealthService.HealthCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
