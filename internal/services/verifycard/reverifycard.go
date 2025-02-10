package verifycard

import (
	"fmt"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// HandlerReVerifyCard re-verifies the user's card information
// @Summary Re-verify user's card information
// @Description Re-verifies and updates the card details, associating it with the user's account
// @Tags verifycard
// @Accept json
// @Produce json
// @Param verifyCardRequest body dto.VerifyCardRequest true "Card re-verification details"
// @Success 200 {object} dto.HttpResponse "Card re-verification successful"
// @Failure 400 {object} dto.HttpResponse "Bad request. Invalid or incomplete data"
// @Failure 500 {object} dto.HttpResponse "Internal server error"
// @Router /api/v1/auth/reverify [patch]
func (h *Handler) HandlerReVerifyCard(c *fiber.Ctx) error {
	// TODO: get payload from jwt (middleware) or something -->
	email := "user3@example.com"
	req := new(dto.VerifyCardRequest)
	req.Picture = "path_to_picture2.jpg"
	req.LaserID = "LASER123"
	req.CitizenID = "2222999567819"
	req.ExpireDate = time.Now()

	// TODO: Upload Image

	updatedUser, err := h.updateCitizenCard(h.store.DB, req, email)
	if err != nil {
		return err
	}

	return c.JSON(dto.HttpResponse{
		Result: updatedUser, 
	})
}

func (h *Handler) updateCitizenCard(db *gorm.DB, req *dto.VerifyCardRequest, email string) (*model.CitizenCard, error) {
	// Start a new transaction
	tx := db.Begin()

	// Rollback if there's any error during the transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find the user by email within the transaction
	var user model.User
	if err := tx.First(&user, "email = ?", email).Error; err != nil {
		tx.Rollback() // Rollback the transaction if the user is not found
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error finding user by email: %v", err))
	}

	// Find the photographer associated with the user within the transaction
	var photographer model.Photographer
	if err := tx.First(&photographer, "user_id = ?", user.ID).Error; err != nil {
		tx.Rollback() // Rollback the transaction if the photographer is not found
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Photographer not found for user: %v", err))
	}

	// Find and delete the old CitizenCard within the transaction
	if photographer.CitizenCardID != nil {
		var oldCitizenCard model.CitizenCard

		if err := tx.First(&oldCitizenCard, "id = ?", *photographer.CitizenCardID).Error; err != nil {
			tx.Rollback() // Rollback the transaction if the old citizen card is not found
			return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error finding old citizen card: %v", err))
		}
		// Delete the old CitizenCard within the transaction
		if err := tx.Delete(&oldCitizenCard).Error; err != nil {
			tx.Rollback() // Rollback the transaction if there's an error deleting the old citizen card
			return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error deleting old citizen card: %v", err))
		}
	}

	// Create the new CitizenCard using the request data within the transaction
	newCitizenCard := model.CitizenCard{
		CitizenID:  req.CitizenID,
		LaserID:    req.LaserID,
		Picture:    req.Picture,
		ExpireDate: req.ExpireDate,
	}

	// Insert the new CitizenCard into the database within the transaction
	if err := tx.Create(&newCitizenCard).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error creating the new citizen card
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error creating new citizen card: %v", err))
	}

	// Update the photographer's CitizenCardID with the new CitizenCard ID within the transaction
	photographer.CitizenCardID = &newCitizenCard.ID
	if err := tx.Save(&photographer).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error updating the photographer
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error updating photographer with new citizen card: %v", err))
	}

	// Commit the transaction after all operations succeed
	if err := tx.Commit().Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error committing transaction: %v", err))
	}

	return &newCitizenCard, nil
}
