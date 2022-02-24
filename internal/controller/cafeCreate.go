package controller

import (
	"bigfood/internal/helpers"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CafeCreateResponse struct {
	Id helpers.Uuid `json:"id" example:"uuid created cafe"`
}

// cafeCreate
// @Summary      Create cafe
// @Security     ApiKeyAuth
// @Description  Create new cafe
// @Tags         cafe
// @Produce      json
// @Success      200  {object}  CafeCreateResponse    "Success"
// @Failure      500  {object}  server.ResponseError  "Internal Server Error"
// @Router       /cafe [post]
func (controller *Controller) cafeCreate(c *gin.Context) {
	id := getUserId(c)

	cafeId, err := controller.handlers.CafeCreateHandler.Run(id)
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &CafeCreateResponse{
		Id: cafeId,
	})
}
