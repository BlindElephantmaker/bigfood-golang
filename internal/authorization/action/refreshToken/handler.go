package refreshToken

import (
	"bigfood/internal/cafeUser/permissions"
	"bigfood/internal/user/userToken"
)

type Message struct {
	RefreshToken userToken.RefreshToken `json:"token" binding:"required" example:"UUID"` // todo: check uuid validation
}

func (h *Handler) Run(m *Message) (*userToken.UserToken, error) {
	oldToken, err := h.tokenRepository.Get(m.RefreshToken)
	if err != nil {
		return nil, err
	}

	userPermissions, err := h.permissionRepository.GetPermissions(oldToken.UserId)
	if err != nil {
		return nil, err
	}

	newToken, err := userToken.NewUserToken(userPermissions)
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
	tokenRepository      userToken.Repository
	permissionRepository permissions.Repository
}

func New(tokens userToken.Repository, permissions permissions.Repository) *Handler {
	return &Handler{
		tokenRepository:      tokens,
		permissionRepository: permissions,
	}
}
