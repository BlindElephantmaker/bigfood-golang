package controller

import (
	"bigfood/internal/authorization/actions/userLogout"
	"bigfood/internal/authorization/userToken"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// logout
// todo: Deprecated. Any reason to do this?
// @Summary      User logout (Deprecated)
// @Security     ApiKeyAuth
// @Description  Any reason to do this?
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  userLogout.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid refresh token"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /auth/logout [delete]
func (controller *Controller) logout(c *gin.Context) {
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

	c.AbortWithStatus(http.StatusOK)
}
