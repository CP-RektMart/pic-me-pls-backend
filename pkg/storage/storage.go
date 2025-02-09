package storage

import (
	"context"

	"github.com/cockroachdb/errors"
	supabase_go "github.com/supabase-community/supabase-go"
)

type Config struct {
	Url    string `env:"URL"`
	Secret string `env:"SECRET"`
}

type Client struct {
	Client *supabase_go.Client
}

func New(ctx context.Context) (*Client, error) {
	client, err := supabase_go.NewClient("", "", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to new Supabase client")
	}

	return &Client{
		Client: client,
	}, nil
}
