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
// @Description			Verify Photographer Citizen Card
// @Tags			citizencard
// @Router			/api/v1/photographer/citizen-card/verify [POST]
// @Security			ApiKeyAuth
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

	response := dto.ToCitizenCardResponse(user)

	return c.JSON(dto.HttpResponse[dto.CitizenCardResponse]{
		Result: response,
	})
}

func (h *Handler) createCitizenCard(userID uint, imageURL, citizenID, laserID string, expireDate time.Time) (*model.CitizenCard, error) {
	var citizenCard model.CitizenCard

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// Check if photographer's citizen card is already verify
		err := h.store.DB.Where("photographer_id = ?", userID).First(&citizenCard).Error
		if err == nil {
			return ErrAlreadyVerified
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrap(err, "failed fetch citizen card")
		}

		// Create the CitizenCard using the request data
		citizenCard.CitizenID = citizenID
		citizenCard.LaserID = laserID
		citizenCard.Picture = imageURL
		citizenCard.ExpireDate = expireDate
		citizenCard.PhotographerID = userID

		// Insert the CitizenCard into the database
		if err := tx.Create(&citizenCard).Error; err != nil {
			return errors.Wrap(err, "Error creating citizen card")
		}

		photographer := model.Photographer{
			UserID:       userID,
			ActiveStatus: true,
			IsVerified:   true,
		}

		// update photographer status
		if err := tx.Model(&photographer).Updates(&photographer).Error; err != nil {
			return errors.Wrap(err, "failed update photographer status")
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
