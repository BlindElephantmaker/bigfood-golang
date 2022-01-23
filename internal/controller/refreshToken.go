package controller

import (
	"bigfood/internal/authorization/actions/refreshToken"
	"bigfood/internal/authorization/userToken"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshTokenResponse struct {
	Success bool   `json:"success" example:"true"`
	Id      string `json:"id" example:"UUID"`
	Access  string `json:"access"`
	Refresh string `json:"refresh" example:"UUID"`
}

func (controller *Controller) RefreshToken(c *gin.Context) {
	var message refreshToken.Message
	if err := c.BindJSON(&message); err != nil {
		// todo: message
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation
		return
	}

	token, err := controller.handlers.RefreshTokenHandler.Run(&message)
	if err == userToken.ErrorInvalidRefreshTokenFormat {
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation
		return
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err) // todo: annotation
		return
	}

	c.JSON(http.StatusOK, &RefreshTokenResponse{
		Success: true,
		Id:      token.UserId.String(),
		Access:  token.Access.String(),
		Refresh: token.Refresh.String(),
	})
}
