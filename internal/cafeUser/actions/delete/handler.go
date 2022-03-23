package cafeUserDelete

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
)

type Message struct {
	CafeUserId helpers.Uuid `json:"cafe-user-id" binding:"required" example:"uuid"`
}

func (h *Handler) Run(m *Message) error {
	return h.CafeUserRepository.Delete(m.CafeUserId)
}

type Handler struct {
	CafeUserRepository cafeUser.Repository
}

func New(cafeUsers cafeUser.Repository) *Handler {
	return &Handler{cafeUsers}
}
