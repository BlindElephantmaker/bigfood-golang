package cafeUserCreate

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/actions"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"errors"
)

type Message struct {
	CafeId  helpers.Uuid      `json:"cafe-id" binding:"required" example:"uuid"`
	Phone   user.Phone        `json:"phone" binding:"required" example:"User phone"`
	Comment *cafeUser.Comment `json:"comment"`
	Roles   []string          `json:"roles" binding:"required" example:"owner,admin,hostess"` // todo: collection
}

var ErrorCafeUserAlreadyExist = errors.New("cafe user already exist")

func (h *Handler) Run(m *Message) (*actions.Response, error) {
	cafeId, phone, comment, roles, err := parseMessage(m)
	if err != nil {
		return nil, err
	}

	usr, err := h.UserService.GetOrNewUser(phone)
	if err != nil {
		return nil, err
	}

	cafeUsr, err := h.CafeUserRepository.GetByCafeAndUserIds(cafeId, usr.Id)
	if err == cafeUser.ErrorNoResult {
		cafeUsr = cafeUser.NewCafeUser(cafeId, usr.Id, comment)
		err = h.CafeUserRepository.Add(cafeUsr, roles)
	} else if err != nil {
		return nil, err
	} else {
		if !cafeUsr.IsDeleted() {
			return nil, ErrorCafeUserAlreadyExist
		}
		cafeUsr.DeletedAt = nil
		cafeUsr.Comment = comment
		err = h.CafeUserRepository.Update(cafeUsr, roles)
	}

	return &actions.Response{
		CafeUser: cafeUsr,
		Roles:    roles,
		Name:     usr.Name,
	}, err
}

func parseMessage(m *Message) (helpers.Uuid, user.Phone, cafeUser.Comment, cafeUser.Roles, error) {
	roles, err := cafeUser.ParseRoles(m.Roles)
	if err != nil {
		return "", "", "", nil, err
	}

	var comment cafeUser.Comment
	if m.Comment != nil { // todo: maybe move it to UnmarshalJSON?
		comment = *m.Comment
	} else {
		comment = cafeUser.NewComment()
	}

	return m.CafeId, m.Phone, comment, roles, nil // todo
}

type Handler struct {
	CafeUserRepository cafeUser.Repository
	UserService        *user.Service
}

func New(cafeUsers cafeUser.Repository, userService *user.Service) *Handler {
	return &Handler{
		CafeUserRepository: cafeUsers,
		UserService:        userService,
	}
}
