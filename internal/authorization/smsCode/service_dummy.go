package smsCode

import "bigfood/internal/helpers"

// todo: code it

type ServiceDummy struct{}

func NewServiceDummy() *ServiceDummy {
	return &ServiceDummy{}
}

func (s *ServiceDummy) Send(text string, phone helpers.Phone) error {
	return nil
}
