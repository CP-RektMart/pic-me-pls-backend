package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary      Get All Users
// @Description  Retrieve a paginated list of users.
// @Tags         admin
// @Router       /api/v1/admin/users [GET]
// @Security     ApiKeyAuth
// @Param        page      query    int  false  "Page number (default: 1)"
// @Param        pageSize  query    int  false  "Page size (default: 5, max: 20)"
// @Param        name      query    string  false  "Filter by user's name (case-insensitive)"
// @Success      200       {object} dto.PaginationResponse[dto.PublicUserResponse]
// @Failure      400       {object} dto.HttpError
// @Failure      404       {object} dto.HttpError
// @Failure      500       {object} dto.HttpError
func (h *Handler) HandleGetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	var params dto.GetUsersRequest

	if err := c.QueryParser(&params); err != nil {
		return apperror.BadRequest("Invalid query parameters", err)
	}

	if err := h.validate.Struct(params); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	query := h.store.DB.Model(&model.User{}).Where("name ILIKE ?", "%"+params.Name+"%")

	page, pageSize, offset := dto.GetPaginationData(params.PaginationRequest, 1, 5)

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return errors.Wrap(err, "failed to count users")
	}
	totalPage := (int(totalCount) + pageSize - 1) / pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("users not found", err)
		}
		return errors.Wrap(err, "failed to fetch users")
	}

	return c.Status(fiber.StatusOK).JSON(dto.PaginationResponse[dto.PublicUserResponse]{
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      dto.ToPublicUserResponses(users),
	})

}
