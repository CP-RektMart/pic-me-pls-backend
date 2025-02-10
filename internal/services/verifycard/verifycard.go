package verifycard

import (
	"fmt"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
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
	// TODO: get payload from jwt (middleware) or something -->
	email := "user3@example.com"
	req := new(dto.VerifyCardRequest)
	req.Picture = "path_to_picture.jpg"
	req.LaserID = "LASER123"
	req.CitizenID = "1519999567819"
	req.ExpireDate = time.Now()

	// TODO: Upload Image

	citizenCard, err := h.createCitizenCard(h.store.DB, req, email)
	if err != nil {
		return err
	}

	return c.JSON(dto.HttpResponse{
		Result: citizenCard,
	})
}

func (h *Handler) createCitizenCard(db *gorm.DB, req *dto.VerifyCardRequest, email string) (*model.CitizenCard, error) {
	var citizenCard model.CitizenCard

	// Find the user by email
	var user model.User
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error finding user by email: %v", err))
	}

	// Find the photographer associated with the user
	var photographer model.Photographer
	if err := db.First(&photographer, "user_id = ?", user.ID).Error; err != nil {
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
