package auth

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/cafe/cafeUser"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"errors"
)

type Handler struct {
	smsCodeRepository  smsCode.Repository
	userRepository     user.Repository
	tokenRepository    userToken.Repository
	cafeUserRepository cafeUser.Repository
}

var ErrorSmsCodeNotConfirmed = errors.New("sms code not confirmed")

func New(
	smsCodeRepository smsCode.Repository,
	users user.Repository,
	tokens userToken.Repository,
	cafeUsers cafeUser.Repository,
) *Handler {
	return &Handler{
		smsCodeRepository:  smsCodeRepository,
		userRepository:     users,
		tokenRepository:    tokens,
		cafeUserRepository: cafeUsers,
	}
}

func (h *Handler) Run(message *Message) (*Response, error) {
	phone, err := user.NewPhone(message.Phone)
	if err != nil {
		return nil, err
	}
	code, err := smsCode.Parse(message.SmsCode)
	if err != nil {
		return nil, err
	}

	err = h.validateSmsCode(phone, code)
	if err != nil {
		return nil, err
	}

	u, err := h.getUser(phone)
	if err != nil {
		return nil, err
	}

	token, err := h.createToken(u.Id)
	if err != nil {
		return nil, err
	}

	return &Response{
		UserToken: token,
		IsNew:     u.IsNew(),
	}, nil
}

func (h *Handler) validateSmsCode(phone user.Phone, code smsCode.Code) error {
	confirmCode, err := h.smsCodeRepository.Get(phone)
	if err != nil {
		return err
	}
	if !confirmCode.Compare(code) {
		return ErrorSmsCodeNotConfirmed
	}

	return nil
}

func (h *Handler) getUser(phone user.Phone) (*user.User, error) {
	isExist, err := h.userRepository.IsExistByPhone(phone)
	if err != nil {
		return nil, err
	}

	if isExist {
		return h.userRepository.GetByPhone(phone)
	}

	newUser := user.New(phone)

	return newUser, h.userRepository.Add(newUser)
}

func (h *Handler) createToken(id helpers.Uuid) (*userToken.UserToken, error) {
	permissions, err := h.cafeUserRepository.GetUserPermissions(id)
	if err != nil {
		return nil, err
	}

	token, err := userToken.NewUserToken(permissions)
	if err != nil {
		return nil, err
	}

	err = h.tokenRepository.Add(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
