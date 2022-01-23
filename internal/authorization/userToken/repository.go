package userToken

type Repository interface {
	Add(*UserToken) error
	Get(*RefreshToken) (*UserToken, error)
	Refresh(newToken, oldToken *UserToken) error
}
