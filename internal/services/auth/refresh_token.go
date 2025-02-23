package auth

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandleRefreshToken(ctx context.Context, req *dto.HumaBody[dto.RefreshTokenRequest]) (*dto.HumaHttpResponse[dto.TokenResponse], error) {

	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request body", err)
	}

	jwtEntity, err := h.jwtService.ParseToken(req.Body.RefreshToken, true)
	if err != nil {
		return nil, huma.Error401Unauthorized("invalid token", err)
	}

	cachedToken, err := h.jwtService.GetCachedTokens(ctx, jwtEntity.ID)
	if err != nil {
		return nil, huma.Error401Unauthorized("invalid token", err)
	}

	if err := h.jwtService.ValidateToken(*cachedToken, jwtEntity, true); err != nil {
		return nil, huma.Error401Unauthorized("invalid token", err)
	}

	tokens, err := h.jwtService.GenerateAndStoreTokenPair(ctx, &model.User{
		Model: gorm.Model{ID: jwtEntity.ID},
		Role:  jwtEntity.Role,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token pair")
	}

	result := dto.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Exp:          tokens.Exp,
	}

	return &dto.HumaHttpResponse[dto.TokenResponse]{
		Body: dto.HttpResponse[dto.TokenResponse]{
			Result: result,
		},
	}, nil
}
