package tableList

import (
	"bigfood/internal/cafe"
	"bigfood/internal/table"
)

type Response struct {
	Tables []*table.Table `json:"tables"`
}

func (h *Handler) Run(cafeId cafe.Id) (*Response, error) {
	tables, err := h.TableRepository.GetByCafe(cafeId)
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		tables = []*table.Table{}
	}

	return &Response{tables}, nil
}

type Handler struct {
	TableRepository table.Repository
}

func New(tables table.Repository) *Handler {
	return &Handler{tables}
}
