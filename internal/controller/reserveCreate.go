package controller

import (
	"bigfood/internal/reserve/actions"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reserveCreate
// @Summary      Create reserve
// @Security     ApiKeyAuth
// @Description  Create reserve
// @Tags         reserve
// @Accept       json
// @Produce      json
// @Param        input  body      reserveAction.Message  true  "Body"
// @Success      200    {object}  reserve.Reserve        "Success"
// @Failure      400    {object}  server.ResponseError   "Invalid data"
// @Failure      401    {object}  server.ResponseError   "Access Denied"
// @Failure      422    {object}  server.ResponseError   "Unprocessable Entity"
// @Failure      500    {object}  server.ResponseError   "Internal Server Error"
// @Router       /reserve [post]
func (controller *Controller) reserveCreate(c *gin.Context) {
	var message reserveAction.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo permission

	response, err := controller.handlers.ReserveCreate.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
