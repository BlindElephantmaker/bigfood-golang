package controller

import (
	"bigfood/internal/authorization/action/refreshToken"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// refreshToken
// @Summary      Refresh jwt token
// @Description  Refresh user refresh and access tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      refreshToken.Message   true  "Body"
// @Success      200    {object}  refreshToken.Response  "Success"
// @Failure      400    {object}  server.ResponseError   "Invalid refresh token format"
// @Failure      500    {object}  server.ResponseError   "Internal Server Error"
// @Router       /auth/refresh-token [put]
func (controller *Controller) refreshToken(c *gin.Context) {
	var message refreshToken.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	response, err := controller.handlers.RefreshToken.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
