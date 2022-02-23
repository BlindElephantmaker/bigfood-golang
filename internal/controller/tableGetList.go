package controller

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table/actions/getList"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableGetList
// @Summary      Get table list
// @Security     ApiKeyAuth
// @Description  Get table list of cafe
// @Tags         table
// @Accept       json
// @Produce      json
// @Param        input  body      getList.Message       true  "Body"
// @Success      200    {object}  TableListResponse     "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /table/list [get]
func (controller *Controller) tableGetList(c *gin.Context) {
	var message getList.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	if !userCanViewTable(c, message.CafeId) {
		server.AccessDenied(c)
		return
	}

	tables, err := controller.handlers.TableGetListHandler.Run(&message)
	if err == helpers.ErrorInvalidUuid {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		server.InternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, &TableListResponse{tables})
}
