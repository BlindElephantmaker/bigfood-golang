package refreshToken

import "bigfood/internal/authorization/userToken"

type Handler struct {
	tokenRepository userToken.Repository
}

func New(tokens userToken.Repository) *Handler {
	return &Handler{
		tokenRepository: tokens,
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

	newToken, err := userToken.New(oldToken.UserId)
	if err != nil {
		return nil, err
	}

	err = h.tokenRepository.Refresh(newToken, oldToken)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}
