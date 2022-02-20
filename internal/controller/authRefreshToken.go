package controller

import (
	"bigfood/internal/authorization/actions/refreshToken"
	"bigfood/internal/authorization/userToken"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RefreshTokenResponse struct {
	Id      string `json:"id" example:"UUID"`
	Access  string `json:"access"`
	Refresh string `json:"refresh" example:"UUID"`
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
	if err == userToken.ErrorInvalidRefreshTokenFormat {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &RefreshTokenResponse{
		Id:      token.UserId.String(),
		Access:  token.Access.String(),
		Refresh: token.Refresh.String(),
	})
}
