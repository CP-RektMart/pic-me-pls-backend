package photographer

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// handlerGetMe godoc
// @summary Get citizen card
// @description Retrieves the authenticated phtographer's citizen card
// @tags user
// @security Bearer
// @id get-citizen-card
// @accept json
// @produce json
// @success 200 {object} dto.BaseUserDTO "OK"
// @failure 400 {object} dto.HttpResponse "Bad Request"
// @failure 500 {object} dto.HttpResponse "Internal Server Error"
// @Router /api/v1/me [GET]
func (h *Handler) HandleGetCitizenCard(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var photographer model.Photographer
	if err := h.store.DB.First(&photographer, "user_id = ?", userId).Error; err != nil {
		return errors.Wrap(err, "Photographer not found for user")
	}

	var citizenCard model.CitizenCard
	if err := h.store.DB.First(&citizenCard, "id = ?", *photographer.CitizenCardID).Error; err != nil {
		return errors.Wrap(err, "Error finding citizen card")
	}

	citizenCardDTO := dto.CitizenCard{
		CitizenID:  citizenCard.CitizenID,
		LaserID:    citizenCard.LaserID,
		Picture:    citizenCard.Picture,
		ExpireDate: citizenCard.ExpireDate,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: citizenCardDTO,
	})
}
