package photographer

import (
	"fmt"
	"path"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// HandlerReVerifyCard re-verifies the user's card information
// @Summary Re-verify user's card information
// @Description Re-verifies and updates the card details, associating it with the user's account
// @Tags photographer
// @Accept json
// @Produce json
// @Param verifyCardRequest body dto.VerifyCardRequest true "Card re-verification details"
// @Success 200 {object} dto.HttpResponse "Card re-verification successful"
// @Failure 400 {object} dto.HttpResponse "Bad request. Invalid or incomplete data"
// @Failure 500 {object} dto.HttpResponse "Internal server error"
// @Router /api/v1/auth/reverify [patch]
func (h *Handler) HandleReVerifyCard(c *fiber.Ctx) error {
	claims, err := h.authMiddleware.GetJWTEntityFromContext(c.UserContext())
	if claims.Role != model.UserRolePhotographer {
		return apperror.Forbidden("FORBIDDEN", fmt.Errorf("user is not photographer"))
	}
	userID := claims.ID
	req := new(dto.VerifyCardRequest)

	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	signedURL, fileUploaded, err := h.uploadCardFile(c, citizenCardFolder(userID))
	if err != nil {
		return errors.Wrap(err, "File upload failed")
	}

	if fileUploaded {
		req.Picture = signedURL
	}

	var oldPictureURL string
	updatedUser, err := h.updateCitizenCard(req, userID, &oldPictureURL)
	if err != nil {
		h.store.Storage.DeleteFile(c.UserContext(), citizenCardFolder(userID)+path.Base(signedURL))
		return err
	}

	if oldPictureURL != "" && oldPictureURL != req.Picture {
		h.store.Storage.DeleteFile(c.UserContext(), citizenCardFolder(userID)+path.Base(oldPictureURL))
	}

	return c.JSON(dto.HttpResponse{
		Result: updatedUser,
	})
}

func (h *Handler) updateCitizenCard(req *dto.VerifyCardRequest, userID uint, oldPictureURL *string) (*model.CitizenCard, error) {
	// Start a new transaction
	tx := h.store.DB.Begin()

	// Rollback if there's any error during the transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Find the photographer associated with the user within the transaction
	var photographer model.Photographer
	if err := tx.First(&photographer, "user_id = ?", userID).Error; err != nil {
		tx.Rollback() // Rollback the transaction if the photographer is not found
		return nil, errors.Wrap(err, "Photographer not found for user")
	}

	// Find and delete the old CitizenCard within the transaction
	if photographer.CitizenCardID != nil {
		var oldCitizenCard model.CitizenCard

		if err := tx.First(&oldCitizenCard, "id = ?", *photographer.CitizenCardID).Error; err != nil {
			tx.Rollback() // Rollback the transaction if the old citizen card is not found
			return nil, errors.Wrap(err, "Error finding old citizen card")
		}
		if oldPictureURL != nil {
			oldPictureURL = &oldCitizenCard.Picture
		}
		// Delete the old CitizenCard within the transaction
		if err := tx.Delete(&oldCitizenCard).Error; err != nil {
			tx.Rollback() // Rollback the transaction if there's an error deleting the old citizen card
			return nil, errors.Wrap(err, "Error deleting old citizen card")
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
		return nil, errors.Wrap(err, "Error creating new citizen card")
	}

	// Update the photographer's CitizenCardID with the new CitizenCard ID within the transaction
	photographer.CitizenCardID = &newCitizenCard.ID
	if err := tx.Save(&photographer).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error updating the photographer
		return nil, errors.Wrap(err, "Error updating photographer with new citizen card")
	}

	// Commit the transaction after all operations succeed
	if err := tx.Commit().Error; err != nil {
		return nil, errors.Wrap(err, "Error committing transaction")
	}

	return &newCitizenCard, nil
}

func citizenCardFolder(userID uint) string {
	return "/citizen_card/" + strconv.FormatUint(uint64(userID), 10)
}
