package auth

import (
	"context"
	"strings"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/utils/convert"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/idtoken"
	"gorm.io/gorm"
)

// @Summary			Register
// @Description			Register
// @Tags			auth
// @Router			/api/v1/auth/Register [POST]
// @Param 			RequestBody 	body 	dto.RegisterRequest 	true 	"request request"
// @Success			200	{object}	dto.HttpResponse{result=dto.RegisterResponse}
// @Failure			400	{object}	dto.HttpResponse
// @Failure			500	{object}	dto.HttpResponse
func (h *Handler) HandleRegister(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(dto.RegisterRequest)
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
	var token *model.Token

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		user, err = h.createUser(tx, OAuthUser)
		if err != nil {
			return errors.Wrap(err, "failed to get or create user")
		}

		token, err = h.jwtService.GenerateAndStoreTokenPair(ctx, user)
		if err != nil {
			return errors.Wrap(err, "failed to generate token pair")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create user and token")
	}

	result := dto.RegisterResponse{
		TokenResponse: dto.TokenResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			Exp:          token.Exp,
		},
		User: convert.ToUserResponse(*user),
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

func (h *Handler) createUser(tx *gorm.DB, user *model.User) (*model.User, error) {
	if err := tx.Save(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, apperror.BadRequest("this account already register", err)
		}

		return nil, errors.Wrap(err, "failed to create user")
	}

	if user.Role == model.UserRolePhotographer {
		if err := tx.Create(&model.Photographer{UserID: user.ID}).Error; err != nil {
			return nil, errors.Wrap(err, "failed to create photographer")
		}
	}

	return user, nil
}
