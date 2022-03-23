package refreshToken

import (
	"bigfood/internal/user/userToken"
)

type Message struct {
	RefreshToken userToken.RefreshToken `json:"token" binding:"required" example:"UUID"`
}

func (h *Handler) Run(m *Message) (*userToken.UserToken, error) {
	oldToken, err := h.tokenRepository.Get(m.RefreshToken)
	if err != nil {
		return nil, err
	}

	newToken, err := userToken.NewUserToken(oldToken.UserId)
	if err != nil {
		return nil, err
	}

	err = h.tokenRepository.Refresh(newToken, oldToken)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

type Handler struct {
	tokenRepository userToken.Repository
}

func New(tokens userToken.Repository) *Handler {
	return &Handler{
		tokenRepository: tokens,
	}
}
