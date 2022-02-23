package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message" example:"Error message"`
}

var ErrorParseJsonRequestToMessage = errors.New("failed parse json request to message")

func NewResponseError(c *gin.Context, code int, err error) {
	message := err.Error()
	logrus.Error(message)
	c.AbortWithStatusJSON(code, ResponseError{message})
}

func InternalServerError(c *gin.Context, err error) {
	logrus.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError{
		Message: "Internal Server Error",
	})
}

func AccessDenied(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseError{
		Message: "Access Denied",
	})
}

func ParseJsonRequestToMessage(c *gin.Context, message interface{}) error {
	err := c.BindJSON(&message)
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, ErrorParseJsonRequestToMessage)
	}
	return err
}
