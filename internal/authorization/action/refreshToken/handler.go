package refreshToken

import (
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
)

type Message struct {
	RefreshToken userToken.RefreshToken `json:"refresh-token" binding:"required" example:"UUID"`
}

type Response struct {
	UserId  user.Id                `json:"user-id" example:"UUID"`
	Access  userToken.AccessToken  `json:"access-token"`
	Refresh userToken.RefreshToken `json:"refresh-token" example:"UUID"`
}

func (h *Handler) Run(m *Message) (*Response, error) {
	refresh, err := h.userTokenRepository.Get(m.RefreshToken)
	if err != nil {
		return nil, err
	}

	access, usrToken, err := h.userTokenService.CreateTokens(refresh.UserId)
	if err != nil {
		return nil, err
	}

	return &Response{
		UserId:  usrToken.UserId,
		Refresh: usrToken.Refresh,
		Access:  access,
	}, nil
}

type Handler struct {
	userTokenRepository userToken.Repository
	userTokenService    *userToken.Service
}

func New(
	userTokens userToken.Repository,
	userTokenService *userToken.Service,
) *Handler {
	return &Handler{
		userTokenRepository: userTokens,
		userTokenService:    userTokenService,
	}
}
