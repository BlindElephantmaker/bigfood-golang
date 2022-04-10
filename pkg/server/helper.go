package server

import (
	"bigfood/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
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

func StatusBadRequest(c *gin.Context, err error) {
	logrus.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusBadRequest, ResponseError{
		Message: err.Error(),
	})
}

// todo: not used
//func AccessDenied(c *gin.Context) {
//	c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseError{
//		Message: "Access Denied",
//	})
//}

func ParseJsonRequestToMessage(c *gin.Context, message interface{}) error {
	err := c.BindJSON(&message)
	if err != nil {
		NewResponseError(c, http.StatusBadRequest, ErrorParseJsonRequestToMessage) // todo: show validation error
	}
	return err
}

func GetQueryParamInt(c *gin.Context, key string, def int) int {
	valStr := c.Request.URL.Query().Get(key)
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return def
	}
	return valInt
}

func GetQueryParamTime(c *gin.Context, key string) (time.Time, error) {
	valStr := c.Request.URL.Query().Get(key)
	return helpers.ParseTime(valStr) // todo: message for this error
}
