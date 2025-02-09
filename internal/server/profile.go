package server

import (
	"fmt"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (s *Server) profileRoute(route fiber.Router) {
	route.Get("/profile", s.authMiddleware.Auth, s.getProfile)
	// POST route
}

func (s *Server) getProfile(c *fiber.Ctx) error {
	userID, err := s.authMiddleware.GetUserIDFromContext(c.UserContext())

	// TODO: GetRoleFromContext please!!
	role := model.UserRoleCustomer

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Failed to get user details from context")
	}

	var profile interface{}

	if role == model.UserRolePhotographer {
		profile, err = getPhotographerFromDB(s.db.DB, userID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error retrieving photographer profile: %v", err))
		}
	} else if role == model.UserRoleCustomer {
		profile, err = getCustomerFromDB(s.db.DB, userID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Error retrieving customer profile: %v", err))
		}
	} else {
		return fiber.NewError(fiber.StatusForbidden, "Invalid role")
	}

	return c.JSON(profile)
}

func getPhotographerFromDB(db *gorm.DB, id uint) (model.Photographer, error) {
	var photographer model.Photographer
	err := db.Where("id = ?", id).First(&photographer).Error
	if err != nil {
		return model.Photographer{}, err
	}
	return photographer, nil
}

func getCustomerFromDB(db *gorm.DB, id uint) (model.User, error) {
	var customer model.User
	err := db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		return model.User{}, err
	}
	return customer, nil
}
