package controller

import (
	"bigfood/internal/table/actions/tableDelete"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableDelete
// @Summary      Delete table
// @Security     ApiKeyAuth
// @Description  Delete table
// @Tags         table
// @Accept       json
// @Param        input  body  tableDelete.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /table [delete]
func (controller *Controller) tableDelete(c *gin.Context) {
	var message tableDelete.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo: permission Admin and table id

	err = controller.handlers.TableDelete.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
