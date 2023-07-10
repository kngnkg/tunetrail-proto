package s3downloader

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/stretchr/testify/assert"
)

func CreateS3DownloaderForTest(dm *DownloaderMock) *S3Downloader {
	return &S3Downloader{
		Sess:       nil,
		Bucket:     "mock-bucket",
		downloader: dm,
	}
}

func TestS3Downloader_Download(t *testing.T) {
	type args struct {
		ctx  context.Context
		item string
	}

	want := func(item string) *os.File {
		filename := filepath.Base(item)
		file, err := os.Create("/tmp/" + filename)
		if err != nil {
			t.Fatal(err)
		}
		return file
	}

	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				item: "/path/to/mock-item.txt",
			},
			want:    want("/path/to/mock-item.txt"),
			wantErr: nil,
		},
		{
			name: "failed to download",
			args: args{
				ctx:  context.Background(),
				item: "wrong-item",
			},
			want:    nil,
			wantErr: ErrorUnableToDownloadItem,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm := &DownloaderMock{}
			dm.DownloadWithContextFunc = func(
				ctx aws.Context,
				w io.WriterAt,
				input *s3.GetObjectInput,
				options ...func(*s3manager.Downloader),
			) (n int64, err error) {
				if tt.args.item == "wrong-item" {
					return 0, errors.New("error from mock")
				}
				return 0, nil
			}

			s3dl := CreateS3DownloaderForTest(dm)

			got, err := s3dl.Download(tt.args.ctx, tt.args.item)
			// テスト終了後にファイルを削除
			t.Cleanup(func() {
				fileName := filepath.Base(tt.args.item)
				if err = os.Remove("/tmp/" + fileName); err != nil {
					t.Fatal(err)
				}
			})

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("want err: %v but got: %v", tt.wantErr, err)
			}

			// 異常系の場合はファイルが作成されない
			if tt.wantErr != nil {
				assert.Nil(t, got)
				return
			}

			// 正常系の場合はファイルが作成される
			assert.Equal(t, tt.want.Name(), got.Name())

			// // ファイルが/tmp/に作成される
			// assert.Equal(t, got.Name(), tt.want.Name())
		})
	}
}
