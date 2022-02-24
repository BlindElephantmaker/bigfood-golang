package user

import (
	"errors"
	"regexp"
)

type Phone string

const pattern = `^\+\d{11}$` // todo: phone number is not always 11 digits

var ErrorPhoneNumberIsInvalid = errors.New("phone number is invalid")

func NewPhone(phone string) (Phone, error) {
	ok, _ := regexp.MatchString(pattern, phone)
	if !ok {
		return "", ErrorPhoneNumberIsInvalid
	}

	return Phone(phone), nil
}
