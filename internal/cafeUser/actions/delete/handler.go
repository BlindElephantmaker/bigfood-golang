package cafeUserDelete

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
)

type Message struct {
	CafeUserId string `json:"cafe-user-id" binding:"required" example:"uuid"`
}

func (h *Handler) Run(m *Message) error {
	cafeUserId, err := helpers.UuidParse(m.CafeUserId)
	if err != nil {
		return err
	}

	return h.CafeUserRepository.Delete(cafeUserId)
}

type Handler struct {
	CafeUserRepository cafeUser.Repository
}

func New(cafeUsers cafeUser.Repository) *Handler {
	return &Handler{cafeUsers}
}
