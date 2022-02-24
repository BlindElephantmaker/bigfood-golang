package sendSmsCode

import (
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/user"
	"errors"
	"time"
)

const (
	maxRetryCount = 3
	ttl           = time.Minute * 30
)

var ErrorRetryCountExceeded = errors.New("retry count of sms code requests exceeded")

type Handler struct {
	smsCodeService    smsCode.Service
	smsCodeRepository smsCode.Repository
}

func New(service smsCode.Service, repository smsCode.Repository) *Handler {
	return &Handler{
		smsCodeService:    service,
		smsCodeRepository: repository,
	}
}

func (h *Handler) Run(m Message) error {
	phone, err := user.NewPhone(m.Phone)
	if err != nil {
		return err
	}

	count, err := h.smsCodeRepository.Count(phone)
	if err != nil {
		return err
	}
	if count >= maxRetryCount {
		return ErrorRetryCountExceeded
	}

	code := smsCode.New()
	err = h.smsCodeRepository.Add(code, phone, ttl)
	if err != nil {
		return err
	}

	err = h.smsCodeService.Send(string(code), phone)
	if err != nil {
		_ = h.smsCodeRepository.DeleteLast(phone)
		return err
	}

	return nil
}
