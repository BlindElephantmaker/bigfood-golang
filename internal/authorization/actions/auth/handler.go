package auth

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/authorization/userToken"
	"bigfood/internal/user"
	"errors"
	"github.com/google/uuid"
)

type Handler struct {
	smsCodeRepository smsCode.Repository
	userRepository    user.Repository
	tokenRepository   userToken.Repository
}

var ErrorSmsCodeNotConfirmed = errors.New("sms code not confirmed")

func New(smsCodeRepository smsCode.Repository, users user.Repository, tokens userToken.Repository) *Handler {
	return &Handler{
		smsCodeRepository: smsCodeRepository,
		userRepository:    users,
		tokenRepository:   tokens,
	}
}

func (handler *Handler) Run(message *Message) (*Response, error) {
	phone, err := user.NewPhone(message.Phone)
	if err != nil {
		return nil, err
	}
	code, err := smsCode.Parse(message.SmsCode)
	if err != nil {
		return nil, err
	}

	err = handler.validateSmsCode(phone, code)
	if err != nil {
		return nil, err
	}

	u, err := handler.getUser(phone)
	if err != nil {
		return nil, err
	}

	token, err := handler.getToken(u.Id)
	if err != nil {
		return nil, err
	}

	return &Response{
		UserToken: token,
		IsNew:     u.IsNew(),
	}, nil
}

func (handler *Handler) validateSmsCode(phone *user.Phone, code *smsCode.Code) error {
	confirmCode, err := handler.smsCodeRepository.Get(phone)
	if err != nil {
		return err
	}
	if !confirmCode.Compare(code) {
		return ErrorSmsCodeNotConfirmed
	}

	return nil
}

func (handler *Handler) getUser(phone *user.Phone) (*user.User, error) {
	isExist, err := handler.userRepository.IsExistByPhone(phone)
	if err != nil {
		return nil, err
	}

	if isExist {
		return handler.userRepository.GetByPhone(phone)
	}

	emptyName, _ := user.NewName("")
	id := uuid.New()
	newUser := user.New(&id, emptyName, phone)

	return newUser, handler.userRepository.Add(newUser)
}

func (handler *Handler) getToken(id *uuid.UUID) (*userToken.UserToken, error) {
	token, err := userToken.New(id)
	if err != nil {
		return nil, err
	}

	err = handler.tokenRepository.Add(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
