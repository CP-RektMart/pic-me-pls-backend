package server

import (
	"context"
	"fmt"
	"time"
	"fmt"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/utils/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/utils/jwt"
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

	OAuthUser, err := s.validateIDToken(c.Context(), req.IDToken)
	if err != nil {
		return errors.Wrap(err, "failed to validate id token")
	}
	OAuthUser.Role = req.Role

	var user *model.User
	var token *dto.TokenResponse

	if err := s.db.DB.Transaction(func(tx *gorm.DB) error {
		user, err = s.getOrCreateUser(tx, OAuthUser)
		if err != nil {
			return err
		}

		token, err = s.generateJWTToken(c.UserContext(), user.ID, user.Role)
		if err != nil {
			return err
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
		Role:              user.Role,
	}

	result := dto.LoginResponse{
		TokenResponse: *token,
		User:          userDTO,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{Result: result})
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

<<<<<<< HEAD
func (s *Server) getOrCreateUser(tx *gorm.DB, user *model.User) (*model.User, error) {
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
func (s *Server) generateJWTToken(c context.Context, ID uint, role string) (*dto.TokenResponse, error) {
	type token struct {
		secret   string
		duration int
		token    string
	}

	tokens := []token{
		{secret: s.config.JWT.AccessSecret, duration: s.config.JWT.AccessDuration},
		{secret: s.config.JWT.RefreshSecret, duration: s.config.JWT.RefreshDuration},
	}

	for i, token := range tokens {
		tokenString, err := jwt.GenerateJWT(ID, role, token.secret, token.duration)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate jwt token")
		}
		if err := s.storeJWTToken(c, ID, tokenString, token.duration); err != nil {
			return nil, errors.Wrap(err, "failed to store jwt token")
		}
		tokens[i].token = tokenString
	}

	return &dto.TokenResponse{
		AcessToken:   tokens[0].token,
		RefreshToken: tokens[1].token,
		Exp:          s.config.JWT.AccessDuration,
	}, nil
}

func (s *Server) storeJWTToken(c context.Context, ID uint, token string, duration int) error {
	key := s.newTokenCacheKey(token, ID)
	ttl := time.Second * time.Duration(duration)

	if err := s.db.Cache.Set(c, key, 1, ttl).Err(); err != nil {
		return errors.Wrap(err, "failed to store jwt token")
	}

	return nil
}

func (s *Server) newTokenCacheKey(token string, ID uint) string {
	return fmt.Sprintf("%s:%d", token, ID)
}
=======
func (s *Server) getOrCreateUser(user *model.User) (*model.User, error) {
	var existingUser model.User
	err := s.db.DB.Where("email = ?", user.Email).First(&existingUser).Error
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
	if err := s.db.DB.Create(&newUser).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}

	return &newUser, nil
}
func (s *Server) generateJWTToken(c context.Context, ID uint, role string) (*dto.TokenResponse, error) {
	type token struct {
		secret   string
		duration int
		token    string
	}

	tokens := []token{
		{secret: s.config.JwtAccessSecret, duration: s.config.JwtAccessDuration},
		{secret: s.config.JWTRefreshSecret, duration: s.config.JwtRefreshDuration},
	}

	for i, token := range tokens {
		tokenString, err := jwt.GenerateJWT(ID, role, token.secret, token.duration)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate jwt token")
		}
		if err := s.storeJWTToken(c, ID, tokenString, token.duration); err != nil {
			return nil, errors.Wrap(err, "failed to store jwt token")
		}
		tokens[i].token = tokenString
	}

	return &dto.TokenResponse{
		AcessToken:   tokens[0].token,
		RefreshToken: tokens[1].token,
		Exp:          s.config.JwtAccessDuration,
	}, nil
}

func (s *Server) storeJWTToken(c context.Context, ID uint, token string, duration int) error {
	key := s.newTokenCacheKey(token, ID)
	ttl := time.Second * time.Duration(duration)

	if err := s.db.Cache.Set(c, key, 1, ttl).Err(); err != nil {
		return errors.Wrap(err, "failed to store jwt token")
	}

	return nil
}
>>>>>>> d7bc73a (chore: set up)
