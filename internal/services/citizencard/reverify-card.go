package citizencard

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Reverify Citizen Card
// @Description		Reverify Photographer Citizen Card
// @Tags			citizencard
// @Router			/api/v1/photographer/citizen-card/reverify [PATCH]
// @Security		ApiKeyAuth
// @Param 			RequestBody 	body 	dto.ReVerifyCitizenCardRequest 	true 	"request request"
// @Success			200	{object}	dto.HttpResponse[dto.CitizenCardResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleReVerifyCard(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.ReVerifyCitizenCardRequest)
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	user, oldImageUrl, err := h.updateCitizenCard(userID, req.ImageURL, req.CitizenID, req.LaserID, req.ExpireDate)
	if err != nil {
		if err := h.store.Storage.DeleteFile(ctx, req.ImageURL); err != nil {
			return errors.Wrap(err, "failed to deleting old picture")
		}
		return errors.Wrap(err, "failed to updating user profile")
	}

	if oldImageUrl != "" && oldImageUrl != req.ImageURL {
		if err := h.store.Storage.DeleteFile(ctx, oldImageUrl); err != nil {
			return errors.Wrap(err, "failed to deleting old picture")
		}
	}

	response := dto.CitizenCardResponse{
		CitizenID:  user.CitizenID,
		LaserID:    user.LaserID,
		Picture:    user.Picture,
		ExpireDate: user.ExpireDate,
	}

	return c.JSON(dto.HttpResponse[dto.CitizenCardResponse]{
		Result: response,
	})
}

func (h *Handler) updateCitizenCard(userID uint, imageURL, citizenID, laserID string, expireDate time.Time) (*model.CitizenCard, string, error) {
	var updatedCitizenCard model.CitizenCard
	oldImageUrl := ""

	err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userID).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		if photographer.CitizenCardID == nil {
			return errors.Wrap(ErrNoExistingCitizenCard, "No existing citizen card found for the photographer")
		}

		var existingCitizenCard model.CitizenCard
		if err := tx.First(&existingCitizenCard, "id = ?", *photographer.CitizenCardID).Error; err != nil {
			return errors.Wrap(err, "Error finding existing citizen card")
		}

		oldImageUrl = existingCitizenCard.Picture

		existingCitizenCard.CitizenID = citizenID
		existingCitizenCard.Picture = imageURL
		existingCitizenCard.LaserID = laserID
		existingCitizenCard.ExpireDate = expireDate

		if err := tx.Save(&existingCitizenCard).Error; err != nil {
			return errors.Wrap(err, "Error updating existing citizen card")
		}

		updatedCitizenCard = existingCitizenCard

		// Update the photographer's CitizenCardID if necessary
		photographer.CitizenCardID = &updatedCitizenCard.ID
		if err := tx.Save(&photographer).Error; err != nil {
			return errors.Wrap(err, "Error updating photographer with citizen card")
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, ErrNoExistingCitizenCard) {
			return nil, "", apperror.BadRequest("no existing citizen card found", err)
		}
		return nil, "", errors.Wrap(err, "Error updating citizen card")
	}

	return &updatedCitizenCard, oldImageUrl, nil
}
