package photographer

import (
	"fmt"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
	claims, err := h.authMiddleware.GetJWTEntityFromContext(c.UserContext())
	if claims.Role != model.UserRolePhotographer {
		return apperror.Forbidden("FORBIDDEN", fmt.Errorf("user is not photographer"))
	}
	userID := claims.ID
	req := new(dto.VerifyCardRequest)

	if err := c.BodyParser(req); err != nil {
		return errors.Wrap(err, "invalid request body")
	}

	folder := "/citizen_card/" + strconv.FormatUint(uint64(userID), 10)
	signedURL, isUploaded, err := h.uploadCardFile(c, folder)
	if err != nil {
		return errors.Wrap(err, "File upload failed")
	}
	if isUploaded {
		req.Picture = signedURL
	} else {
		return errors.Wrap(errors.Errorf("Field Missing"), "citizen_card is require")
	}

	err = h.createCitizenCard(req, userID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) uploadCardFile(c *fiber.Ctx, folder string) (string, bool, error) {
	file, err := c.FormFile("citizen_card")
	if err != nil {
		// This allows because if the field is not provided mean they dont change the picutre
		return "", false, nil
	}

	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", false, errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	signedURL, err := h.store.Storage.UploadFile(c.UserContext(), folder+file.Filename, contentType, src, true)
	if err != nil {

		return "", false, errors.Wrap(err, "failed to upload file")
	}

	return signedURL, true, nil
}

func (h *Handler) createCitizenCard(req *dto.VerifyCardRequest, userID uint) error {
	var citizenCard model.CitizenCard

	// Find the photographer associated with the user
	var photographer model.Photographer
	if err := h.store.DB.First(&photographer, "user_id = ?", userID).Error; err != nil {
		return errors.Wrap(err, "Photographer not found for user")
	}

	// Check if the photographer already has a CitizenCard
	if photographer.CitizenCardID != nil {
		// Return custom error code and message if already verified
		return errors.Wrap(errors.Errorf("Already verified"), "photographer already has a citizen card")
	}

	// Create the CitizenCard using the request data
	citizenCard.CitizenID = req.CitizenID
	citizenCard.LaserID = req.LaserID
	citizenCard.Picture = req.Picture
	citizenCard.ExpireDate = req.ExpireDate

	// Insert the CitizenCard into the database
	if err := h.store.DB.Create(&citizenCard).Error; err != nil {
		return errors.Wrap(err, "Error creating citizen card")
	}

	// Update the photographer's CitizenCardID with the new CitizenCard ID
	photographer.CitizenCardID = &citizenCard.ID
	if err := h.store.DB.Save(&photographer).Error; err != nil {
		return errors.Wrap(err, "Error updating photographer with citizen card")
	}

	return nil
}
