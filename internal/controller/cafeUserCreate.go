package controller

import (
	"bigfood/internal/cafeUser/actions/create"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cafeUserCreate
// @Summary      Create cafe user
// @Security     ApiKeyAuth
// @Description  Create cafe user
// @Tags         cafe user
// @Accept       json
// @Produce      json
// @Param        input  body      cafeUserCreate.Message  true  "Body"
// @Success      200    {object}  actions.Response        "Success"
// @Failure      400    {object}  server.ResponseError    "Invalid data"
// @Failure      401    {object}  server.ResponseError    "Access Denied"
// @Failure      422    {object}  server.ResponseError    "Cafe user already exist"
// @Failure      500    {object}  server.ResponseError    "Internal Server Error"
// @Router       /cafe/user [post]
func (controller *Controller) cafeUserCreate(c *gin.Context) {
	var message cafeUserCreate.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo
	//if !userIsAdmin(c, message.CafeId) {
	//	server.AccessDenied(c)
	//	return
	//}

	response, err := controller.handlers.CafeUserCreateHandler.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
