package server

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

func (s *Server) handleLogin(c *fiber.Ctx) error {
	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := s.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	user, err := s.validateIDToken(c.Context(), req.IDToken)
	if err != nil {
		return errors.Wrap(err, "failed to validate id token")
	}
	user.Role = req.Role

	user, err = s.getOrCreateUser(user)
	if err != nil {
		return err
	}


	return c.Status(fiber.StatusOK).JSON(dto.BaseUserDTO{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role,
	})
}

func (s *Server) validateIDToken(c context.Context, idToken string) (*model.User, error) {
	payload, err := idtoken.Validate(c, idToken, s.config.GoogleClientID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate id token")
	}

	name, ok := payload.Claims["name"].(string)
	if !ok {
		return nil, errors.New("name claim not found in id token")
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return nil, errors.New("email claim not found in id token")
	}

	picture, ok := payload.Claims["picture"].(string)
	if !ok {
		return nil, errors.New("picture claim not found in id token")
	}

	return &model.User{
		Name:              name,
		Email:             email,
		ProfilePictureURL: picture,
	}, nil
}

func (s *Server) getOrCreateUser(user *model.User) (*model.User, error) {
	// Check if user exists
	// if user exists, return user
	// if user does not exist, create user with name and email then return user

	// this is my mock user
	return &model.User{
		Model:             gorm.Model{ID: 123},
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
		PhoneNumber:       "0123456789",
		Role:              user.Role,
	}, nil
}

func (s *Server) generateJWTToken() error {
	return nil
}
