package cafeUserCreate

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/actions"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"errors"
)

type Message struct {
	CafeId  string   `json:"cafe-id" binding:"required" example:"uuid"`
	Phone   string   `json:"phone" binding:"required" example:"User phone"`
	Comment *string  `json:"comment"`
	Roles   []string `json:"roles" binding:"required" example:"owner,admin,hostess"`
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
	phone, err := user.NewPhone(m.Phone)
	if err != nil {
		return "", "", "", nil, err
	}
	cafeId, err := helpers.UuidParse(m.CafeId)
	if err != nil {
		return "", "", "", nil, err
	}

	var comment cafeUser.Comment
	if m.Comment != nil {
		comment, err = cafeUser.ParseComment(*m.Comment)
		if err != nil {
			return "", "", "", nil, err
		}
	} else {
		comment = cafeUser.NewComment()
	}

	return cafeId, phone, comment, roles, nil
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
