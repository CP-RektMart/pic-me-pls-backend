package auth

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/pkg/errors"
)

func (h *Handler) HandleLogout(ctx context.Context, req *struct{}) (*dto.HumaHttpResponse[dto.TokenResponse], error) {
	userID, err := h.authmiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}

	if err := h.jwtService.RemoveToken(ctx, userID); err != nil {
		return nil, errors.Wrap(err, "failed to remove token")
	}

	return nil, nil
}
