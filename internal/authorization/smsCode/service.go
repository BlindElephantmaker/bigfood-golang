package smsCode

import "bigfood/internal/user"

type Service interface {
	Send(text string, phone user.Phone) error
}
