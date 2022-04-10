package controller

import (
	"bigfood/internal/reserve/actions/delete"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reserveDelete
// @Summary      Delete reserve
// @Security     ApiKeyAuth
// @Description  Delete reserve
// @Tags         reserve
// @Accept       json
// @Param        input  body  reserveDelete.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /reserve [delete]
func (controller *Controller) reserveDelete(c *gin.Context) {
	var message reserveDelete.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo: permissions
	//if !true {
	//	server.AccessDenied(c)
	//	return
	//}

	if err := controller.handlers.ReserveDelete.Run(&message); err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
