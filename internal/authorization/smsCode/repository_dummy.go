package smsCode

import (
	"bigfood/internal/helpers"
	"time"
)

// todo: code it or replace to redis

type RepositoryDummy struct {
	memory map[helpers.Phone]*dummy
}

type dummy struct {
	code  SmsCode
	count int
}

func NewRepositoryDummy() *RepositoryDummy {
	return &RepositoryDummy{make(map[helpers.Phone]*dummy)}
}

func (r *RepositoryDummy) Add(code SmsCode, phone helpers.Phone, ttl time.Duration) error {
	d := r.memory[phone]
	if d != nil {
		d.code = code
		d.count++
	} else {
		r.memory[phone] = &dummy{
			code:  code,
			count: 1,
		}
	}

	return nil
}

func (r *RepositoryDummy) Get(phone helpers.Phone) (SmsCode, error) {
	return "1234", nil
}

func (r *RepositoryDummy) Count(phone helpers.Phone) (int, error) {
	d := r.memory[phone]
	if d != nil {
		return d.count, nil
	}

	return 0, nil
}

func (r *RepositoryDummy) DeleteLast(helpers.Phone) error {
	return nil
}
