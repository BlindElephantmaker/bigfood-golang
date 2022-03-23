package smsCode

import (
	"bigfood/internal/user"
	"time"
)

type Repository interface {
	Add(code SmsCode, phone user.Phone, ttl time.Duration) error
	Get(user.Phone) (SmsCode, error)
	Count(user.Phone) (int, error)
	DeleteLast(user.Phone) error
}
