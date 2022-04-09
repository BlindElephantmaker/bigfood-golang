package reserveHistoryByTable

import (
	"bigfood/internal/reserve"
	"bigfood/internal/table"
)

type Response struct {
	Reserves []*reserve.Reserve `json:"reserves"`
}

func (h *Handler) Run(tableId table.Id, limit int, offset int) (*Response, error) {
	reserves, err := h.reserveRepository.GetHistoryByTableId(tableId, limit, offset)
	if err != nil {
		return nil, err
	}

	return &Response{reserves}, nil
}

type Handler struct {
	reserveRepository reserve.Repository
}

func New(reserveRepository reserve.Repository) *Handler {
	return &Handler{reserveRepository}
}
