package user

import (
	"fmt"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandlerUpdateProfile(c *fiber.Ctx) error {
	// TODO: get payload from jwt (middleware) or something bababa -->
	req := new(dto.UpdateUserRequest)
	req.Email = "customer@example.com"
	req.PhoneNumber = "0810824581"
	req.ProfilePictureURL = "///"
	req.Facebook = "testtest"
	req.Instagram = "test"
	req.BankBranch = "test"
	req.Bank = "kabnk"
	req.AccountNo = "123123123123"
	// -->

	updatedUser, err := h.updateUserDB(h.store.DB, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error updating user profile: %v", err))
	}

	return c.JSON(dto.HttpResponse{
		Result: updatedUser,
	})
}

func (h *Handler) updateUserDB(db *gorm.DB, req *dto.UpdateUserRequest) (*model.User, error) {
	var user model.User

	if err := db.First(&user, "email = ?", req.Email).Error; err != nil {
		return nil, err
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

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
