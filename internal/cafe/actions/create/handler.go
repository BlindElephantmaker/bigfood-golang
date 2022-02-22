package createCafe

import (
	"bigfood/internal/cafe"
	"bigfood/internal/cafe/cafeUser"
	"bigfood/internal/helpers"
	"github.com/google/uuid"
)

type Handler struct {
	cafeRepository cafe.Repository
}

func New(cafes cafe.Repository) *Handler {
	return &Handler{
		cafeRepository: cafes,
	}
}

func (h *Handler) Run(userId *uuid.UUID) (*uuid.UUID, error) {
	newCafe := cafe.New()
	newCafeUser := cafeUser.NewCafeUser(newCafe.Id, userId)
	now := helpers.Now()
	err := h.cafeRepository.Add(newCafe, newCafeUser, &now)
	if err != nil {
		return nil, err
	}

	return newCafe.Id, nil
}
