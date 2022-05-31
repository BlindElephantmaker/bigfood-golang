package tableGet

import "bigfood/internal/table"

type Handler struct {
	repository table.Repository
}

func (h *Handler) Run(tableId table.Id) (*table.Table, error) {
	return h.repository.Get(tableId)
}

func New(repository table.Repository) *Handler {
	return &Handler{repository: repository}
}
