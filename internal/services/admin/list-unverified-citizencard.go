package admin

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			list unverified citizencard
// @Tags			admin
// @Router			/api/v1/admin/citizenCards/unverify [GET]
// @Security		ApiKeyAuth
// @Param        page      query    int    	    false  "Page number"
// @Param        pageSize  query    int    		false  "Page size"
// @Success			200 {object}	dto.HttpResponse[dto.PaginationResponse[dto.CitizenCardResponse]]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListCitizenCard(c *fiber.Ctx) error {
	return nil
}
