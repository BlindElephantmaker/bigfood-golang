package smsCode

import (
	"bigfood/internal/user"
	"time"
)

// todo: code it or replace to redis

type RepositoryDummy struct {
	memory map[string]*dummy
}

type dummy struct {
	code  string
	count int
}

func NewRepositoryDummy() *RepositoryDummy {
	return &RepositoryDummy{make(map[string]*dummy)}
}

func (r *RepositoryDummy) Add(code *Code, phone *user.Phone, ttl time.Duration) error {
	d := r.memory[phone.String()]
	if d != nil {
		d.code = code.String()
		d.count++
	} else {
		r.memory[phone.String()] = &dummy{
			code:  code.String(),
			count: 1,
		}
	}

	return nil
}

func (r *RepositoryDummy) Get(phone *user.Phone) (*Code, error) {
	return Parse("1234")
}

func (r *RepositoryDummy) Count(phone *user.Phone) (int, error) {
	d := r.memory[phone.String()]
	if d != nil {
		return d.count, nil
	}

	return 0, nil
}

func (r *RepositoryDummy) DeleteLast(*user.Phone) error {
	return nil
}
