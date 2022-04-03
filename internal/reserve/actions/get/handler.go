package reserveGet

import "bigfood/internal/reserve"

func (h *Handler) Run(reserveId reserve.Id) (*reserve.Reserve, error) {
	return h.reserveRepository.Get(reserveId)
}

type Handler struct {
	reserveRepository reserve.Repository
}

func New(reserves reserve.Repository) *Handler {
	return &Handler{
		reserveRepository: reserves,
	}
}
