package photographer

import (
	"path"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	// if error mean cannot get file just ignore.
	// because field is not provide mean not change.
	if err == nil {
		signedURL, err := h.uploadCardFile(c.UserContext(), file, citizenCardFolder(userId))
		if err != nil {
			return errors.Wrap(err, "File upload failed")
		}
		req.Picture = signedURL
	}

	var oldPictureURL string
	updatedUser, err := h.updateCitizenCard(req, userId, &oldPictureURL)
	if err != nil {
		if req.Picture != "" {
			err = h.store.Storage.DeleteFile(c.UserContext(), citizenCardFolder(userId)+path.Base(req.Picture))
			if err != nil {
				return errors.Wrap(err, "Fail to delete the picture")
			}
		}
		return errors.Wrap(err, "Error updating user profile")
	}

	if oldPictureURL != "" && oldPictureURL != req.Picture {
		err = h.store.Storage.DeleteFile(c.UserContext(), citizenCardFolder(userId)+path.Base(oldPictureURL))
		if err != nil {
			return errors.Wrap(err, "Fail to delete old picture")
		}
	}

	return c.JSON(dto.HttpResponse{
		Result: updatedUser,
	})
}

func (h *Handler) updateCitizenCard(req *dto.VerifyCardRequest, userId uint, oldPictureURL *string) (*model.CitizenCard, error) {
	var newCitizenCard model.CitizenCard

	err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userId).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		// If there's an existing CitizenCard, delete it
		// (Photographer is created before citizen card in a first place)
		if photographer.CitizenCardID != nil {
			err := tx.Transaction(func(tx2 *gorm.DB) error {
				var oldCitizenCard model.CitizenCard
				if err := tx2.First(&oldCitizenCard, "id = ?", *photographer.CitizenCardID).Error; err != nil {
					return errors.Wrap(err, "Error finding old citizen card")
				}
				// Assign old picture URL if needed
				if oldPictureURL != nil {
					*oldPictureURL = oldCitizenCard.Picture
				}
				if err := tx2.Delete(&oldCitizenCard).Error; err != nil {
					return errors.Wrap(err, "Error deleting old citizen card")
				}
				return nil
			})
			if err != nil {
				return err
			}
		}

		// Create a new CitizenCard
		err := tx.Transaction(func(tx3 *gorm.DB) error {
			newCitizenCard = model.CitizenCard{
				CitizenID:  req.CitizenID,
				LaserID:    req.LaserID,
				Picture:    req.Picture,
				ExpireDate: req.ExpireDate,
			}
			if err := tx3.Create(&newCitizenCard).Error; err != nil {
				return errors.Wrap(err, "Error creating new citizen card")
			}
			return nil
		})
		if err != nil {
			return err // Rollback if creating new CitizenCard fails
		}

		// Update the photographer's CitizenCardID
		photographer.CitizenCardID = &newCitizenCard.ID
		if err := tx.Save(&photographer).Error; err != nil {
			return errors.Wrap(err, "Error updating photographer with new citizen card")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &newCitizenCard, nil
}

func citizenCardFolder(userId uint) string {
	return "/citizen_card/" + strconv.FormatUint(uint64(userId), 10)
}
