package controller

import (
	"bigfood/internal/authorization/actions/refreshToken"
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

// RefreshToken
// todo: annotation
func (controller *Controller) RefreshToken(c *gin.Context) {
	var message refreshToken.Message
	if err := c.BindJSON(&message); err != nil {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}

	userToken, err := controller.handlers.RefreshTokenHandler.Run(&message)
	if err != nil {
		server.NewResponseError(c, http.StatusInternalServerError, err) // todo: code response
		return
	}

	c.JSON(http.StatusOK, &RefreshTokenResponse{
		Success: true,
		Id:      userToken.UserId.String(),
		Access:  userToken.Access.String(),
		Refresh: userToken.Refresh.String(),
	})
}
