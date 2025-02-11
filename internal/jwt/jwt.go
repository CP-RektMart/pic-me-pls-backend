package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	AccessTokenSecret  string `env:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `env:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpire  int64  `env:"ACCESS_TOKEN_EXPIRE"`
	RefreshTokenExpire int64  `env:"REFRESH_TOKEN_EXPIRE"`
	AutoLogout         int64  `env:"AUTO_LOGOUT"`
}

type JWTentity struct {
	ID   uint           `json:"id"` // User ID
	UID  uuid.UUID      `json:"uid"`
	Role model.UserRole `json:"role"`
	jwt.MapClaims
}

type JWT struct {
	config Config
	cache  *redis.Client
}

func New(config Config, cache *redis.Client) *JWT {
	return &JWT{
		config: config,
		cache:  cache,
	}
}

func (j *JWT) CreateToken(userID uint, expire int64, secret string, role model.UserRole) (token string, uid uuid.UUID, exp int64, err error) {
	exp = time.Now().Add(time.Second * time.Duration(expire)).Unix()
	uid = uuid.New()
	claims := &JWTentity{
		ID:   userID,
		UID:  uid,
		Role: role,
		MapClaims: jwt.MapClaims{
			"exp": exp,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", uuid.Nil, 0, errors.Wrap(err, "can't create token")
	}

	return token, uid, exp, nil
}

func (j *JWT) GenerateTokenPair(user model.User) (
	cahcedToken *model.CachedTokens,
	accessToken string,
	refreshToken string,
	exp int64,
	err error,
) {
	var accessUID, refreshUID uuid.UUID
	accessToken, accessUID, exp, err = j.CreateToken(
		user.ID,
		j.config.AccessTokenExpire,
		j.config.AccessTokenSecret,
		user.Role,
	)
	if err != nil {
		return nil, "", "", 0, errors.Wrap(err, "can't create access token")
	}

	refreshToken, refreshUID, _, err = j.CreateToken(
		user.ID,
		j.config.RefreshTokenExpire,
		j.config.RefreshTokenSecret,
		user.Role,
	)
	if err != nil {
		return nil, "", "", 0, errors.Wrap(err, "can't create refresh token")
	}

	cachedToken := &model.CachedTokens{
		AccessUID:  accessUID,
		RefreshUID: refreshUID,
	}

	return cachedToken, accessToken, refreshToken, exp, nil
}

func (j *JWT) ValidateToken(cachedToken model.CachedTokens, token JWTentity, isRefreshToken bool) error {
	var tokenUID uuid.UUID
	if isRefreshToken {
		tokenUID = cachedToken.RefreshUID
	} else {
		tokenUID = cachedToken.AccessUID
	}

	if tokenUID != token.UID {
		return errors.New("invalid token")
	}

	return nil
}

func (j *JWT) ParseToken(tokenString string, isRefreshToken bool) (JWTentity, error) {
	var secret string
	if isRefreshToken {
		secret = j.config.RefreshTokenSecret
	} else {
		secret = j.config.AccessTokenSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTentity{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return JWTentity{}, errors.Wrap(err, "can't parse token")
	}

	claims, ok := token.Claims.(*JWTentity)
	if !ok {
		return JWTentity{}, errors.New("can't parse token")
	}
	if !token.Valid {
		return JWTentity{}, errors.New("invalid token")
	}

	return *claims, nil
}

func (j *JWT) GenerateAndStoreTokenPair(ctx context.Context, user *model.User) (*model.Token, error) {
	cachedToken, accessToken, refreshToken, exp, err := j.GenerateTokenPair(*user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token pair")
	}

	if err := j.storeCacheTokens(
		ctx,
		cachedToken,
		user.ID,
		int(j.config.RefreshTokenExpire),
	); err != nil {
		return nil, errors.Wrap(err, "failed to store cache tokens")
	}

	return &model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Exp:          exp,
	}, nil
}

func (j *JWT) GetCachedTokens(ctx context.Context, userID uint) (*model.CachedTokens, error) {
	var cachedToken model.CachedTokens
	val, err := j.cache.Get(ctx, j.newTokenKey(userID)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cached token")
	}

	if err := json.Unmarshal([]byte(val), &cachedToken); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal cached token")
	}

	return &cachedToken, nil
}

func (j *JWT) RemoveToken(ctx context.Context, userID uint) error {
	return j.cache.Del(ctx, j.newTokenKey(userID)).Err()
}

func (j *JWT) newTokenKey(userID uint) string {
	return fmt.Sprintf("auth:token:%d", userID)
}

func (j *JWT) storeCacheTokens(
	ctx context.Context,
	tokens *model.CachedTokens,
	userID uint,
	ttl int,
) error {
	tokensJSON, err := json.Marshal(tokens)
	if err != nil {
		return errors.Wrap(err, "failed to marshal tokens")
	}

	return j.cache.Set(
		ctx,
		j.newTokenKey(userID),
		tokensJSON,
		time.Second*time.Duration(ttl),
	).Err()
}
