package user

import (
	"context"
	"mime/multipart"
	"path"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Update me
// @Description		Update user's profile
// @Tags			user
// @Router			/api/v1/me [PATCH]
// @Security		ApiKeyAuth
// @Accept			multipart/form-data
// @Param 			profilePicture 	formData 	file		false	"Profile picture (optional)"
// @Param 			name 			formData 	string		false	"Name"
// @Param 			phoneNumber 	formData 	string		false	"Phone Number"
// @Param 			facebook 		formData 	string		false	"Facebook"
// @Param 			instagram 		formData 	string		false	"Instagram"
// @Param 			bank 			formData 	string		false	"Bank"
// @Param 			accountNo 		formData 	string		false	"Account No"
// @Param 			bankBranch 		formData 	string		false	"Bank Branch"
// @Success			200	{object}	dto.HttpResponse[dto.UserResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateMe(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.UserUpdateRequest)
	req.Name = c.FormValue("name")
	req.PhoneNumber = c.FormValue("phoneNumber")
	req.Facebook = c.FormValue("facebook")
	req.Instagram = c.FormValue("instagram")
	req.Bank = c.FormValue("bank")
	req.AccountNo = c.FormValue("accountNo")
	req.BankBranch = c.FormValue("bankBranch")

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	file, err := c.FormFile("profilePicture")
	// if error mean cannot get file just ignore.
	// because field is not provide mean not change.
	var signedURL string
	if err == nil {
		signedURL, err = h.uploadProfileFile(c.UserContext(), file, profileFolder(userId))
		if err != nil {
			return errors.Wrap(err, "File upload failed")
		}
	}

	// var oldPictureURL string
	updatedUser, err := h.updateUserDB(userId, req, signedURL, nil)
	if err != nil {
		if signedURL != "" {
			err = h.store.Storage.DeleteFile(c.UserContext(), profileFolder(userId)+path.Base(signedURL))
			if err != nil {
				return errors.Wrap(err, "Fail to delete the picture")
			}
		}
		return errors.Wrap(err, "Error updating user profile")
	}

	// if oldPictureURL != "" && oldPictureURL != signedURL {
	// 	err = h.store.Storage.DeleteFile(c.UserContext(), profileFolder(userId)+path.Base(oldPictureURL))
	// 	if err != nil {
	// 		return errors.Wrap(err, "Fail to delete old picture")
	// 	}
	// }

	response := dto.UserResponse{
		ID:                updatedUser.ID,
		Name:              updatedUser.Name,
		Email:             updatedUser.Email,
		PhoneNumber:       updatedUser.PhoneNumber,
		ProfilePictureURL: updatedUser.ProfilePictureURL,
		Role:              updatedUser.Role.String(),
		Facebook:          updatedUser.Facebook,
		Instagram:         updatedUser.Instagram,
		Bank:              updatedUser.Bank,
		AccountNo:         updatedUser.AccountNo,
		BankBranch:        updatedUser.BankBranch,
	}

	return c.JSON(dto.HttpResponse[dto.UserResponse]{
		Result: response,
	})
}

func (h *Handler) uploadProfileFile(c context.Context, file *multipart.FileHeader, folder string) (string, error) {
	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	var signedURL string
	if signedURL, err = h.store.Storage.UploadFile(c, folder+file.Filename, contentType, src, true); err != nil {
		return "", errors.Wrap(err, "failed to upload file")
	}

	return signedURL, nil
}

func (h *Handler) updateUserDB(userID uint, req *dto.UserUpdateRequest, signedURL string, oldPictureURL *string) (*model.User, error) {
	var user model.User

	err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, "id = ?", userID).Error; err != nil {
			return errors.Wrap(err, "User not found")
		}

		// Assign old picture URL if needed
		if oldPictureURL != nil {
			*oldPictureURL = user.ProfilePictureURL
		}

		updateField := func(field *string, newValue string) {
			if newValue != "" {
				*field = newValue
			}
		}

		updateField(&user.Name, req.Name)
		updateField(&user.PhoneNumber, req.PhoneNumber)
		updateField(&user.ProfilePictureURL, signedURL)
		updateField(&user.Facebook, req.Facebook)
		updateField(&user.Instagram, req.Instagram)
		updateField(&user.Bank, req.Bank)
		updateField(&user.AccountNo, req.AccountNo)
		updateField(&user.BankBranch, req.BankBranch)

		if err := tx.Save(&user).Error; err != nil {
			return errors.Wrap(err, "Failed to update user")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func profileFolder(userID uint) string {
	return "profile/" + strconv.FormatUint(uint64(userID), 10) + "/"
}
