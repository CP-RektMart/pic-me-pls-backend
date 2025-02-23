package user

import (
	"context"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

func (h *Handler) RegisterUpdateMe(api huma.API, middlewares huma.Middlewares) {
	huma.Register(api, huma.Operation{
		OperationID: "update-me",
		Method:      http.MethodPatch,
		Path:        "/api/v1/me",
		Summary:     "Update my profile",
		Description: "Update my profile",
		Tags:        []string{"user"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleUpdateMe)
}

func (h *Handler) HandleUpdateMe(ctx context.Context, req *dto.HumaFormData[dto.UserUpdateRequest]) (*dto.HumaHttpResponse[dto.UserResponse], error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request", err)
	}

	file, ok := req.RawBody.Form.File["profilePicture"]
	if !ok {
		return nil, huma.Error400BadRequest("invalid request", errors.New("profilePicture is required"))
	}

	var signedURL string
	if len(file) > 0 {
		signedURL, err = h.uploadProfileFile(ctx, file[0], profileFolder(userId))
		if err != nil {
			return nil, errors.Wrap(err, "File upload failed")
		}
	}

	data := req.RawBody.Data()
	// var oldPictureURL string
	updatedUser, err := h.updateUserDB(userId, data, signedURL, nil)
	if err != nil {
		if signedURL != "" {
			err = h.store.Storage.DeleteFile(ctx, profileFolder(userId)+path.Base(signedURL))
			if err != nil {
				return nil, errors.Wrap(err, "Fail to delete the picture")
			}
		}
		return nil, errors.Wrap(err, "Error updating user profile")
	}

	// if oldPictureURL != "" && oldPictureURL != signedURL {
	// 	err = h.store.Storage.DeleteFile(c.UserContext(), profileFolder(userId)+path.Base(oldPictureURL))
	// 	if err != nil {
	// 		return errors.Wrap(err, "Fail to delete old picture")
	// 	}
	// }

	response := dto.UserResponse{
		ID:                updatedUser.ID,
		Name:              updatedUser.Name,
		Email:             updatedUser.Email,
		PhoneNumber:       updatedUser.PhoneNumber,
		ProfilePictureURL: updatedUser.ProfilePictureURL,
		Role:              updatedUser.Role.String(),
		Facebook:          updatedUser.Facebook,
		Instagram:         updatedUser.Instagram,
		Bank:              updatedUser.Bank,
		AccountNo:         updatedUser.AccountNo,
		BankBranch:        updatedUser.BankBranch,
	}

	return &dto.HumaHttpResponse[dto.UserResponse]{
		Body: dto.HttpResponse[dto.UserResponse]{
			Result: response,
		},
	}, nil
}

func (h *Handler) uploadProfileFile(c context.Context, file *multipart.FileHeader, folder string) (string, error) {
	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer src.Close()
	var signedURL string
	if signedURL, err = h.store.Storage.UploadFile(c, folder+file.Filename, contentType, src, true); err != nil {
		return "", errors.Wrap(err, "failed to upload file")
	}

	return signedURL, nil
}

func (h *Handler) updateUserDB(userID uint, req *dto.UserUpdateRequest, signedURL string, oldPictureURL *string) (*model.User, error) {
	var user model.User

	err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&user, "id = ?", userID).Error; err != nil {
			return errors.Wrap(err, "User not found")
		}

		// Assign old picture URL if needed
		if oldPictureURL != nil {
			*oldPictureURL = user.ProfilePictureURL
		}

		updateField := func(field *string, newValue string) {
			if newValue != "" {
				*field = newValue
			}
		}

		updateField(&user.Name, req.Name)
		updateField(&user.PhoneNumber, req.PhoneNumber)
		updateField(&user.ProfilePictureURL, signedURL)
		updateField(&user.Facebook, req.Facebook)
		updateField(&user.Instagram, req.Instagram)
		updateField(&user.Bank, req.Bank)
		updateField(&user.AccountNo, req.AccountNo)
		updateField(&user.BankBranch, req.BankBranch)

		if err := tx.Save(&user).Error; err != nil {
			return errors.Wrap(err, "Failed to update user")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func profileFolder(userID uint) string {
	return "profile/" + strconv.FormatUint(uint64(userID), 10) + "/"
}
