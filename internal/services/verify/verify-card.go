package verify

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

var (
	ErrAlreadyVerified = errors.New("ALREADY_VERIFIED")
)

func (h *Handler) RegisterVerifyCard(api huma.API, middlewares ...func(ctx huma.Context, next func(huma.Context))) {
	huma.Register(api, huma.Operation{
		OperationID: "verify-card",
		Method:      http.MethodPost,
		Path:        "/api/v1/photographer/verify",
		Summary:     "Verify citizen card",
		Description: "Verify citizen card",
		Tags:        []string{"verify"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleVerifyCard)
}

func (h *Handler) HandleVerifyCard(ctx context.Context, req *dto.HumaFormData[dto.CitizenCardRequest]) (*dto.HumaHttpResponse[dto.CitizenCardResponse], error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request", err)
	}

	file, ok := req.RawBody.Form.File["cardPicture"]
	if !ok {
		return nil, huma.Error400BadRequest("invalid request", errors.New("cardPicture is required"))
	}

	var signedURL string
	if len(file) > 0 {
		signedURL, err = h.uploadCardFile(ctx, file[0], citizenCardFolder(userId))
		if err != nil {
			return nil, errors.Wrap(err, "File upload failed")
		}
	}

	data := req.RawBody.Data()
	user, err := h.createCitizenCard(data, signedURL, userId)
	if err != nil {
		return nil, errors.Wrap(err, "Fail to create citizen card")
	}

	response := dto.CitizenCardResponse{
		CitizenID:  user.CitizenID,
		LaserID:    user.LaserID,
		Picture:    user.Picture,
		ExpireDate: user.ExpireDate,
	}

	return &dto.HumaHttpResponse[dto.CitizenCardResponse]{
		Body: dto.HttpResponse[dto.CitizenCardResponse]{
			Result: response,
		},
	}, nil
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
			return errors.WithStack(ErrAlreadyVerified)
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
		if errors.Is(err, ErrAlreadyVerified) {
			return nil, huma.Error400BadRequest("already verified", err)
		}
		return nil, errors.Wrap(err, "failed to create citizen card")
	}

	return &citizenCard, nil
}
