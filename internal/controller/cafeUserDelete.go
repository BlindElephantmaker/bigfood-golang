package controller

import (
	"bigfood/internal/cafeUser/actions/delete"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cafeUserDelete
// @Summary      Delete cafe user
// @Security     ApiKeyAuth
// @Description  Delete cafe user
// @Tags         cafe user
// @Accept       json
// @Param        input  body  cafeUserDelete.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /cafe/user [delete]
func (controller *Controller) cafeUserDelete(c *gin.Context) {
	var message cafeUserDelete.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo: permissions role and cafeUserId

	err = controller.handlers.CafeUserDeleteHandler.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
