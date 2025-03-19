package citizencard

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Get Citizen Card
// @Description			Get Photographer Citizen Card
// @Tags			citizencard
// @Router			/api/v1/photographer/citizen-card [GET]
// @Security			ApiKeyAuth
// @Success			200	{object}	dto.HttpResponse[dto.CitizenCardResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetCitizenCard(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	citizenCard, err := h.getCitizenCard(userID)
	if err != nil {
		return errors.Wrap(err, "failed get citizen card")
	}

	citizenCardDTO := dto.CitizenCardResponse{
		CitizenID:  citizenCard.CitizenID,
		LaserID:    citizenCard.LaserID,
		Picture:    citizenCard.Picture,
		ExpireDate: citizenCard.ExpireDate,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.CitizenCardResponse]{
		Result: citizenCardDTO,
	})
}

func (h *Handler) getCitizenCard(photographerID uint) (*model.CitizenCard, error) {
	var citizenCard model.CitizenCard
	if err := h.store.DB.Where("photographer_id = ?", photographerID).First(&citizenCard).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.NotFound("citizen card not found", err)
		}
		return nil, errors.Wrap(err, "failed fetch citizen card")
	}
	return &citizenCard, nil
}
