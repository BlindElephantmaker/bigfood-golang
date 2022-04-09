package reserveListByTable

import (
	"bigfood/internal/reserve"
	"bigfood/internal/table"
)

type Response struct {
	Actual  []*reserve.Reserve `json:"actual"`
	Deleted []*reserve.Reserve `json:"deleted"`
}

func (h *Handler) Run(tableId table.Id) (*Response, error) {
	actual, err := h.reserveRepository.GetActualByTableId(tableId)
	if err != nil {
		return nil, err
	}
	deleted, err := h.reserveRepository.GetDeletedByTableId(tableId)
	if err != nil {
		return nil, err
	}

	return &Response{
		Actual:  actual,
		Deleted: deleted,
	}, nil
}

type Handler struct {
	reserveRepository reserve.Repository
}

func New(reserveRepository reserve.Repository) *Handler {
	return &Handler{reserveRepository}
}
