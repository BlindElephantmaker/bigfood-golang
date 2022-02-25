package refreshToken

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/user/userToken"
)

type Handler struct {
	tokenRepository    userToken.Repository
	cafeUserRepository cafeUser.Repository
}

func New(tokens userToken.Repository, cafeUsers cafeUser.Repository) *Handler {
	return &Handler{
		tokenRepository:    tokens,
		cafeUserRepository: cafeUsers,
	}
}

func (h *Handler) Run(message *Message) (*userToken.UserToken, error) {
	refresh, err := userToken.ParseRefresh(message.Token)
	if err != nil {
		return nil, err
	}
	oldToken, err := h.tokenRepository.Get(refresh)
	if err != nil {
		return nil, err
	}

	permissions, err := h.cafeUserRepository.GetUserPermissions(oldToken.UserId)
	if err != nil {
		return nil, err
	}

	newToken, err := userToken.NewUserToken(permissions)
	if err != nil {
		return nil, err
	}

	err = h.tokenRepository.Refresh(newToken, oldToken)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}
