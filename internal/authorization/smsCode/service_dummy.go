package smsCode

import "bigfood/internal/user"

// todo: code it

type ServiceDummy struct{}

func NewServiceDummy() *ServiceDummy {
	return &ServiceDummy{}
}

func (s *ServiceDummy) Send(text string, phone *user.Phone) error {
	return nil
}
