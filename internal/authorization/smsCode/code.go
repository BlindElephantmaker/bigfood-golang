package smsCode

import (
	"encoding/json"
	"errors"
	"math/rand"
	"regexp"
)

type SmsCode string

const (
	length  = 4
	symbols = "0123456789"
	pattern = `^\d{4}$`
)

var ErrorSmsCodeIsInvalid = errors.New("sms code is invalid")

func New() SmsCode {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = symbols[rand.Intn(len(symbols))]
	}

	return SmsCode(buf)
}

func (sc *SmsCode) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	smsCode, err := parse(value)
	if err != nil {
		return err
	}

	*sc = smsCode
	return nil
}

func (sc *SmsCode) Compare(another SmsCode) bool {
	return *sc == another
}

func parse(code string) (SmsCode, error) {
	ok, _ := regexp.MatchString(pattern, code)
	if !ok {
		return "", ErrorSmsCodeIsInvalid
	}

	return SmsCode(code), nil
}
