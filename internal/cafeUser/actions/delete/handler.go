package cafeUserDelete

import (
	"bigfood/internal/cafeUser"
)

type Message struct {
	CafeUserId cafeUser.Id `json:"cafe-user-id" binding:"required" example:"uuid"`
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
