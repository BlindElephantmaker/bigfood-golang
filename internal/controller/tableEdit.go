package controller

import (
	"bigfood/internal/table/actions/tableEdit"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableEdit
// @Summary      Edit table
// @Security     ApiKeyAuth
// @Description  Edit table
// @Tags         table
// @Accept       json
// @Param        input  body  tableEdit.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /table [put]
func (controller *Controller) tableEdit(c *gin.Context) {
	var message tableEdit.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo permission

	err = controller.handlers.TableEdit.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
