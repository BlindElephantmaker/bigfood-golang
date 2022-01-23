package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ResponseError struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message"`
}

func NewResponseError(c *gin.Context, code int, err error) {
	message := err.Error()
	logrus.Error(message)
	c.AbortWithStatusJSON(code, ResponseError{false, message})
}

func NewResponseInternalServerError(c *gin.Context, err error) {
	logrus.Error(err.Error())
	c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError{
		Success: false,
		Message: "Internal Server Error",
	})
}
