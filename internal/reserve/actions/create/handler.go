package reserveCreate

import (
	"bigfood/internal/helpers"
	"bigfood/internal/reserve"
	"bigfood/internal/reserve/actions"
)

func (h *Handler) Run(m *reserveAction.Message) (*reserve.Reserve, error) {
	newReserve, err := h.reserveActionHelper.ParseMessage(reserve.NewId(), m)
	if err != nil {
		return nil, err
	}

	if err := h.reserveRepository.Add(newReserve, helpers.NowTime()); err != nil {
		return nil, err
	}

	return newReserve, nil
}

type Handler struct {
	reserveRepository   reserve.Repository
	reserveActionHelper *reserveAction.Helper
}

func New(
	reserveRepository reserve.Repository,
	reserveActionHelper *reserveAction.Helper,
) *Handler {
	return &Handler{
		reserveRepository:   reserveRepository,
		reserveActionHelper: reserveActionHelper,
	}
}
