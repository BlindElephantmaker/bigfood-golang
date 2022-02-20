package controller

import (
	"bigfood/internal/user"
	"bigfood/internal/user/actions/userEdit"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserEditResponse struct {
	Success bool `json:"success" example:"true"`
}

// userEdit
// @Summary      Edit user
// @Security     ApiKeyAuth
// @Description  Edit user information
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        input  body      userEdit.Message  true  "Body"
// @Success      200    {object}  UserEditResponse
// @Failure      400    {object}  server.ResponseError  "Invalid user data"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /user [put]
func (controller *Controller) userEdit(c *gin.Context) {
	var message userEdit.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	id, _ := c.Get(UserId)
	message.Id = id.(*uuid.UUID)

	err = controller.handlers.UserEditHandler.Run(&message)
	if err == user.ErrorUsernameIsTooShort || err == user.ErrorUsernameIsTooLong || err == userEdit.ErrorUserNameIsEmpty {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &UserEditResponse{
		Success: true,
	})
}
