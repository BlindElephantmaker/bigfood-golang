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

// Auth
// @Summary      User authorization
// @Description  Get new refresh and access tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      auth.Message  true  "Body"
// @Success      200    {object}  AuthResponse
// @Failure      422    {object}  server.ResponseError  "SMS code not confirmed"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /auth [post]
func (controller *Controller) Auth(c *gin.Context) {
	var message auth.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	response, err := controller.handlers.UserAuthHandler.Run(&message)
	if err == user.ErrorPhoneNumberIsInvalid || err == smsCode.ErrorSmsCodeIsInvalid {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err == auth.ErrorSmsCodeNotConfirmed {
		server.NewResponseError(c, http.StatusUnprocessableEntity, err)
		return
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err)
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
