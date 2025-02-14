package photographer

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Get Citizen Card
// @Description		Get Photographer Citizen Card
// @Tags			photographer
// @Router			/api/v1/photographer/citizen-card [GET]
// @Security		ApiKeyAuth
// @Success			200	{object}	dto.HttpResponse{result=dto.CitizenCardResponse}
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetCitizenCard(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var photographer model.Photographer
	if err := h.store.DB.First(&photographer, "user_id = ?", userId).Error; err != nil {
		return errors.Wrap(err, "Photographer not found for user")
	}

	if photographer.CitizenCardID == nil {
		return apperror.NotFound("Citizen card is null", err)
	}

	var citizenCard model.CitizenCard
	if err := h.store.DB.First(&citizenCard, "id = ?", photographer.CitizenCardID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("Citizen card not found", err)
		}
		return errors.Wrap(err, "Error finding citizen card")
	}

	citizenCardDTO := dto.CitizenCardResponse{
		CitizenID:  citizenCard.CitizenID,
		LaserID:    citizenCard.LaserID,
		Picture:    citizenCard.Picture,
		ExpireDate: citizenCard.ExpireDate,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: citizenCardDTO,
	})
}
