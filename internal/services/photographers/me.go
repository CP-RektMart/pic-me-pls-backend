package photographers

import (
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Get me (photographer)
// @Tags			photographers
// @Router			/api/v1/photographer/me [GET]
// @Security		ApiKeyAuth
// @Success			200	{object}	dto.HttpResponse[dto.PhotographerMeResponse]
// @Failure     	400   {object}  dto.HttpError
// @Failure     	401   {object}  dto.HttpError
// @Failure     	403   {object}  dto.HttpError
// @Failure     	404   {object}  dto.HttpError
// @Failure     	500   {object}  dto.HttpError
func (h *Handler) HandleGetMe(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	photographer, err := h.getMe(userID)
	if err != nil {
		return errors.Wrap(err, "failed fetching photographer profile")
	}

	resp := dto.ToPhotographerMeResponse(photographer)

	return c.Status(fiber.StatusOK).JSON(dto.Success(resp))
}

func (h *Handler) getMe(userID uint) (model.Photographer, error) {
	var photographer model.Photographer
	if err := h.store.DB.Joins("User").First(&photographer, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Photographer{}, apperror.NotFound("photographer not found", err)
		}
		return model.Photographer{}, err
	}
	return photographer, nil
}
