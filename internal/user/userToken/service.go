package userToken

import (
	"bigfood/internal/helpers"
	"bigfood/internal/user"
)

type Service struct {
	userTokenRepository Repository
}

func NewService(userTokens Repository) *Service {
	return &Service{
		userTokenRepository: userTokens,
	}
}

func (s *Service) CreateTokens(userId user.Id) (AccessToken, *UserToken, error) {
	now := helpers.NowTime()
	access, err := newAccess(userId, now)
	if err != nil {
		return "", nil, err
	}

	userToken := newUserToken(userId, now)
	if err := s.userTokenRepository.Add(userToken); err != nil {
		return "", nil, err
	}

	return access, userToken, nil
}
