package getList

import (
	"bigfood/internal/cafe"
	"bigfood/internal/table"
)

type Message struct {
	CafeId cafe.Id `json:"cafe-id" binding:"required" example:"uuid"`
}

func (h *Handler) Run(m *Message) ([]*table.Table, error) {
	tables, err := h.TableRepository.GetByCafe(m.CafeId)
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		tables = []*table.Table{}
	}

	return tables, nil
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
