package user

import (
	"path"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/pkg/errors"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// HandlerUpdateProfile updates the user's profile information
// @Summary Update user profile
// @Description Updates the user's profile information including email, phone number, social media links, and bank account details
// @Tags user
// @Accept json
// @Produce json
// @Param updateUserRequest body dto.UpdateUserRequest true "User profile update data"
// @Success 200 {object} dto.HttpResponse "Profile updated successfully"
// @Failure 400 {object} dto.HttpResponse "Bad request, invalid input data"
// @Failure 404 {object} dto.HttpResponse "User not found"
// @Failure 500 {object} dto.HttpResponse "Internal server error"
// @Router /api/v1/user/profile [patch]
func (h *Handler) HandleUpdateProfile(c *fiber.Ctx) error {
	userID, _ := h.authMiddleware.GetUserIDFromContext(c.UserContext())

	req := new(dto.BaseUserDTO)

	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	signedURL, fileUploaded, err := h.uploadProfileFile(c, profileFolder(userID))
	if err != nil {
		return errors.Wrap(err, "File upload failed")
	}

	if fileUploaded {
		req.ProfilePictureURL = signedURL
	}

	var oldPictureURL string
	updatedUser, err := h.updateUserDB(userID, req, &oldPictureURL)
	if err != nil {
		if fileUploaded {
			h.store.Storage.DeleteFile(c.UserContext(), profileFolder(userID)+path.Base(signedURL))
		}
		if err == gorm.ErrRecordNotFound {
			return errors.Wrap(err, "User not found")
		}
		return errors.Wrap(err, "Error updating user profile")
	}

	if oldPictureURL != "" && oldPictureURL != req.ProfilePictureURL {
		h.store.Storage.DeleteFile(c.UserContext(), profileFolder(userID)+path.Base(oldPictureURL))
	}

	return c.JSON(dto.HttpResponse{
		Result: updatedUser,
	})
}

func (h *Handler) uploadProfileFile(c *fiber.Ctx, folder string) (string, bool, error) {
	file, err := c.FormFile("profile")
	if err != nil {
		// This allows because if the field is not provided mean they dont change the pictrue
		return "", false, nil
	}

	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", false, errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	var signedURL string
	if signedURL, err = h.store.Storage.UploadFile(c.UserContext(), folder+file.Filename, contentType, src, true); err != nil {
		return "", false, errors.Wrap(err, "failed to upload file")
	}

	return signedURL, true, nil
}

func (h *Handler) updateUserDB(userID uint, req *dto.BaseUserDTO, oldPictureURL *string) (*model.User, error) {
	var user model.User

	if err := h.store.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, errors.Wrap(err, "User not found")
	}
	if oldPictureURL != nil {
		oldPictureURL = &user.ProfilePictureURL
	}

	updateField := func(field *string, newValue string) {
		if newValue != "" {
			*field = newValue
		}
	}

	updateField(&user.Name, req.Name)
	updateField(&user.PhoneNumber, req.PhoneNumber)
	updateField(&user.ProfilePictureURL, req.ProfilePictureURL)
	updateField(&user.Facebook, req.Facebook)
	updateField(&user.Instagram, req.Instagram)
	updateField(&user.Bank, req.Bank)
	updateField(&user.AccountNo, req.AccountNo)
	updateField(&user.BankBranch, req.BankBranch)

	if err := h.store.DB.Save(&user).Error; err != nil {
		return nil, errors.Wrap(err, "File to update user")
	}

	return &user, nil
}

func profileFolder(userID uint) string {
	return "/profile/" + strconv.FormatUint(uint64(userID), 10)
}
