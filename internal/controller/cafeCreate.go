package controller

import (
	"bigfood/internal/cafe"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CafeCreateResponse struct {
	Id cafe.Id `json:"id" example:"uuid created cafe"`
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
	cafeId, err := controller.handlers.CafeCreateHandler.Run(getUserId(c))
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &CafeCreateResponse{
		Id: cafeId,
	})
}
