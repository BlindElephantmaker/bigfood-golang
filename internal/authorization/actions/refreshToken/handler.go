package refreshToken

import (
	"bigfood/internal/cafeUser/permissions"
	"bigfood/internal/user/userToken"
)

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

func (h *Handler) Run(message *Message) (*userToken.UserToken, error) {
	refresh, err := userToken.ParseRefresh(message.Token)
	if err != nil {
		return nil, err
	}
	oldToken, err := h.tokenRepository.Get(refresh)
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
