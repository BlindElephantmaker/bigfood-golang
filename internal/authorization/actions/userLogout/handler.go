package userLogout

import (
	"bigfood/internal/user/userToken"
)

type Handler struct {
	userTokenRepository userToken.Repository
}

func New(tokens userToken.Repository) *Handler {
	return &Handler{
		userTokenRepository: tokens,
	}
}

func (h *Handler) Run(message *Message) error {
	token, err := userToken.ParseRefresh(message.Token)
	if err != nil {
		return err
	}

	return h.userTokenRepository.Delete(token, message.UserId)
}
