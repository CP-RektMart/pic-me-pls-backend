package photographer

import (
	"context"
	"mime/multipart"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// verifyCardHandler verifies the user's card information
// @Summary Verify user's card details
// @Description Verifies the user's card information and associates it with their account
// @Tags photographer
// @Accept json
// @Produce json
// @Param verifyCardRequest body dto.VerifyCardRequest true "Card verification details"
// @Success 200 {object} dto.HttpResponse "Verification successful"
// @Failure 400 {object} dto.HttpResponse "Bad request. Card already verified or invalid data"
// @Failure 500 {object} dto.HttpResponse "Internal server error"
// @Router /api/v1/auth/verify [post]
func (h *Handler) HandleVerifyCard(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.VerifyCardRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	file, err := c.FormFile("citizen_card")
	if err != nil {
		return apperror.BadRequest("citizen_card is require", errors.Errorf("Field Missing"))
	}

	signedURL, err := h.uploadCardFile(c.UserContext(), file, citizenCardFolder(userId))
	if err != nil {
		return errors.Wrap(err, "File upload failed")
	}
	req.Picture = signedURL

	err = h.createCitizenCard(req, userId)
	if err != nil {
		return errors.Wrap(err, "Fail to create citizen card")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) uploadCardFile(c context.Context, file *multipart.FileHeader, folder string) (string, error) {
	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	signedURL, err := h.store.Storage.UploadFile(c, folder+file.Filename, contentType, src, true)
	if err != nil {

		return "", errors.Wrap(err, "failed to upload file")
	}

	return signedURL, nil
}

func (h *Handler) createCitizenCard(req *dto.VerifyCardRequest, userId uint) error {
	var citizenCard model.CitizenCard

	err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userId).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		// Check if the photographer already has a CitizenCard
		if photographer.CitizenCardID != nil {
			return apperror.BadRequest("photographer already has a citizen card", errors.Errorf("Already verified"))
		}

		// Create the CitizenCard using the request data
		citizenCard.CitizenID = req.CitizenID
		citizenCard.LaserID = req.LaserID
		citizenCard.Picture = req.Picture
		citizenCard.ExpireDate = req.ExpireDate

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
	})

	return err
}
