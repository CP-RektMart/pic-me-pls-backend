package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
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

func CreateToken(userID uint, expire int64, secret string, role model.UserRole) (token string, uid uuid.UUID, exp int64, err error) {
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

func GenerateTokenPair(user model.User, accessTokenSecret, refreshTokenSecret string, accessTokenExpire, refreshTokenExpire int64) (
	cahcedToken *model.CachedTokens,
	accessToken string,
	refreshToken string,
	exp int64,
	err error,
) {
	var accessUID, refreshUID uuid.UUID
	accessToken, accessUID, exp, err = CreateToken(user.ID, accessTokenExpire, accessTokenSecret, user.Role)
	if err != nil {
		return nil, "", "", 0, errors.Wrap(err, "can't create access token")
	}

	refreshToken, refreshUID, _, err = CreateToken(user.ID, refreshTokenExpire, refreshTokenSecret, user.Role)
	if err != nil {
		return nil, "", "", 0, errors.Wrap(err, "can't create refresh token")
	}

	cachedToken := &model.CachedTokens{
		AccessUID:  accessUID,
		RefreshUID: refreshUID,
	}

	return cachedToken, accessToken, refreshToken, exp, nil
}

func ValidateToken(cachedToken model.CachedTokens, token JWTentity, isRefreshToken bool) error {
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

func ParseToken(tokenString string, secret string) (JWTentity, error) {
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

func NewTokenKey(userID uint) string {
	return fmt.Sprintf("auth:token:%d", userID)
}

func StoreCacheTokens(ctx context.Context, cache *redis.Client, tokens *model.CachedTokens, userID uint, ttl int) error {
	tokensJSON, err := json.Marshal(tokens)
	if err != nil {
		return errors.Wrap(err, "failed to marshal tokens")
	}

	return cache.Set(ctx, NewTokenKey(userID), tokensJSON, time.Second*time.Duration(ttl)).Err()
}

func GetCachedTokens(ctx context.Context, cache *redis.Client, userID uint) (*model.CachedTokens, error) {
	var cachedToken *model.CachedTokens
	val, err := cache.Get(ctx, NewTokenKey(userID)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cached token")
	}

	if err := json.Unmarshal([]byte(val), cachedToken); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal cached token")
	}

	return cachedToken, nil
}

func GenerateAndStoreTokenPair(
	ctx context.Context,
	cache *redis.Client,
	user model.User,
	accessTokenSecret,
	refreshTokenSecret string,
	accessTokenExpire,
	refreshTokenExpire int64,
) (*dto.TokenResponse, error) {
	cachedToken, accessToken, refreshToken, exp, err := GenerateTokenPair(
		user,
		accessTokenSecret,
		refreshTokenSecret,
		accessTokenExpire,
		refreshTokenExpire,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token pair")
	}

	if err := StoreCacheTokens(
		ctx,
		cache,
		cachedToken,
		user.ID,
		int(refreshTokenExpire),
	); err != nil {
		return nil, errors.Wrap(err, "failed to store cache tokens")
	}

	return &dto.TokenResponse{
		AcessToken:   accessToken,
		RefreshToken: refreshToken,
		Exp:          exp,
	}, nil
}
