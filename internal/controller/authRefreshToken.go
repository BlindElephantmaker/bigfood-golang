package controller

import (
	"bigfood/internal/authorization/action/refreshToken"
	"bigfood/internal/helpers"
	"bigfood/internal/user/userToken"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshTokenResponse struct {
	Id      helpers.Uuid           `json:"id" example:"UUID"`
	Access  userToken.AccessToken  `json:"access"`
	Refresh userToken.RefreshToken `json:"refresh" example:"UUID"`
}

// refreshToken
// @Summary      Refresh jwt token
// @Description  Refresh user refresh and access tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      refreshToken.Message  true  "Body"
// @Success      200    {object}  RefreshTokenResponse  "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid refresh token format"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /auth/refresh-token [put]
func (controller *Controller) refreshToken(c *gin.Context) {
	var message refreshToken.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	token, err := controller.handlers.RefreshTokenHandler.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &RefreshTokenResponse{
		Id:      token.UserId,
		Access:  *token.Access,
		Refresh: token.Refresh,
	})
}
