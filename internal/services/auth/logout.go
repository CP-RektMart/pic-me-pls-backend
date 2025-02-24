package auth

import (
	"context"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
)

func (h *Handler) RegisterLogout(api huma.API, middlewares ...func(ctx huma.Context, next func(huma.Context))) {
	huma.Register(api, huma.Operation{
		OperationID: "logout",
		Method:      http.MethodPost,
		Path:        "/api/v1/auth/logout",
		Summary:     "Logout",
		Description: "Logout",
		Tags:        []string{"auth"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleLogout)
}

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
