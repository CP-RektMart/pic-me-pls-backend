package redis

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Url string `env:"URL"`
}

func New(ctx context.Context, config Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(config.Url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse redis url")
	}
	rdb := redis.NewClient(opt)

	return rdb, nil
}
