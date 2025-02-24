package verify

import (
	"context"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) RegisterGetCitizenCard(api huma.API, middlewares ...func(ctx huma.Context, next func(huma.Context))) {
	huma.Register(api, huma.Operation{
		OperationID: "get-citizen-card",
		Method:      http.MethodGet,
		Path:        "/api/v1/photographer/citizen-card",
		Summary:     "Get citizen card",
		Description: "Get citizen card",
		Tags:        []string{"verify"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleGetCitizenCard)
}

func (h *Handler) HandleGetCitizenCard(ctx context.Context, req *struct{}) (*dto.HumaHttpResponse[dto.CitizenCardResponse], error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	var photographer model.Photographer
	if err := h.store.DB.First(&photographer, "user_id = ?", userId).Error; err != nil {
		return nil, errors.Wrap(err, "Photographer not found for user")
	}

	if photographer.CitizenCardID == nil {
		return nil, apperror.NotFound("Citizen card is null", err)
	}

	var citizenCard model.CitizenCard
	if err := h.store.DB.First(&citizenCard, "id = ?", photographer.CitizenCardID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("Citizen card not found", err)
		}
		return nil, errors.Wrap(err, "Error finding citizen card")
	}

	citizenCardDTO := dto.CitizenCardResponse{
		CitizenID:  citizenCard.CitizenID,
		LaserID:    citizenCard.LaserID,
		Picture:    citizenCard.Picture,
		ExpireDate: citizenCard.ExpireDate,
	}

	return &dto.HumaHttpResponse[dto.CitizenCardResponse]{
		Body: dto.HttpResponse[dto.CitizenCardResponse]{
			Result: citizenCardDTO,
		},
	}, nil
}
