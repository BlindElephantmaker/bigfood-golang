package controller

import (
	"bigfood/internal/cafe"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableList
// @Summary      Get table list
// @Security     ApiKeyAuth
// @Description  Get table list by cafe
// @Tags         table
// @Produce      json
// @Param        cafe-id  path      string                true  "cafe-id"
// @Success      200      {object}  tableList.Response    "Success"
// @Failure      400      {object}  server.ResponseError  "Invalid data"
// @Failure      401      {object}  server.ResponseError  "Access Denied"
// @Failure      500      {object}  server.ResponseError  "Internal Server Error"
// @Router       /table/list/{cafe-id} [get]
func (controller *Controller) tableList(c *gin.Context) {
	cafeId, err := cafe.ParseId(c.Param("cafe-id"))
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	// todo permission

	response, err := controller.handlers.TableList.Run(cafeId)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
