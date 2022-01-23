package controller

import (
	"bigfood/internal/authorization/actions/auth"
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
// todo: annotation
func (controller *Controller) Auth(c *gin.Context) {
	var message auth.Message
	if err := c.BindJSON(&message); err != nil {
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation
		return
	}

	response, err := controller.handlers.UserAuthHandler.Run(&message)
	if err != nil {
		server.NewResponseError(c, http.StatusInternalServerError, err) // todo: change code response
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
