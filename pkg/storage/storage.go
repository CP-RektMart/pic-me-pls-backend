package storage

import (
	"context"
	"io"

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
	if _, err := c.Client.UploadFile(c.Bucket, path, data, storage_go.FileOptions{
		ContentType: &contentType,
		Upsert:      &overwrite,
	}); err != nil {
		return errors.Wrap(err, "failed to upload file")
	}

	return nil
}

func (c *Client) MoveFile(ctx context.Context, source string, destination string) error {
	if _, err := c.Client.MoveFile(c.Bucket, source, destination); err != nil {
		return errors.Wrap(err, "failed to move file")
	}

	return nil
}

func (c *Client) DeleteFile(ctx context.Context, path string) error {
	if err := c.DeleteFiles(ctx, []string{path}); err != nil {
		return errors.Wrap(err, "failed to delete file")
	}

	return nil
}

func (c *Client) DeleteFiles(ctx context.Context, path []string) error {
	if _, err := c.Client.RemoveFile(c.Bucket, path); err != nil {
		return errors.Wrap(err, "failed to delete file")
	}

	return nil
}
