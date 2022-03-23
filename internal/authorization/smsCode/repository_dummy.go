package smsCode

import (
	"bigfood/internal/user"
	"time"
)

// todo: code it or replace to redis

type RepositoryDummy struct {
	memory map[user.Phone]*dummy
}

type dummy struct {
	code  SmsCode
	count int
}

func NewRepositoryDummy() *RepositoryDummy {
	return &RepositoryDummy{make(map[user.Phone]*dummy)}
}

func (r *RepositoryDummy) Add(code SmsCode, phone user.Phone, ttl time.Duration) error {
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

func (r *RepositoryDummy) Get(phone user.Phone) (SmsCode, error) {
	return "1234", nil
}

func (r *RepositoryDummy) Count(phone user.Phone) (int, error) {
	d := r.memory[phone]
	if d != nil {
		return d.count, nil
	}

	return 0, nil
}

func (r *RepositoryDummy) DeleteLast(user.Phone) error {
	return nil
}
