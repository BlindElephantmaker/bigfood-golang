package controller

import (
	"bigfood/internal/table/actions/tableDeleteAll"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableDeleteAll
// @Summary      Delete all tables
// @Security     ApiKeyAuth
// @Description  Delete all tables
// @Tags         table
// @Accept       json
// @Param        input  body  tableDeleteAll.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /table/delete-all [delete]
func (controller *Controller) tableDeleteAll(c *gin.Context) {
	var message tableDeleteAll.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo permission

	err = controller.handlers.TableDeleteAll.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
