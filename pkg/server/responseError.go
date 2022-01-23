package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
