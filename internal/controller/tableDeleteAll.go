package controller

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table/actions/tableDeleteAll"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TableDeleteAllResponse struct {
	Success bool `json:"success" example:"true"`
}

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
	if !userIsAdmin(c, message.CafeId) {
		server.AccessDenied(c)
		return
	}

	err = controller.handlers.TableDeleteAllHandler.Run(&message)
	if err == helpers.ErrorInvalidUuid {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
