package photographer

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Verify Citizen Card
// @Description		Verify Photographer Citizen Card
// @Tags			photographer
// @Router			/api/v1/photographer/verify [POST]
// @Security		ApiKeyAuth
// @Accept			multipart/form-data
// @Param 			cardPicture 	formData 	file		false	"Card picture (optional)"
// @Param 			citizenId 		formData 	string		true	"Citizen ID"
// @Param 			laserId 		formData 	string		true	"Laser ID"
// @Param 			expireDate 		formData 	string		true	"Expire Date"
// @Success			200	{object}	dto.HttpResponse[dto.CitizenCardResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleVerifyCard(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.CitizenCardRequest)
	req.CitizenID = c.FormValue("citizenId")
	req.LaserID = c.FormValue("laserId")
	req.ExpireDate, err = time.Parse(time.RFC3339, c.FormValue("expireDate"))
	if err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	file, err := c.FormFile("cardPicture")
	if err != nil {
		return apperror.BadRequest("card Picture is require", errors.Errorf("Field Missing"))
	}

	var signedURL string
	signedURL, err = h.uploadCardFile(c.UserContext(), file, citizenCardFolder(userId))
	if err != nil {
		return errors.Wrap(err, "File upload failed")
	}

	user, err := h.createCitizenCard(req, signedURL, userId)
	if err != nil {
		return errors.Wrap(err, "Fail to create citizen card")
	}

	response := dto.CitizenCardResponse{
		CitizenID:  user.CitizenID,
		LaserID:    user.LaserID,
		Picture:    user.Picture,
		ExpireDate: user.ExpireDate,
	}

	return c.JSON(dto.HttpResponse[dto.CitizenCardResponse]{
		Result: response,
	})
}

func (h *Handler) uploadCardFile(c context.Context, file *multipart.FileHeader, folder string) (string, error) {
	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	signedURL, err := h.store.Storage.UploadFile(c, folder+file.Filename, contentType, src, true)
	if err != nil {

		return "", errors.Wrap(err, "failed to upload file")
	}

	return signedURL, nil
}

func (h *Handler) createCitizenCard(req *dto.CitizenCardRequest, signedURL string, userId uint) (*model.CitizenCard, error) {
	var citizenCard model.CitizenCard

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userId).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		// Check if the photographer already has a CitizenCard
		if photographer.CitizenCardID != nil {
			return apperror.BadRequest("photographer already has a citizen card", errors.Errorf("Already verified"))
		}

		// Create the CitizenCard using the request data
		citizenCard.CitizenID = req.CitizenID
		citizenCard.LaserID = req.LaserID
		citizenCard.Picture = signedURL
		citizenCard.ExpireDate = req.ExpireDate

		// Insert the CitizenCard into the database
		if err := tx.Create(&citizenCard).Error; err != nil {
			return errors.Wrap(err, "Error creating citizen card")
		}

		// Update the photographer's CitizenCardID with the new CitizenCard ID
		photographer.CitizenCardID = &citizenCard.ID
		if err := tx.Save(&photographer).Error; err != nil {
			return errors.Wrap(err, "Error updating photographer with citizen card")
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "failed to create citizen card")
	}

	return &citizenCard, nil
}
