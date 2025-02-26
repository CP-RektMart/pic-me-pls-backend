package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"

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
	config Config
}

func New(ctx context.Context, config Config) (*Client, error) {
	client, err := supabase.NewClient(config.Url, config.Secret, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new Supabase client")
	}

	return &Client{
		Client: client.Storage,
		config: config,
	}, nil
}

func (c *Client) UploadFile(ctx context.Context, path string, contentType string, data io.Reader, overwrite bool) (string, error) {
	if _, err := c.Client.UploadFile(c.config.Bucket, path, data, storage_go.FileOptions{
		ContentType: &contentType,
		Upsert:      &overwrite,
	}); err != nil {
		return "", errors.Wrap(err, "failed to upload file")
	}

	fileURL, err := c.CleanURL(c.Client.GetPublicUrl(c.config.Bucket, path).SignedURL)
	if err != nil {
		return "", errors.Wrap(err, "failed to clean url")
	}

	return fileURL, nil
}

func (c *Client) MoveFile(ctx context.Context, source string, destination string) error {
	if _, err := c.Client.MoveFile(c.config.Bucket, source, destination); err != nil {
		return errors.Wrap(err, "failed to move file")
	}

	return nil
}

// You can pass any path like /folder/file.jpg or https://your-supabase-url.supabase.co/storage/v1/object/public/your-bucket/folder/file.jpg
func (c *Client) DeleteFile(ctx context.Context, path string) error {
	if err := c.DeleteFiles(ctx, []string{path}); err != nil {
		return errors.Wrap(err, "failed to delete file")
	}
	return nil
}

// You can pass any path like /folder/file.jpg or https://your-supabase-url.supabase.co/storage/v1/object/public/your-bucket/folder/file.jpg
func (c *Client) DeleteFiles(ctx context.Context, path []string) error {
	relativePaths := make([]string, 0, len(path))
	for _, p := range path {
		relativePath, err := c.RelativePath(p)
		if err != nil {
			return errors.Wrap(err, "failed to get relative path")
		}
		relativePaths = append(relativePaths, relativePath)
	}

	// There's some bug. If I start the server, and trigger remove file, it will work.
	// But if you trigger upload file before remove file, it will not work.
	client, err := supabase.NewClient(c.config.Url, c.config.Secret, nil)
	if err != nil {
		return errors.Wrap(err, "failed to new Supabase client")
	}

	if _, err := client.Storage.RemoveFile(c.config.Bucket, relativePaths); err != nil {
		return errors.Wrap(err, "failed to delete file")
	}

	return nil
}

func (c *Client) CleanURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse url")
	}

	// Clean up the path using path.Clean
	parsedURL.Path = path.Clean(parsedURL.Path)

	return parsedURL.String(), nil
}

func (c *Client) RelativePath(path string) (string, error) {
	bucketURL, err := url.Parse(c.config.Url)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse bucket url")
	}

	cleanedUrl, err := c.CleanURL(path)
	if err != nil {
		return "", errors.Wrap(err, "failed to clean url")
	}

	return strings.Replace(cleanedUrl, fmt.Sprintf("https://%s/storage/v1/object/public/%s/", bucketURL.Host, c.config.Bucket), "", 1), nil
}
