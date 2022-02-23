package getList

import "bigfood/internal/table"

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}

func (h *Handler) Run(message *Message) ([]*table.Table, error) {
	return h.TableRepository.Get(message.CafeId)
}
