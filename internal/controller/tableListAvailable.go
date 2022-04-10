package controller

import (
	"bigfood/internal/cafe"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableListAvailable
// @Summary      Get list available tables
// @Security     ApiKeyAuth
// @Description  Get list available tables
// @Tags         table
// @Produce      json
// @Param        cafe-id     path      string                       true  "cafe-id"
// @Param        from-date   query     string                       true  "from"
// @Param        until-date  query     string                       true  "until"
// @Success      200         {object}  tableListAvailable.Response  "Success"
// @Failure      400         {object}  server.ResponseError         "Invalid data"
// @Failure      401         {object}  server.ResponseError         "Access Denied"
// @Failure      500         {object}  server.ResponseError         "Internal Server Error"
// @Router       /table/list/{cafe-id}/available [get]
func (controller *Controller) tableListAvailable(c *gin.Context) {
	cafeId, err := cafe.ParseId(c.Param("cafe-id"))
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	from, err := server.GetQueryParamTime(c, "from-date")
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	until, err := server.GetQueryParamTime(c, "until-date")
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	// todo: permissions

	response, err := controller.handlers.TableListAvailable.Run(cafeId, from, until)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
