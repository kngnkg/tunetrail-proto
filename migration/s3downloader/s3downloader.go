package s3downloader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . Downloader
type Downloader interface {
	DownloadWithContext(ctx aws.Context, w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error)
}

type S3Downloader struct {
	Sess       *session.Session // AWSセッション
	Bucket     string           // バケット名
	downloader Downloader       // S3からファイルをダウンロードするためのインターフェース
}

func New(sess *session.Session, bucket string) *S3Downloader {
	downloader := s3manager.NewDownloader(sess)

	return &S3Downloader{
		Sess:       sess,
		Bucket:     bucket,
		downloader: downloader,
	}
}

var (
	ErrUnableToCreateFile     = errors.New("s3download: unable to create file")
	ErrorUnableToDownloadItem = errors.New("s3download: unable to download item")
)

// S3からファイルを/tmp/にダウンロードする
// item: バケット内のオブジェクトのキー
func (d *S3Downloader) Download(ctx context.Context, item string) (*os.File, error) {
	filename := filepath.Base(item) // ファイル名のみを抽出

	// Lambdaの/tmpディレクトリにファイルを作成
	file, err := os.Create("/tmp/" + filename)
	if err != nil {
		return nil, fmt.Errorf("%w: item=%v: %w", ErrUnableToCreateFile, item, err)
	}

	_, err = d.downloader.DownloadWithContext(ctx, file,
		&s3.GetObjectInput{
			Bucket: aws.String(d.Bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		return nil, fmt.Errorf("%w: item=%v: %w", ErrorUnableToDownloadItem, item, err)
	}

	return file, nil
}
