package getList

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table"
)

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}

func (h *Handler) Run(message *Message) ([]*table.Table, error) {
	cafeId, err := helpers.UuidParse(message.CafeId)
	if err != nil {
		return nil, err
	}

	tables, err := h.TableRepository.GetByCafe(cafeId)
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		tables = []*table.Table{}
	}

	return tables, nil
}
