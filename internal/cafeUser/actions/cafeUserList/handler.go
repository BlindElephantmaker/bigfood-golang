package cafeUserList

import (
	"bigfood/internal/cafe"
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/actions"
	"bigfood/internal/user"
)

type Response struct {
	CafeUsers []*actions.Response `json:"cafe-users"`
}

func (h *Handler) Run(cafeId cafe.Id) (*Response, error) {
	responseEmpty := &Response{[]*actions.Response{}}

	cafeUsers, err := h.CafeUserRepository.GetByCafe(cafeId)
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
			UserName: usr.Name,
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
