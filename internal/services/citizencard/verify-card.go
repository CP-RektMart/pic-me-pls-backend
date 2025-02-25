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

// @Summary			Verify Citizen Card
// @Description		Verify Photographer Citizen Card
// @Tags			citizencard
// @Router			/api/v1/photographer/citizen-card/verify [POST]
// @Security		ApiKeyAuth
// @Param 			RequestBody 	body 	dto.VerifyCitizenCardRequest 	true 	"request request"
// @Success			200	{object}	dto.HttpResponse[dto.CitizenCardResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleVerifyCard(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.VerifyCitizenCardRequest)
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	user, err := h.createCitizenCard(userID, req.ImageURL, req.CitizenID, req.LaserID, req.ExpireDate)
	if err != nil {
		return errors.Wrap(err, "Fail to create citizen card")
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

func (h *Handler) createCitizenCard(userId uint, imageURL, citizenID, laserID string, expireDate time.Time) (*model.CitizenCard, error) {
	var citizenCard model.CitizenCard

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userId).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		// Check if the photographer already has a CitizenCard
		if photographer.CitizenCardID != nil {
			return errors.Wrap(ErrAlreadyVerified, "Photographer already has a citizen card")
		}

		// Create the CitizenCard using the request data
		citizenCard.CitizenID = citizenID
		citizenCard.LaserID = laserID
		citizenCard.Picture = imageURL
		citizenCard.ExpireDate = expireDate

		// Insert the CitizenCard into the database
		if err := tx.Create(&citizenCard).Error; err != nil {
			return errors.Wrap(err, "Error creating citizen card")
		}

		// Update the photographer's CitizenCardID with the new CitizenCard ID
		photographer.CitizenCardID = &citizenCard.ID
		if err := tx.Save(&photographer).Error; err != nil {
			return errors.Wrap(err, "Error updating photographer with citizen card")
		}

		return nil
	}); err != nil {
		if errors.Is(err, ErrAlreadyVerified) {
			return nil, apperror.BadRequest("photographer already has a citizen card", errors.Errorf("Already verified"))
		}
		return nil, errors.Wrap(err, "failed to create citizen card")
	}

	return &citizenCard, nil
}
