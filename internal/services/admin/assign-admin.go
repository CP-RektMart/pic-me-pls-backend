package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			assign user to admin
// @Tags			admin
// @Router			/api/v1/admin/users/{userID}/role [PATCH]
// @Security		ApiKeyAuth
// @Param			userID		path		int	true	"user id"
// @Param			detail		body		dto.AssignAdminRequest true "detail"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleAssignAdmin(c *fiber.Ctx) error {
	var req dto.AssignAdminRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	if err := h.assignAdmin(req.UserID, *req.Admin); err != nil {
		return errors.Wrap(err, "failed assign admin")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) assignAdmin(userID uint, admin bool) error {
	var user model.User
	if err := h.store.DB.First(&user, userID).Error; err != nil {
		return apperror.NotFound("user not found", err)
	}

	if admin {
		if user.Role == model.UserRoleAdmin {
			return apperror.BadRequest("user already be an admin", nil)
		}

		user.Role = model.UserRoleAdmin
	} else {
		if user.Role != model.UserRoleAdmin {
			return apperror.BadRequest("user already not an admin", nil)
		}

		err := h.store.DB.First(&model.Photographer{}, userID).Error
		isPhotographer := err == nil

		if isPhotographer {
			user.Role = model.UserRolePhotographer
		} else {
			user.Role = model.UserRoleCustomer
		}
	}

	return h.store.DB.Updates(&user).Error
}
