package storage

import (
	"context"
	"io"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/cockroachdb/errors"
	storage_go "github.com/supabase-community/storage-go"
	"github.com/supabase-community/supabase-go"
)

type Config struct {
	Url    string `env:"URL"`
	Secret string `env:"SECRET"`
	Bucket string `env:"BUCKET"`
}

type Client struct {
	Client *storage_go.Client
	Bucket string
}

func New(ctx context.Context, config Config) (*Client, error) {
	client, err := supabase.NewClient(config.Url, config.Secret, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new Supabase client")
	}

	return &Client{
		Client: client.Storage,
		Bucket: config.Bucket,
	}, nil
}

func (c *Client) UploadFile(ctx context.Context, path string, contentType string, data io.Reader, overwrite bool) error {
	response, err := c.Client.UploadFile(c.Bucket, path, data, storage_go.FileOptions{
		ContentType: &contentType,
		Upsert:      &overwrite,
	})
	if err != nil {
		logger.ErrorContext(ctx, "failed to upload file", err)
		return errors.Wrap(err, "failed to upload file")
	}

	if response.Error != "" {
		logger.ErrorContext(ctx, "failed to upload file", response.Error)
		return errors.New(response.Error)
	}

	return nil
}
