package tableDelete

import (
	"bigfood/internal/table"
)

type Message struct {
	TableId table.Id `json:"table-id" binding:"required" example:"uuid"`
}

func (h *Handler) Run(m *Message) error {
	return h.TableRepository.Delete(m.TableId)
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
