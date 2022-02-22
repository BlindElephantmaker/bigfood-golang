package userToken

import (
	"bigfood/internal/cafe/cafeUser/userRole"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgev" // todo: generate key.pem
)

type AccessToken struct {
	value string
}

type UserClaims struct {
	jwt.StandardClaims
	userRole.Permissions
}

var (
	ErrorInvalidSigningMethod = errors.New("invalid signing method")
	ErrorInvalidClaims        = errors.New("invalid user claims")
)

func NewAccess(permissions *userRole.Permissions, time *time.Time, ttl time.Duration) (*AccessToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{ // todo: change to SigningMethodES256
		jwt.StandardClaims{
			IssuedAt:  time.Unix(),
			ExpiresAt: time.Add(ttl).Unix(),
		},
		*permissions,
	})
	value, err := token.SignedString([]byte(signingKey))

	return &AccessToken{value}, err
}

func ParseAccess(token string) (*UserClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // todo: check after change SigningMethodHS256 to SigningMethodES256
			return nil, ErrorInvalidSigningMethod
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*UserClaims)
	if !ok {
		return nil, ErrorInvalidClaims
	}

	return claims, nil
}

func (t *AccessToken) String() string {
	return t.value
}
