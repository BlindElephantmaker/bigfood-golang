package smsCode

import (
	"bigfood/internal/helpers"
	"time"
)

type Repository interface {
	Add(code SmsCode, phone helpers.Phone, ttl time.Duration) error
	Get(helpers.Phone) (SmsCode, error)
	Count(helpers.Phone) (int, error)
	DeleteLast(helpers.Phone) error
}
