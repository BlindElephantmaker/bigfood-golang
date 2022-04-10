package tableGetListAvailable

import (
	"bigfood/internal/cafe"
	"bigfood/internal/table"
	"time"
)

type Response struct {
	Tables []*table.Table `json:"tables"`
}

func (h *Handler) Run(cafeId cafe.Id, from, until time.Time) (*Response, error) {
	tables, err := h.repository.GetListAvailable(cafeId, from, until)
	if err != nil {
		return nil, err
	}

	return &Response{tables}, nil
}

type Handler struct {
	repository Repository
}

func New(repository Repository) *Handler {
	return &Handler{repository}
}
