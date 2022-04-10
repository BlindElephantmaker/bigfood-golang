package controller

import (
	"bigfood/internal/reserve/actions/undelete"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reserveUndelete
// @Summary      Reserve undelete
// @Security     ApiKeyAuth
// @Description  Reserve undelete
// @Tags         reserve
// @Accept       json
// @Param        input  body  reserveUndelete.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /reserve/undelete [put]
func (controller *Controller) reserveUndelete(c *gin.Context) {
	var message reserveUndelete.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo: permissions
	//if !true {
	//	server.AccessDenied(c)
	//	return
	//}

	if err := controller.handlers.ReserveUndelete.Run(&message); err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
