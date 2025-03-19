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
// @Description			Reverify Photographer Citizen Card
// @Tags			citizencard
// @Router			/api/v1/photographer/citizen-card/reverify [PATCH]
// @Security			ApiKeyAuth
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

	user, oldImageURL, err := h.updateCitizenCard(userID, req.ImageURL, req.CitizenID, req.LaserID, req.ExpireDate)
	if err != nil {
		if err := h.store.Storage.DeleteFile(ctx, req.ImageURL); err != nil {
			return errors.Wrap(err, "failed to deleting old picture")
		}
		return errors.Wrap(err, "failed to updating user profile")
	}

	if oldImageURL != "" && oldImageURL != req.ImageURL {
		if err := h.store.Storage.DeleteFile(ctx, oldImageURL); err != nil {
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
	var citizencard model.CitizenCard
	oldImageURL := ""

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// Check if photographer's citizen card is already verify
		if err := h.store.DB.Where("photographer_id = ?", userID).First(&citizencard).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrNoExistingCitizenCard
			}
			return errors.Wrap(err, "failed fetching citizen card")
		}

		oldImageURL = citizencard.Picture

		citizencard.CitizenID = citizenID
		citizencard.Picture = imageURL
		citizencard.LaserID = laserID
		citizencard.ExpireDate = expireDate

		if err := tx.Save(&citizencard).Error; err != nil {
			return errors.Wrap(err, "Error updating existing citizen card")
		}

		return nil
	}); err != nil {
		if errors.Is(err, ErrNoExistingCitizenCard) {
			return nil, "", apperror.BadRequest("no existing citizen card found", err)
		}
		return nil, "", errors.Wrap(err, "Error updating citizen card")
	}

	return &citizencard, oldImageURL, nil
}
