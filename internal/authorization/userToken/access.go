package userToken

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const (
	signingKey = "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgev" // todo: generate key.pem
)

type AccessToken struct {
	value string
}

var (
	ErrorInvalidSigningMethod = errors.New("invalid signing method")
	ErrorInvalidTokenClaims   = errors.New("token claims are not type of StandardClaims")
)

func NewAccess(userId *uuid.UUID, time *time.Time, ttl time.Duration) (*AccessToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{ // todo: change to SigningMethodES256
		Id:        userId.String(),
		IssuedAt:  time.Unix(),
		ExpiresAt: time.Add(ttl).Unix(),
	})
	value, err := token.SignedString([]byte(signingKey))

	return &AccessToken{value}, err
}

func ParseAccess(token string) (*uuid.UUID, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // todo: check after change SigningMethodHS256 to SigningMethodES256
			return nil, ErrorInvalidSigningMethod
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, ErrorInvalidTokenClaims
	}

	id, err := uuid.Parse(claims.Id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (t *AccessToken) String() string {
	return t.value
}
