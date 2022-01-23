package controller

import (
	"bigfood/internal/authorization/actions/auth"
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/user"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthResponse struct {
	Success bool   `json:"success" example:"true"`
	IsNew   bool   `json:"is-new"`
	UserId  string `json:"user-id" example:"UUID"`
	Access  string `json:"access-token"`
	Refresh string `json:"refresh-token" example:"UUID"`
}

func (controller *Controller) Auth(c *gin.Context) {
	var message auth.Message
	if err := c.BindJSON(&message); err != nil {
		// todo: message
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation
		return
	}

	response, err := controller.handlers.UserAuthHandler.Run(&message)
	if err == user.ErrorPhoneNumberIsInvalid || err == smsCode.ErrorSmsCodeIsInvalid {
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation two
		return
	}
	if err == auth.ErrorSmsCodeNotConfirmed {
		server.NewResponseError(c, http.StatusUnprocessableEntity, err) // todo: annotation
		return
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err) // todo: annotation
		return
	}

	c.JSON(http.StatusOK, &AuthResponse{
		Success: true,
		IsNew:   response.IsNew,
		UserId:  response.UserToken.UserId.String(),
		Access:  response.UserToken.Access.String(),
		Refresh: response.UserToken.Refresh.String(),
	})
}
