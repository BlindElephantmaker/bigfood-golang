package controller

import (
	"bigfood/internal/helpers"
	"bigfood/internal/table/actions/createMass"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// tableCreateMass
// @Summary      Mass creation
// @Security     ApiKeyAuth
// @Description  Create N-quantity of tables
// @Tags         table
// @Accept       json
// @Produce      json
// @Param        input  body      createMass.Message    true  "Body"
// @Success      200    {object}  TableListResponse     "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /table/mass-create [post]
func (controller *Controller) tableCreateMass(c *gin.Context) {
	var message createMass.Message
	err := server.ParseJsonRequestToMessage(c, &message) // todo: strange handle case "quantity": 0
	if err != nil {
		return
	}
	if !userIsAdmin(c, message.CafeId) {
		server.AccessDenied(c)
		return
	}

	tables, err := controller.handlers.TableCreateMassHandler.Run(&message)
	if err == createMass.ErrorQuantityIsTooLow ||
		err == createMass.ErrorQuantityIsTooHigh ||
		err == helpers.ErrorInvalidUuid {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &TableListResponse{tables})
}