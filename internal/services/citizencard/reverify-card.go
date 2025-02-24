package citizencard

import (
	"mime/multipart"
	"path"
	"strconv"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Reverify Citizen Card
// @Description		Reverify Photographer Citizen Card
// @Tags			photographer
// @Router			/api/v1/photographer/reverify [PATCH]
// @Security		ApiKeyAuth
// @Accept			multipart/form-data
// @Param 			cardPicture 	formData 	file		false	"Card picture (optional)"
// @Param 			citizenId 		formData 	string		true	"Citizen ID"
// @Param 			laserId 		formData 	string		true	"Laser ID"
// @Param 			expireDate 		formData 	string		true	"Expire Date"
// @Success			200	{object}	dto.HttpResponse[dto.CitizenCardResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleReVerifyCard(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.CitizenCardRequest)
	req.CitizenID = c.FormValue("citizenId")
	req.LaserID = c.FormValue("laserId")
	req.ExpireDate, err = time.Parse(time.RFC3339, c.FormValue("expireDate"))
	if err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	var file *multipart.FileHeader
	file, err = c.FormFile("cardPicture")
	// if error mean cannot get file just ignore.
	// because field is not provide mean not change.
	var signedURL string = ""
	if err == nil && file.Size != 0 {
		signedURL, err = h.uploadCardFile(c.UserContext(), file, citizenCardFolder(userId))
		if err != nil {
			return errors.Wrap(err, "File upload failed")
		}
	}

	// var oldPictureURL string
	user, err := h.updateCitizenCard(req, userId, signedURL, nil)
	if err != nil {
		if signedURL != "" {
			err = h.store.Storage.DeleteFile(c.UserContext(), citizenCardFolder(userId)+path.Base(signedURL))
			if err != nil {
				return errors.Wrap(err, "Fail to delete the picture")
			}
		}
		return errors.Wrap(err, "Error updating user profile")
	}

	// if oldPictureURL != "" && oldPictureURL != signedURL {
	// 	err = h.store.Storage.DeleteFile(c.UserContext(), citizenCardFolder(userId)+path.Base(oldPictureURL))
	// 	if err != nil {
	// 		return errors.Wrap(err, "Fail to delete old picture")
	// 	}
	// }

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

func (h *Handler) updateCitizenCard(req *dto.CitizenCardRequest, userId uint, signedURL string, oldPictureURL *string) (*model.CitizenCard, error) {
	var updatedCitizenCard model.CitizenCard

	updateField := func(field *string, newValue string) {
		if newValue != "" {
			*field = newValue
		}
	}

	err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userId).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		// Check if the photographer has an existing CitizenCard
		if photographer.CitizenCardID != nil {
			// Find the existing CitizenCard
			var existingCitizenCard model.CitizenCard
			if err := tx.First(&existingCitizenCard, "id = ?", *photographer.CitizenCardID).Error; err != nil {
				return errors.Wrap(err, "Error finding existing citizen card")
			}

			// Optionally assign old picture URL if needed
			if oldPictureURL != nil {
				*oldPictureURL = existingCitizenCard.Picture
			}

			existingCitizenCard.CitizenID = req.CitizenID
			existingCitizenCard.LaserID = req.LaserID
			updateField(&existingCitizenCard.Picture, signedURL)
			existingCitizenCard.ExpireDate = req.ExpireDate

			// Save the updated CitizenCard
			if err := tx.Save(&existingCitizenCard).Error; err != nil {
				return errors.Wrap(err, "Error updating existing citizen card")
			}

			// Return the updated citizen card
			updatedCitizenCard = existingCitizenCard
		} else {
			// If no CitizenCard exists, create a new one (this block can be omitted if you want to handle that elsewhere)
			return errors.New("No existing citizen card found for the photographer")
		}

		// Update the photographer's CitizenCardID if necessary
		photographer.CitizenCardID = &updatedCitizenCard.ID
		if err := tx.Save(&photographer).Error; err != nil {
			return errors.Wrap(err, "Error updating photographer with citizen card")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &updatedCitizenCard, nil
}

func citizenCardFolder(userId uint) string {
	return "citizen_card/" + strconv.FormatUint(uint64(userId), 10) + "/"
}
