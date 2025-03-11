package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary      	List Photographer's packages
// @Description  	List Photographer's packages
// @Tags         	packages
// @Router       	/api/v1/photographer/packages [GET]
// @Security	 	ApiKeyAuth
// @Success      	200    {object}  dto.HttpListResponse[dto.SmallPackageResponse]
// @Failure      	400    {object}  dto.HttpError
// @Failure      	500    {object}  dto.HttpError
func (h *Handler) HandlerListPhotographerPackages(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	packages, err := h.getPackagesFromUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed fetch packages")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.SmallPackageResponse]{
		Result: dto.ToSmallPackageResponses(packages),
	})
}

func (h *Handler) getPackagesFromUserID(ID uint) ([]model.Package, error) {
	var photographer model.Photographer
	if err := h.store.DB.Where("user_id = ?", ID).First(&photographer).Error; err != nil {
		return nil, errors.Wrap(err, "failed fetch photographer")
	}

	packages := make([]model.Package, 0)
	if err := h.store.DB.Where("photographer_id = ?", photographer.ID).Find(&packages).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.BadRequest("packages not found", err)
		}
		return nil, errors.Wrap(err, "failed fetch packages")
	}

	return packages, nil
}
