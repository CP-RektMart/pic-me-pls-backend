package photographer

import (
	"context"
	"path"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandleReVerifyCard(ctx context.Context, req *dto.HumaFormData[dto.CitizenCardRequest]) (*dto.HumaHttpResponse[dto.CitizenCardResponse], error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}
	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request", err)
	}

	file, ok := req.RawBody.Form.File["profilePicture"]
	if !ok {
		return nil, huma.Error400BadRequest("invalid request", errors.New("profilePicture is required"))
	}

	// if error mean cannot get file just ignore.
	// because field is not provide mean not change.
	var signedURL string
	if len(file) > 0 {
		signedURL, err = h.uploadCardFile(ctx, file[0], citizenCardFolder(userId))
		if err != nil {
			return nil, errors.Wrap(err, "File upload failed")
		}
	}

	data := req.RawBody.Data()
	// var oldPictureURL string
	user, err := h.updateCitizenCard(data, userId, signedURL, nil)
	if err != nil {
		if signedURL != "" {
			err = h.store.Storage.DeleteFile(ctx, citizenCardFolder(userId)+path.Base(signedURL))
			if err != nil {
				return nil, errors.Wrap(err, "Fail to delete the picture")
			}
		}
		return nil, errors.Wrap(err, "Error updating user profile")
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

	return &dto.HumaHttpResponse[dto.CitizenCardResponse]{
		Body: dto.HttpResponse[dto.CitizenCardResponse]{
			Result: response,
		},
	}, nil
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
