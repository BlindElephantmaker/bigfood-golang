package tableDeleteAll

import (
	"bigfood/internal/cafe"
	"bigfood/internal/table"
)

type Message struct {
	CafeId cafe.Id `json:"cafe-id" binding:"required" example:"uuid"`
}

func (h *Handler) Run(m *Message) error {
	return h.TableRepository.DeleteAll(m.CafeId)
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
