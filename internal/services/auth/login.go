package auth

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

// handlerLogin godoc
// @summary login with external service provider
// @description provider can be GOOGLE,  role can be CUSTOMER, PHOTOGRAPHER, ADMIN
// @tags auth
// @security Bearer
// @id login
// @accept json
// @produce json
// @param LoginRequest body dto.LoginRequest true "login request"
// @response 200 {object} dto.LoginResponse "OK"
// @response 400 {object} dto.HttpResponse "Bad Request"
// @response 500 {object} dto.HttpResponse "Internal Server Error"
// @Router /api/v1/auth/login [POST]
func (h *Handler) HandleLogin(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	OAuthUser, err := h.validateIDToken(ctx, req.IDToken)
	if err != nil {
		return errors.Wrap(err, "failed to validate id token")
	}
	OAuthUser.Role = model.UserRole(req.Role)

	var user *model.User
	var token *dto.TokenResponse

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		user, err = h.getOrCreateUser(tx, OAuthUser)
		if err != nil {
			return errors.Wrap(err, "failed to get or create user")
		}

		token, err = h.jwtService.GenerateAndStoreTokenPair(ctx, user)
		if err != nil {
			return errors.Wrap(err, "failed to generate token pair")
		}

		return nil
	}); err != nil {
		return err
	}

	userDTO := dto.BaseUserDTO{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role.String(),
		PhoneNumber:       user.PhoneNumber,
		Facebook:          "",
		Instagram:         "",
		AccountNo:         "",
		Bank:              "",
		BankBranch:        "",
	}

	result := dto.LoginResponse{
		TokenResponse: *token,
		User:          userDTO,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: result,
	})
}

func (h *Handler) validateIDToken(c context.Context, idToken string) (*model.User, error) {
	payload, err := idtoken.Validate(c, idToken, h.googleClientID)
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

func (h *Handler) getOrCreateUser(tx *gorm.DB, user *model.User) (*model.User, error) {
	var existingUser model.User
	err := tx.Where("email = ?", user.Email).First(&existingUser).Error
	if err == nil {
		return &existingUser, nil
	}

	if err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "failed to get user")
	}

	newUser := model.User{
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role,
	}
	if err := tx.Create(&newUser).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &newUser, nil
}
