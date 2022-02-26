package controller

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
	"bigfood/internal/user/userToken"
	"bigfood/pkg/server"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	claims              = "claims"
)

var (
	ErrorEmptyAuthorizationHeader   = errors.New("empty authorization header")
	ErrorInvalidAuthorizationHeader = errors.New("invalid authorization header")
)

func (controller *Controller) userIdentity(c *gin.Context) {
	// todo: how compare cases?
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		server.NewResponseError(c, http.StatusUnauthorized, ErrorEmptyAuthorizationHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		server.NewResponseError(c, http.StatusUnauthorized, ErrorInvalidAuthorizationHeader)
		return
	}

	userClaims, err := userToken.ParseAccess(headerParts[1])
	if err != nil {
		server.NewResponseError(c, http.StatusUnauthorized, err)
		return
	}

	c.Set(claims, userClaims)
}

func getClaims(c *gin.Context) *userToken.UserClaims {
	permissions, _ := c.Get(claims)
	return permissions.(*userToken.UserClaims)
}

func getUserId(c *gin.Context) helpers.Uuid {
	claims := getClaims(c)
	return claims.Permissions.UserId
}

func userHasRole(c *gin.Context, cafeId helpers.Uuid, role cafeUser.Role) bool {
	claims := getClaims(c)
	return claims.Permissions.HasRole(cafeId, role)
}
