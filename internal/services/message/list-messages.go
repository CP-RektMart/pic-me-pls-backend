package message

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

// @Summary      list all chats
// @Description  list all chats for each individual user
// @Tags         chat
// @Router       /api/v1/chats [GET]
// @Security	 ApiKeyAuth
// @Success      200    {object}  dto.HttpListResponse[dto.ChatResponse]
// @Failure      401    {object}  dto.HttpError
// @Failure      404    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleListMessages(c *fiber.Ctx) error {
	userID, err := h.authentication.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed getting userID from context")
	}

	var messages []model.Message

	if err := h.store.DB.
		Order("created_at").
		Joins("Receiver").
		Joins("Sender").
		Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Find(&messages).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.NotFound("Messages not found", err)
		}
		return errors.Wrap(err, "failed getting sended messages")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.ChatResponse]{
		Result: h.toChatResponse(userID, messages),
	})
}

func (h *Handler) toChatResponse(userID uint, messages []model.Message) []dto.ChatResponse {
	chats := make(map[uint]*dto.ChatResponse)
	var talker model.User

	for _, msg := range messages {
		if msg.ReceiverID == userID {
			talker = msg.Sender
		} else {
			talker = msg.Receiver
		}

		chat, ok := chats[talker.ID]
		if !ok {
			chats[talker.ID] = &dto.ChatResponse{
				User:     dto.ToPublicUserResponse(talker),
				Messages: make([]dto.RealTimeMessageResponse, 0),
			}
			chat = chats[talker.ID]
		}

		chat.Messages = append(chat.Messages, dto.ToRealTimeMessageResponse(msg))
	}

	return lo.MapToSlice(chats, func(_ uint, chat *dto.ChatResponse) dto.ChatResponse {
		return *chat
	})
}
