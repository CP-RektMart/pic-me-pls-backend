package auth

import (
	"context"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) RegisterLogin(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "login",
		Method:      http.MethodPost,
		Path:        "/api/v1/auth/login",
		Summary:     "Login",
		Description: "Login",
		Tags:        []string{"auth"},
	}, h.HandleLogin)
}

func (h *Handler) HandleLogin(ctx context.Context, req *dto.HumaBody[dto.LoginRequest]) (*dto.HumaHttpResponse[dto.LoginResponse], error) {
	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request", err)
	}

	OAuthUser, err := h.validateIDToken(ctx, req.Body.IDToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate id token")
	}

	var user *model.User
	var token *model.Token

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("email = ?", OAuthUser.Email).First(&user).Error; err != nil {
			return errors.Wrap(err, "failed getting user")
		}

		token, err = h.jwtService.GenerateAndStoreTokenPair(ctx, user)
		if err != nil {
			return errors.Wrap(err, "failed to generate token pair")
		}

		return nil
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("user not found")
		}
		return nil, errors.Wrap(err, "failed to get user and token")
	}

	result := dto.LoginResponse{
		TokenResponse: dto.TokenResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			Exp:          token.Exp,
		},
		User: dto.ToUserResponse(*user),
	}

	return &dto.HumaHttpResponse[dto.LoginResponse]{
		Body: dto.HttpResponse[dto.LoginResponse]{
			Result: result,
		},
	}, nil
}
