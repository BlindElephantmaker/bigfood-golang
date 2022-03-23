package tableDeleteAll

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
)

type Message struct {
	CafeId helpers.Uuid `json:"cafe-id" binding:"required" example:"uuid"`
}

func (h *Handler) Run(message *Message) error {
	return h.TableRepository.DeleteAll(message.CafeId)
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
