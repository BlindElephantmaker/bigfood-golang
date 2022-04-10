package controller

import (
	"bigfood/internal/reserve/actions/edit"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reserveEdit
// @Summary      Reserve edit
// @Security     ApiKeyAuth
// @Description  Reserve edit
// @Tags         reserve
// @Accept       json
// @Produce      json
// @Param        input  body      reserveEdit.Message   true  "Body"
// @Success      200    {object}  reserve.Reserve       "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      422    {object}  server.ResponseError  "Unprocessable Entity"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /reserve [put]
func (controller *Controller) reserveEdit(c *gin.Context) {
	var message reserveEdit.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo permission

	response, err := controller.handlers.ReserveEdit.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
