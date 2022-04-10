package controller

import (
	"bigfood/internal/cafeUser/actions/edit"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cafeUserEdit
// @Summary      Edit cafe user
// @Security     ApiKeyAuth
// @Description  Edit cafe user
// @Tags         cafe user
// @Accept       json
// @Produce      json
// @Param        input  body      cafeUserEdit.Message  true  "Body"
// @Success      200    {object}  actions.Response      "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid data"
// @Failure      401    {object}  server.ResponseError  "Access Denied"
// @Failure      422    {object}  server.ResponseError  "Owner role could not be set"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /cafe/user [put]
func (controller *Controller) cafeUserEdit(c *gin.Context) {
	var message cafeUserEdit.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo permission

	response, err := controller.handlers.CafeUserEdit.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
