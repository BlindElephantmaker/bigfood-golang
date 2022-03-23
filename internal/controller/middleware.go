package controller

import (
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"bigfood/pkg/server"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userId              = "userId"
)

var (
	errorEmptyAuthorizationHeader   = errors.New("empty authorization header")
	errorInvalidAuthorizationHeader = errors.New("invalid authorization header")
)

func (controller *Controller) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		server.NewResponseError(c, http.StatusUnauthorized, errorEmptyAuthorizationHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		server.NewResponseError(c, http.StatusUnauthorized, errorInvalidAuthorizationHeader)
		return
	}

	claims, err := userToken.ParseAccess(headerParts[1])
	if err != nil {
		server.NewResponseError(c, http.StatusUnauthorized, err)
		return
	}

	id, err := user.ParseId(claims.Id)
	if err != nil {
		server.NewResponseError(c, http.StatusUnauthorized, err)
		return
	}
	c.Set(userId, id)
}

func getUserId(c *gin.Context) user.Id {
	userId, _ := c.Get(userId)
	return userId.(user.Id)
}
