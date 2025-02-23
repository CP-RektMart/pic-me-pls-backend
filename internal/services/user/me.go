package user

import (
	"context"
	"net/http"

	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

func (h *Handler) RegisterMe(api huma.API, middlewares huma.Middlewares) {
	huma.Register(api, huma.Operation{
		OperationID: "get-me",
		Method:      http.MethodGet,
		Path:        "/api/v1/me",
		Summary:     "Get my profile",
		Description: "Get my profile",
		Tags:        []string{"user"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleGetMe)
}

func (h *Handler) HandleGetMe(ctx context.Context, req *struct{}) (*dto.HumaHttpResponse[dto.UserResponse], error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	var user model.User
	result := h.store.DB.First(&user, userId)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get user")
	}

	response := dto.UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role.String(),
		Facebook:          user.Facebook,
		Instagram:         user.Instagram,
		Bank:              user.Bank,
		AccountNo:         user.AccountNo,
		BankBranch:        user.BankBranch,
	}

	return &dto.HumaHttpResponse[dto.UserResponse]{
		Body: dto.HttpResponse[dto.UserResponse]{
			Result: response,
		},
	}, nil
}
