package createCafe

import (
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
)

type Handler struct {
	cafeRepository cafe.Repository
}

func New(cafes cafe.Repository) *Handler {
	return &Handler{
		cafeRepository: cafes,
	}
}

func (h *Handler) Run(userId helpers.Uuid) (helpers.Uuid, error) {
	newCafe := cafe.New()
	newCafeUser := cafeUser.NewCafeUser(newCafe.Id, userId, cafeUser.NewComment())
	err := h.cafeRepository.Add(newCafe, newCafeUser, cafeUser.Roles{
		cafeUser.Owner,
		cafeUser.Admin,
		cafeUser.Hostess,
	})
	if err != nil {
		return "", err
	}

	return newCafe.Id, nil
}
