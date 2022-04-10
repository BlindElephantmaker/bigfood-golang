package controller

import (
	"bigfood/internal/table/actions/tableCreate"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableCreate
// @Summary      Create table
// @Security     ApiKeyAuth
// @Description  Create table
// @Tags         table
// @Accept       json
// @Produce      json
// @Param        input  body      tableCreate.Message   true  "Body"
// @Success      200    {object}  table.Table           "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /table [post]
func (controller *Controller) tableCreate(c *gin.Context) {
	var message tableCreate.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo
	//if !userIsAdmin(c, message.CafeId) {
	//	server.AccessDenied(c)
	//	return
	//}

	response, err := controller.handlers.TableCreate.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
