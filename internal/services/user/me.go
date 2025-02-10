package user

import (
	"fmt"
	"strconv"

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
// @Router /api/v1/user/profile [post]
func (h *Handler) HandlerUpdateProfile(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: User ID not found")
	}

	req := new(dto.UpdateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
	}

	// Upload file (if present)
	folder := strconv.FormatUint(uint64(userID), 10) + "/profile/"
	signedURL, fileUploaded, err := h.UploadProfileFile(folder, c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("File upload failed: %v", err))
	}

	if fileUploaded {
		req.ProfilePictureURL = signedURL
	}

	updatedUser, _, err := h.updateUserDB(userID, req)
	if err != nil {
		// TODO: remove picture [Currently bug: failed to delete file: body must be object]
		// if fileUploaded {
		// 	h.store.Storage.DeleteFile(c.Context(), folder+signedURL)
		// }
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error updating user profile: %v", err))
	}

	// TODO: remove old picture [Currently bug: failed to delete file: body must be object]
	// if oldPictureURL != "" && oldPictureURL != req.ProfilePictureURL {
	// 	fileName := path.Base(oldPictureURL)
	// 	err := h.store.Storage.DeleteFile(c.Context(), folder+fileName)
	// 	fmt.Println(err)
	// }

	return c.JSON(dto.HttpResponse{
		Result: updatedUser,
	})
}

func (h *Handler) UploadProfileFile(folder string, c *fiber.Ctx) (string, bool, error) {
	file, err := c.FormFile("profile")
	if err != nil {
		return "", false, nil
	}

	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", false, errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	var signedURL string
	if signedURL, err = h.store.Storage.UploadFile(c.Context(), folder+file.Filename, contentType, src, true); err != nil {
		return "", false, errors.Wrap(err, "failed to upload file")
	}

	return signedURL, true, nil
}

func (h *Handler) updateUserDB(userID uint, req *dto.UpdateUserRequest) (*model.User, string, error) {
	var user model.User

	if err := h.store.DB.First(&user, "id = ?", userID).Error; err != nil {
		return nil, "", err
	}

	oldPictureURL := user.ProfilePictureURL

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
		return nil, "", err
	}

	return &user, oldPictureURL, nil
}
