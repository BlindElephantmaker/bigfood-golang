package cafeUserList

import (
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/actions"
	"bigfood/internal/user"
)

type Message struct {
	CafeId cafe.Id `json:"cafe-id" binding:"required" example:"uuid"`
}

type Response struct {
	CafeUsers []*actions.Response `json:"cafe-users"`
}

func (h *Handler) Run(m *Message) (*Response, error) {
	responseEmpty := &Response{[]*actions.Response{}}

	cafeUsers, err := h.CafeUserRepository.GetByCafe(m.CafeId)
	if err != nil {
		return responseEmpty, err
	}

	response := []*actions.Response{}
	for _, cafeUsr := range cafeUsers {
		usr, err := h.UserRepository.Get(cafeUsr.UserId)
		if err != nil {
			return responseEmpty, err
		}

		roles, err := h.CafeUserRepository.GetUserRoles(cafeUsr.Id)
		if err != nil {
			return responseEmpty, err
		}

		response = append(response, &actions.Response{
			CafeUser: cafeUsr,
			Roles:    *roles,
			Name:     usr.Name,
		})
	}

	return &Response{response}, nil
}

type Handler struct {
	UserRepository     user.Repository
	CafeUserRepository cafeUser.Repository
}

func New(users user.Repository, cafeUsers cafeUser.Repository) *Handler {
	return &Handler{
		UserRepository:     users,
		CafeUserRepository: cafeUsers,
	}
}
