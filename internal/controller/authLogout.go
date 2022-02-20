package controller

import (
	"bigfood/internal/authorization/actions/userLogout"
	"bigfood/internal/authorization/userToken"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type AuthLogoutResponse struct {
	Success bool `json:"success" example:"true"`
}

// Logout
// @Summary      User logout
// @Security     ApiKeyAuth
// @Description  Delete user refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      userLogout.Message  true  "Body"
// @Success      200    {object}  AuthLogoutResponse
// @Failure      400    {object}  server.ResponseError  "Invalid refresh token"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /auth/logout [delete]
func (controller *Controller) Logout(c *gin.Context) {
	var message userLogout.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	id, _ := c.Get(UserId)
	message.UserId = id.(*uuid.UUID)

	err = controller.handlers.UserLogoutHandler.Run(&message)
	if err == userToken.ErrorInvalidRefreshTokenFormat {
		server.NewResponseError(c, http.StatusBadRequest, err)
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &AuthLogoutResponse{
		Success: true,
	})
}
