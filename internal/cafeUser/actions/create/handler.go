package cafeUserCreate

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/cafeUser/role"
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

type Response struct {
	*cafeUser.User
	Name user.Name `json:"name"`
}

var ErrorCafeUserAlreadyExist = errors.New("cafe user already exist")

func (h *Handler) Run(m *Message) (*Response, error) {
	cafeId, phone, comment, roles, err := parseMessage(m)
	if err != nil {
		return nil, err
	}

	u, err := h.UserService.GetOrNewUser(phone)
	if err != nil {
		return nil, err
	}

	cu, err := h.CafeUserRepository.Get(cafeId, u.Id)
	if err == cafeUser.ErrorNoResult {
		cu = cafeUser.NewCafeUser(cafeId, u.Id, comment, roles)
		err = h.CafeUserRepository.Add(cu)
	} else if err != nil {
		return nil, err
	} else {
		if !cu.IsDeleted() {
			return nil, ErrorCafeUserAlreadyExist
		}
		cu.DeletedAt = nil
		cu.Comment = comment
		cu.Roles = roles
		err = h.CafeUserRepository.Update(cu)
	}

	return &Response{
		User: cu,
		Name: u.Name,
	}, err
}

func parseMessage(m *Message) (helpers.Uuid, user.Phone, cafeUser.Comment, role.Roles, error) {
	roles, err := role.ParseRoles(m.Roles)
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
