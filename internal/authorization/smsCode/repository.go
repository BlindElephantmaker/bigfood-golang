package smsCode

import (
	"bigfood/internal/user"
	"time"
)

type Repository interface {
	Add(code Code, phone user.Phone, ttl time.Duration) error
	Get(user.Phone) (Code, error)
	Count(user.Phone) (int, error)
	DeleteLast(user.Phone) error
}
