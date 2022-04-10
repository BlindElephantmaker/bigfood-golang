package controller

import (
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cafeCreate
// @Summary      Create cafe
// @Security     ApiKeyAuth
// @Description  Create new cafe
// @Tags         cafe
// @Produce      json
// @Success      200  {object}  createCafe.Response   "Success"
// @Failure      500  {object}  server.ResponseError  "Internal Server Error"
// @Router       /cafe [post]
func (controller *Controller) cafeCreate(c *gin.Context) {
	response, err := controller.handlers.CafeCreate.Run(getUserId(c))
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
