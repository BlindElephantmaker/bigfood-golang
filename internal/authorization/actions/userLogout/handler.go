package userLogout

import (
	userToken2 "bigfood/internal/user/userToken"
)

type Handler struct {
	userTokenRepository userToken2.Repository
}

func New(tokens userToken2.Repository) *Handler {
	return &Handler{
		userTokenRepository: tokens,
	}
}

func (h *Handler) Run(message *Message) error {
	token, err := userToken2.ParseRefresh(message.Token)
	if err != nil {
		return err
	}

	return h.userTokenRepository.Delete(token, message.UserId)
}
