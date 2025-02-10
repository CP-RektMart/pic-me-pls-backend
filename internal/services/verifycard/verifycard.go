package verifycard

import (
	"fmt"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// verifyCardHandler verifies the user's card information
// @Summary Verify user's card details
// @Description Verifies the user's card information and associates it with their account
// @Tags verifycard
// @Accept json
// @Produce json
// @Param verifyCardRequest body dto.VerifyCardRequest true "Card verification details"
// @Success 200 {object} dto.HttpResponse "Verification successful"
// @Failure 400 {object} dto.HttpResponse "Bad request. Card already verified or invalid data"
// @Failure 500 {object} dto.HttpResponse "Internal server error"
// @Router /api/v1/auth/verify [post]
func (h *Handler) HandlerVerifyCard(c *fiber.Ctx) error {
	if claims, _ := h.authMiddleware.GetJWTEntityFromContext(c.Context()); claims.Role != model.UserRolePhotographer {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: User should be photographer")
	}
	userID, err := h.authMiddleware.GetUserIDFromContext(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: User ID not found")
	}

	req := new(dto.VerifyCardRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
	}

	folder := strconv.FormatUint(uint64(userID), 10) + "/citizen_card/"
	signedURL, err := h.UploadCardFile(folder, c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("File upload failed: %v", err))
	}

	req.Picture = signedURL

	citizenCard, err := h.createCitizenCard(h.store.DB, req, userID)
	if err != nil {
		return err
	}

	return c.JSON(dto.HttpResponse{
		Result: citizenCard,
	})
}

func (h *Handler) UploadCardFile(folder string, c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("citizen_card")
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "failed to get file (citizen_card field is empty)")
	}

	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	var signedURL string
	if signedURL, err = h.store.Storage.UploadFile(c.Context(), folder+file.Filename, contentType, src, true); err != nil {
		return "", errors.Wrap(err, "failed to upload file")
	}

	return signedURL, nil
}

func (h *Handler) createCitizenCard(db *gorm.DB, req *dto.VerifyCardRequest, userID uint) (*model.CitizenCard, error) {
	var citizenCard model.CitizenCard

	// Find the photographer associated with the user
	var photographer model.Photographer
	if err := db.First(&photographer, "user_id = ?", userID).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Photographer not found for user: %v", err))
	}

	// Check if the photographer already has a CitizenCard
	if photographer.CitizenCardID != nil {
		// Return custom error code and message if already verified
		return nil, fiber.NewError(fiber.StatusBadRequest, "Already verified, photographer already has a citizen card")
	}

	// Create the CitizenCard using the request data
	citizenCard.CitizenID = req.CitizenID
	citizenCard.LaserID = req.LaserID
	citizenCard.Picture = req.Picture
	citizenCard.ExpireDate = req.ExpireDate

	// Insert the CitizenCard into the database
	if err := db.Create(&citizenCard).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error creating citizen card: %v", err))
	}

	// Update the photographer's CitizenCardID with the new CitizenCard ID
	photographer.CitizenCardID = &citizenCard.ID
	if err := db.Save(&photographer).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error updating photographer with citizen card: %v", err))
	}

	return &citizenCard, nil
}
