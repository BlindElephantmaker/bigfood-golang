package controller

import (
	"bigfood/internal/table"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reserveHistoryByTable
// @Summary      Get reserve history
// @Security     ApiKeyAuth
// @Description  Get reserve history list by table
// @Tags         reserve
// @Produce      json
// @Param        table-id  path      string                          true   "table-id"
// @Param        limit     query     int                             false  "limit"
// @Param        offset    query     int                             false  "offset"
// @Success      200       {object}  reserveHistoryByTable.Response  "Success"
// @Failure      400       {object}  server.ResponseError            "Invalid data"
// @Failure      401       {object}  server.ResponseError            "Access Denied"
// @Failure      500       {object}  server.ResponseError            "Internal Server Error"
// @Router       /reserve/table/{table-id}/history [get]
func (controller *Controller) reserveHistoryByTable(c *gin.Context) {
	tableId, err := table.ParseId(c.Param("table-id"))
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	limit := server.GetQueryParamInt(c, "limit", 20)
	offset := server.GetQueryParamInt(c, "offset", 0)
	// todo: permissions

	response, err := controller.handlers.ReserveHistoryByTableHandler.Run(tableId, limit, offset)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
