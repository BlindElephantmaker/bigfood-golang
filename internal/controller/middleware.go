package controller

import (
	"bigfood/internal/authorization/userToken"
	"bigfood/pkg/server"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	UserId              = "UserId"
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

	id, err := userToken.ParseAccess(headerParts[1])
	if err != nil {
		server.NewResponseError(c, http.StatusUnauthorized, err)
		return
	}

	c.Set(UserId, id)
}

func getUserId(c *gin.Context) *uuid.UUID {
	id, _ := c.Get(UserId)
	return id.(*uuid.UUID)
}
