package cafeUserEdit

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/actions"
	"bigfood/internal/user"
	"errors"
	"fmt"
)

type Message struct {
	CafeUserId cafeUser.Id       `json:"cafe-user-id" binding:"required" example:"uuid"`
	Comment    *cafeUser.Comment `json:"comment"`
	Roles      *[]string         `json:"roles"` // todo: bad array swagger and parse collection
}

var ErrorOwnerRoleCouldNotBeSet = errors.New(fmt.Sprintf("%s role could not be set", cafeUser.Owner))

func (h *Handler) Run(m *Message) (*actions.Response, error) {
	roles, err := parseMessage(m)

	cafeUsr, err := h.CafeUserRepository.Get(m.CafeUserId)
	if err != nil {
		return nil, err
	}
	usr, err := h.UserRepository.Get(cafeUsr.UserId)
	if err != nil {
		return nil, err
	}

	if m.Comment != nil {
		cafeUsr.Comment = *m.Comment
	}

	if roles == nil {
		roles, err = h.CafeUserRepository.GetUserRoles(m.CafeUserId)
		if err != nil {
			return nil, err
		}
	}

	if err := h.CafeUserRepository.Update(cafeUsr, *roles); err != nil {
		return nil, err
	}

	return &actions.Response{
		CafeUser: cafeUsr,
		Roles:    *roles,
		Name:     usr.Name,
	}, nil
}

func parseMessage(m *Message) (*cafeUser.Roles, error) {
	var roles cafeUser.Roles
	if m.Roles != nil {
		roles, err := cafeUser.ParseRoles(*m.Roles)
		for _, role := range roles {
			if role == cafeUser.Owner {
				return nil, ErrorOwnerRoleCouldNotBeSet
			}
		}
		if err != nil {
			return nil, err
		}
	}

	return &roles, nil
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
