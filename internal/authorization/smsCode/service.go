package smsCode

import "bigfood/internal/helpers"

type Service interface {
	Send(text string, phone helpers.Phone) error
}
