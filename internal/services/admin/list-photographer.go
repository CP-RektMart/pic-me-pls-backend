package admin

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			list photographers
// @Tags			admin
// @Router			/api/v1/admin/photographers [GET]
// @Security		ApiKeyAuth
// @Param 	 	 name	   query    string    	false  "Filter by photographer's name"
// @Param        page      query    int    	    false  "Page number"
// @Param        pageSize  query    int    		false  "Page size"
// @Success			200 {object}	dto.HttpResponse[dto.PaginationResponse[dto.PhotographerResponse]]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListPhotographer(c *fiber.Ctx) error {
	return nil
}
