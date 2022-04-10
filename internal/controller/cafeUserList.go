package controller

import (
	"bigfood/internal/cafe"
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
// @Param        cafe-id  path      string                true  "cafe-id"
// @Success      200    {object}  cafeUserList.Response  "Success"
// @Failure      400    {object}  server.ResponseError   "Invalid data"
// @Failure      401    {object}  server.ResponseError   "Access Denied"
// @Failure      500    {object}  server.ResponseError   "Internal Server Error"
// @Router       /cafe/user/list/{cafe-id} [get]
func (controller *Controller) cafeUserList(c *gin.Context) {
	cafeId, err := cafe.ParseId(c.Param("cafe-id"))
	if err != nil {
		server.StatusBadRequest(c, err)
		return
	}
	// todo permission

	response, err := controller.handlers.CafeUserList.Run(cafeId)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
