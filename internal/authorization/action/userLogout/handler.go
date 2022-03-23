package userLogout

import (
	"bigfood/internal/helpers"
	"bigfood/internal/user/userToken"
)

type Message struct {
	UserId       helpers.Uuid           `swaggerignore:"true"`
	RefreshToken userToken.RefreshToken `json:"token" binding:"required" example:"uuid"` // todo: check uuid validation
}

func (h *Handler) Run(message *Message) error {
	return h.userTokenRepository.Delete(message.RefreshToken, message.UserId)
}

type Handler struct {
	userTokenRepository userToken.Repository
}

func New(tokens userToken.Repository) *Handler {
	return &Handler{
		userTokenRepository: tokens,
	}
}
