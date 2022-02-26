package cafeUserEdit

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/actions"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"errors"
	"fmt"
)

type Message struct {
	CafeUserId string    `json:"cafe-user-id" binding:"required" example:"uuid"`
	Comment    *string   `json:"comment"`
	Roles      *[]string `json:"roles"` // todo: bad array swagger
}

var ErrorOwnerRoleCouldNotBeSet = errors.New(fmt.Sprintf("%s role could not be set", cafeUser.Owner))

func (h *Handler) Run(m *Message) (*actions.Response, error) {
	cafeUserId, comment, roles, err := parseMessage(m)

	cafeUsr, err := h.CafeUserRepository.Get(cafeUserId)
	if err != nil {
		return nil, err
	}
	usr, err := h.UserRepository.Get(cafeUsr.UserId)
	if err != nil {
		return nil, err
	}

	if comment != nil {
		cafeUsr.Comment = *comment
	}

	if roles == nil {
		roles, err = h.CafeUserRepository.GetUserRoles(cafeUserId)
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

func parseMessage(m *Message) (helpers.Uuid, *cafeUser.Comment, *cafeUser.Roles, error) {
	cafeUserId, err := helpers.UuidParse(m.CafeUserId)
	if err != nil {
		return "", nil, nil, err
	}

	var comment *cafeUser.Comment
	if m.Comment != nil {
		c, err := cafeUser.ParseComment(*m.Comment)
		if err != nil {
			return "", nil, nil, err
		}
		comment = &c
	}

	var roles cafeUser.Roles
	if m.Roles != nil {
		roles, err = cafeUser.ParseRoles(*m.Roles)
		for _, role := range roles {
			if role == cafeUser.Owner {
				return "", nil, nil, ErrorOwnerRoleCouldNotBeSet
			}
		}
		if err != nil {
			return "", nil, nil, err
		}
	}

	return cafeUserId, comment, &roles, nil
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
