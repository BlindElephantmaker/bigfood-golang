package userToken

import (
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgev" // todo: generate key.pem
	accessTTL  = 1 * helpers.Day
)

type AccessToken string

var (
	ErrorInvalidSigningMethod = errors.New("invalid signing method")
	ErrorInvalidClaims        = errors.New("invalid user claims")
)

func newAccess(userId user.Id, now time.Time) (AccessToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{ // todo: change to SigningMethodES256
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(accessTTL).Unix(),
		Id:        string(userId),
	})
	value, err := token.SignedString([]byte(signingKey))

	return AccessToken(value), err
}

func ParseAccess(token string) (*jwt.StandardClaims, error) {
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
		return nil, ErrorInvalidClaims
	}

	return claims, nil
}
