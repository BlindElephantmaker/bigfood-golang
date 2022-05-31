package controller

import (
	"bigfood/internal/table"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableGet
// @Summary      Get table
// @Security     ApiKeyAuth
// @Description  Get table by table id
// @Tags         table
// @Produce      json
// @Param        table-id  path      string                true  "table-id"
// @Success      200       {object}  table.Table           "Success"
// @Failure      400       {object}  server.ResponseError  "Invalid data"
// @Failure      401       {object}  server.ResponseError  "Access Denied"
// @Failure      500       {object}  server.ResponseError  "Internal Server Error"
// @Router       /table/{table-id} [get]
func (controller *Controller) tableGet(c *gin.Context) {
	tableId, err := table.ParseId(c.Param("table-id"))
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	// todo permission

	responce, err := controller.handlers.TableGet.Run(tableId)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, responce)
}
