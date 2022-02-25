package user

type Service struct {
	userRepository Repository
}

func NewService(users Repository) *Service {
	return &Service{users}
}

func (s *Service) GetOrNewUser(userPhone Phone) (*User, error) {
	isExist, err := s.userRepository.IsExistByPhone(userPhone)
	if err != nil {
		return nil, err
	}

	if isExist {
		return s.userRepository.GetByPhone(userPhone)
	}

	newUser := New(userPhone)
	return newUser, s.userRepository.Add(newUser)
}
