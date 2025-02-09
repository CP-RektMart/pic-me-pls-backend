package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	jwtConfig jwt.Config
	cache     *redis.Client
}

func NewService(jwtConfig jwt.Config, cache *redis.Client) *Service {
	return &Service{
		jwtConfig: jwtConfig,
		cache:     cache,
	}
}

func (s *Service) GenerateTokenPair(ctx context.Context, user *model.User) (*dto.TokenResponse, error) {
	cachedToken, accessToken, refreshToken, exp, err := jwt.GenerateTokenPair(
		*user,
		s.jwtConfig.AccessTokenSecret,
		s.jwtConfig.RefreshTokenSecret,
		s.jwtConfig.AccessTokenExpire,
		s.jwtConfig.RefreshTokenExpire,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token pair")
	}

	if err := s.storeCacheTokens(
		ctx,
		cachedToken,
		user.ID,
		int(s.jwtConfig.RefreshTokenExpire),
	); err != nil {
		return nil, errors.Wrap(err, "failed to store cache tokens")
	}

	return &dto.TokenResponse{
		AcessToken:   accessToken,
		RefreshToken: refreshToken,
		Exp:          exp,
	}, nil
}

func (s *Service) GetCachedTokens(ctx context.Context, userID uint) (*model.CachedTokens, error) {
	var cachedToken model.CachedTokens
	val, err := s.cache.Get(ctx, s.newTokenKey(userID)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cached token")
	}

	if err := json.Unmarshal([]byte(val), &cachedToken); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal cached token")
	}

	return &cachedToken, nil
}

func (s *Service) RemoveToken(ctx context.Context, cache *redis.Client, userID uint) error {
	return cache.Del(ctx, s.newTokenKey(userID)).Err()
}

func (s *Service) newTokenKey(userID uint) string {
	return fmt.Sprintf("auth:token:%d", userID)
}

func (s *Service) storeCacheTokens(
	ctx context.Context,
	tokens *model.CachedTokens,
	userID uint,
	ttl int,
) error {
	tokensJSON, err := json.Marshal(tokens)
	if err != nil {
		return errors.Wrap(err, "failed to marshal tokens")
	}

	return s.cache.Set(
		ctx,
		s.newTokenKey(userID),
		tokensJSON,
		time.Second*time.Duration(ttl),
	).Err()
}
