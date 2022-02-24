package controller

import (
	"bigfood/internal/authorization/actions/auth"
	"bigfood/internal/authorization/smsCode"
	"bigfood/internal/helpers"
	"bigfood/internal/user"
	"bigfood/internal/user/userToken"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthResponse struct {
	IsNew   bool                   `json:"is-new"`
	UserId  helpers.Uuid           `json:"user-id" example:"UUID"`
	Access  userToken.AccessToken  `json:"access-token"`
	Refresh userToken.RefreshToken `json:"refresh-token" example:"UUID"`
}

// auth
// @Summary      User authorization
// @Description  Get new refresh and access tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      auth.Message          true  "Body"
// @Success      200    {object}  AuthResponse          "Success"
// @Failure      422    {object}  server.ResponseError  "SMS code not confirmed"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /auth [post]
func (controller *Controller) auth(c *gin.Context) {
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
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &AuthResponse{
		IsNew:   response.IsNew,
		UserId:  response.UserToken.UserId,
		Access:  *response.UserToken.Access,
		Refresh: response.UserToken.Refresh,
	})
}
