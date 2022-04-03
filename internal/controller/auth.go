package controller

import (
	"bigfood/internal/authorization/action/auth"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// auth
// @Summary      User authorization
// @Description  Get new refresh and access tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      auth.Message          true  "Body"
// @Success      200    {object}  auth.Response         "Success"
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
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
