package controller

import (
	"bigfood/internal/reserve"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// reserveGet
// @Summary      Get reserve
// @Security     ApiKeyAuth
// @Description  Get reserve
// @Tags         reserve
// @Produce      json
// @Param        reserve-id  path      string                true  "reserve-id"
// @Success      200         {object}  reserve.Reserve       "Success"
// @Failure      400         {object}  server.ResponseError  "Invalid data"
// @Failure      401         {object}  server.ResponseError  "Access Denied"
// @Failure      500         {object}  server.ResponseError  "Internal Server Error"
// @Router       /reserve/{reserve-id} [get]
func (controller *Controller) reserveGet(c *gin.Context) {
	reserveId, err := reserve.ParseId(c.Param("reserve-id"))
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	// todo: permissions
	//if !true {
	//	server.AccessDenied(c)
	//	return
	//}

	response, err := controller.handlers.ReserveGet.Run(reserveId)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
