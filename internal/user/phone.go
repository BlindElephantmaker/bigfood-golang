package user

import (
	"errors"
	"regexp"
)

type Phone struct {
	value string
}

const pattern = `^\+\d{11}$` // todo: phone number is not always 11 digits

func NewPhone(phone string) (*Phone, error) {
	ok, _ := regexp.MatchString(pattern, phone)
	if !ok {
		return nil, errors.New("phone number is invalid")
	}

	return &Phone{phone}, nil
}

func (p *Phone) String() string {
	return p.value
}
