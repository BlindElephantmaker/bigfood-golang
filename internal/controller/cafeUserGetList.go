package controller

import (
	"bigfood/internal/cafeUser/actions/list"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cafeUserList
// @Summary      Get cafe user list
// @Security     ApiKeyAuth
// @Description  Get cafe user list
// @Tags         cafe user
// @Accept       json
// @Produce      json
// @Param        input  body      cafeUserList.Message   true  "Body"
// @Success      200    {object}  cafeUserList.Response  "Success"
// @Failure      400    {object}  server.ResponseError   "Invalid data"
// @Failure      401    {object}  server.ResponseError   "Access Denied"
// @Failure      500    {object}  server.ResponseError   "Internal Server Error"
// @Router       /cafe/user/list [get]
func (controller *Controller) cafeUserList(c *gin.Context) {
	var message cafeUserList.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}
	// todo
	//if !userIsHostess(c, message.CafeId) {
	//	server.AccessDenied(c)
	//	return
	//}

	response, err := controller.handlers.CafeUserListHandler.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
