package controller

import (
	"bigfood/internal/user/actions/userEdit"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// userEdit
// @Summary      Edit user
// @Security     ApiKeyAuth
// @Description  Edit user information
// @Tags         user
// @Accept       json
// @Param        input  body  userEdit.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Invalid user data"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /user [put]
func (controller *Controller) userEdit(c *gin.Context) {
	var message userEdit.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	message.Id = getUserId(c)

	err = controller.handlers.UserEditHandler.Run(&message)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
