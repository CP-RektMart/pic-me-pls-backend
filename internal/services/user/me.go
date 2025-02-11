package user

import (
	"context"
	"mime/multipart"
	"path"
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/pkg/errors"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// handlerGetMe godoc
// @summary Get user profile
// @description Retrieves the authenticated user's profile.
// @tags user
// @security Bearer
// @id get-me
// @accept json
// @produce json
// @success 200 {object} dto.BaseUserDTO "OK"
// @failure 400 {object} dto.HttpResponse "Bad Request"
// @failure 500 {object} dto.HttpResponse "Internal Server Error"
// @Router /api/v1/me [GET]
func (h *Handler) HandleGetMe(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var user model.User
	result := h.store.DB.First(&user, userId)
	if result.Error != nil {
		return apperror.Internal("failed to get user", nil)
	}

	userDTO := dto.BaseUserDTO{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role.String(),
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{Result: userDTO})
}

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
func (h *Handler) HandleUpdateMe(c *fiber.Ctx) error {
	userID, _ := h.authMiddleware.GetUserIDFromContext(c.UserContext())

	req := new(dto.BaseUserDTO)

	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	var signedURL string = ""
	file, err := c.FormFile("profile")
	if err != nil {
		signedURL, err = h.uploadProfileFile(c.UserContext(), file, profileFolder(userID))
		if err != nil {
			return errors.Wrap(err, "File upload failed")
		}
		req.ProfilePictureURL = signedURL
	}

	var oldPictureURL string
	updatedUser, err := h.updateUserDB(userID, req, &oldPictureURL)
	if err != nil {
		if signedURL != "" {
			err = h.store.Storage.DeleteFile(c.UserContext(), profileFolder(userID)+path.Base(signedURL))
			if err != nil {
				return errors.Wrap(err, "Fail to delete the picture")
			}
		}
		return errors.Wrap(err, "Error updating user profile")
	}

	if oldPictureURL != "" && oldPictureURL != req.ProfilePictureURL {
		err = h.store.Storage.DeleteFile(c.UserContext(), profileFolder(userID)+path.Base(oldPictureURL))
		if err != nil {
			return errors.Wrap(err, "Fail to delete old picture")
		}
	}

	return c.JSON(dto.HttpResponse{
		Result: updatedUser,
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
