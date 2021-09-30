package server

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	client     *s3.Client
	bucketName string
}

func NewStorage(ctx context.Context, bucketName string) (*Storage, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)

	s := Storage{
		client:     client,
		bucketName: bucketName,
	}

	return &s, nil
}

func (s *Storage) List(ctx context.Context) ([]Receipt, error) {
	receipts := []Receipt{}

	res, err := s.client.ListObjects(ctx, &s3.ListObjectsInput{
		Bucket: &s.bucketName,
	})
	if err != nil {
		return nil, err
	}
	for _, o := range res.Contents {
		r := Receipt{
			Name: *o.Key,
		}

		receipts = append(receipts, r)
	}

	return receipts, nil
}
