package user

import (
	"bigfood/internal/helpers"
	"encoding/json"
	"regexp"
)

type Phone string

const pattern = `^\+\d{11}$` // todo: phone number is not always 11 digits

var errorPhoneNumberIsInvalid = helpers.NewErrorBadRequest("phone number is invalid")

func (p *Phone) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	phone, err := parsePhone(value)
	if err != nil {
		return err
	}

	*p = phone
	return nil
}

func parsePhone(phone string) (Phone, error) {
	ok, _ := regexp.MatchString(pattern, phone)
	if !ok {
		return "", errorPhoneNumberIsInvalid
	}

	return Phone(phone), nil
}
