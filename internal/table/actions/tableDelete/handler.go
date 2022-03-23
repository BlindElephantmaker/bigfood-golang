package tableDelete

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
)

type Message struct {
	TableId helpers.Uuid `json:"table-id" binding:"required" example:"uuid"`
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
