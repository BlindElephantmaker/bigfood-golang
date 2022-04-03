package user

import "bigfood/internal/helpers"

type Service struct {
	userRepository Repository
}

func NewService(users Repository) *Service {
	return &Service{users}
}

func (s *Service) GetOrNewUser(userPhone helpers.Phone) (*User, error) {
	user, err := s.userRepository.GetByPhone(userPhone)
	if err == NotExist {
		user = New(userPhone)
		err = s.userRepository.Add(user)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
