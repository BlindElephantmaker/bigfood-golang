package smsCode

import (
	"errors"
	"math/rand"
	"regexp"
)

type Code struct {
	value string
}

const (
	length  = 4
	symbols = "0123456789"
	pattern = `^\d{4}$`
)

var ErrorSmsCodeIsInvalid = errors.New("sms code is invalid")

func New() *Code {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = symbols[rand.Intn(len(symbols))]
	}

	return &Code{string(buf)}
}

func Parse(code string) (*Code, error) {
	ok, _ := regexp.MatchString(pattern, code)
	if !ok {
		return nil, ErrorSmsCodeIsInvalid
	}

	return &Code{code}, nil
}

func (c *Code) String() string {
	return c.value
}

func (c *Code) Compare(another *Code) bool {
	return c.String() != another.String()
}
