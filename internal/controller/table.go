package controller

import (
	"bigfood/internal/cafe/cafeUser/role"
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"bigfood/internal/table/createMass"
	"bigfood/internal/table/getList"
	"bigfood/internal/table/tableEdit"
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
	}

	tables, err := controller.handlers.TableGetListHandler.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &TableListResponse{tables})
}

// tableEdit
// @Summary      Edit table
// @Security     ApiKeyAuth
// @Description  Edit table
// @Tags         table
// @Accept       json
// @Param        input  body  tableEdit.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /table [put]
func (controller *Controller) tableEdit(c *gin.Context) {
	var message tableEdit.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo: permission Admin and table id

	err = controller.handlers.TableEditHandler.Run(&message)
	if helpers.ErrorInvalidUuid == err ||
		table.ErrorTableTitleIsTooLong == err ||
		table.ErrorTableTitleIsTooShort == err ||
		table.ErrorTableCommentIsTooLong == err ||
		table.ErrorTableSeatsInvalidValue == err {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

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
	if !userCanEditTable(c, message.CafeId) {
		server.AccessDenied(c)
	}

	tables, err := controller.handlers.TableCreateMassHandler.Run(&message)
	if err == createMass.ErrorQuantityIsTooLow || err == createMass.ErrorQuantityIsTooHigh {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &TableListResponse{tables})
}

type TableListResponse struct {
	Tables []*table.Table `json:"tables"`
}

func userCanEditTable(c *gin.Context, cafeId helpers.Uuid) bool {
	return userHasRole(c, cafeId, role.Admin)
}

func userCanViewTable(c *gin.Context, cafeId helpers.Uuid) bool {
	return userHasRole(c, cafeId, role.Admin) || userHasRole(c, cafeId, role.Hostess)
}
