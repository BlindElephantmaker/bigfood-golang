package reserveDelete

import (
	"bigfood/internal/reserve"
)

type Message struct {
	ReserveId reserve.Id `json:"reserve-id" binding:"required" example:"uuid"`
}

func (h *Handler) Run(m *Message) error {
	return h.reserveRepository.Delete(m.ReserveId)
}

type Handler struct {
	reserveRepository reserve.Repository
}

func New(reserves reserve.Repository) *Handler {
	return &Handler{reserves}
}
