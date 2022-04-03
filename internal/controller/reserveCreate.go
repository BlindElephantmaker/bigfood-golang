package controller

import (
	"bigfood/internal/reserve/actions/create"
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
// @Param        input  body      reserveCreate.Message  true  "Body"
// @Success      200    {object}  reserve.Reserve        "Success"
// @Failure      400    {object}  server.ResponseError   "Invalid data"
// @Failure      401    {object}  server.ResponseError   "Access Denied"
// @Failure      500    {object}  server.ResponseError   "Internal Server Error"
// @Router       /reserve [post]
func (controller *Controller) reserveCreate(c *gin.Context) {
	var message reserveCreate.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	//todo: permissions
	//if !true {
	//	server.AccessDenied(c)
	//	return
	//}

	response, err := controller.handlers.ReserveCreateHandler.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
