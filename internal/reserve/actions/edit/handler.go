package reserveEdit

import (
	"bigfood/internal/reserve"
	"bigfood/internal/reserve/actions"
)

type Message struct {
	ReserveId reserve.Id `json:"reserve-id" binding:"required" example:"uuid"`
	*reserveAction.Message
}

func (h *Handler) Run(m *Message) (*reserve.Reserve, error) {
	reserv, err := h.reserveActionHelper.ParseMessage(m.ReserveId, m.Message)
	if err != nil {
		return nil, err
	}

	if err := h.reserveRepository.Update(reserv); err != nil {
		return nil, err
	}

	return reserv, nil
}

type Handler struct {
	reserveRepository   reserve.Repository
	reserveActionHelper *reserveAction.Helper
}

func New(reserveRepository reserve.Repository, reserveActionHelper *reserveAction.Helper) *Handler {
	return &Handler{
		reserveRepository:   reserveRepository,
		reserveActionHelper: reserveActionHelper,
	}
}
