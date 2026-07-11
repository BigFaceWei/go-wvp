package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Storage struct {
	client     *minio.Client
	bucket     string
	basePrefix string
}

type S3Config struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	Bucket     string
	Region     string
	UseSSL     bool
	BasePrefix string
}

func NewS3Storage(cfg *S3Config) (*S3Storage, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
		Region: cfg.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("create minio client failed: %w", err)
	}

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("check bucket failed: %w", err)
	}

	if !exists {
		if err := client.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{Region: cfg.Region}); err != nil {
			return nil, fmt.Errorf("create bucket failed: %w", err)
		}
	}

	return &S3Storage{
		client:     client,
		bucket:     cfg.Bucket,
		basePrefix: cfg.BasePrefix,
	}, nil
}

func (s *S3Storage) objectName(fileName string) string {
	if s.basePrefix != "" {
		return s.basePrefix + "/" + fileName
	}
	return fileName
}

func (s *S3Storage) Store(fileName string, reader io.Reader) (*FileInfo, error) {
	objectName := s.objectName(fileName)

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read data failed: %w", err)
	}

	_, err = s.client.PutObject(context.Background(), s.bucket, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("upload to s3 failed: %w", err)
	}

	return &FileInfo{
		FileName:    fileName,
		FilePath:    objectName,
		FileSize:    int64(len(data)),
		StorageType: 2,
	}, nil
}

func (s *S3Storage) Delete(filePath string) error {
	return s.client.RemoveObject(context.Background(), s.bucket, s.objectName(filePath), minio.RemoveObjectOptions{})
}

func (s *S3Storage) Exists(filePath string) (bool, error) {
	_, err := s.client.StatObject(context.Background(), s.bucket, s.objectName(filePath), minio.StatObjectOptions{})
	if err == nil {
		return true, nil
	}
	return false, nil
}

func (s *S3Storage) GetReader(filePath string) (io.ReadCloser, error) {
	obj, err := s.client.GetObject(context.Background(), s.bucket, s.objectName(filePath), minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *S3Storage) GetType() int {
	return 2
}
