package controller

import (
	"bigfood/internal/table"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reserveListByTable
// @Summary      Get reserve list
// @Security     ApiKeyAuth
// @Description  Get reserve list by table
// @Tags         reserve
// @Produce      json
// @Param        table-id  path      string                       true  "table-id"
// @Success      200       {object}  reserveListByTable.Response  "Success"
// @Failure      400       {object}  server.ResponseError         "Invalid data"
// @Failure      401       {object}  server.ResponseError         "Access Denied"
// @Failure      500       {object}  server.ResponseError         "Internal Server Error"
// @Router       /reserve/table/{table-id} [get]
func (controller *Controller) reserveListByTable(c *gin.Context) {
	tableId, err := table.ParseId(c.Param("table-id"))
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	// todo: permissions

	response, err := controller.handlers.ReserveListByTable.Run(tableId)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
