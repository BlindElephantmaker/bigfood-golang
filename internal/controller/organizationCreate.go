package controller

import (
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrganizationCreateResponse struct {
	Id string `json:"id" example:"uuid created organization"`
}

// organizationCreate
// @Summary      Create organization
// @Security     ApiKeyAuth
// @Description  Create new organization
// @Tags         organization
// @Produce      json
// @Success      200  {object}  OrganizationCreateResponse  "Success"
// @Failure      500  {object}  server.ResponseError        "Internal Server Error"
// @Router       /organization [post]
func (controller *Controller) organizationCreate(c *gin.Context) {
	id := getUserId(c)

	organizationId, err := controller.handlers.OrganizationCreateHandler.Run(id)
	if err != nil {
		server.NewResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &OrganizationCreateResponse{
		Id: organizationId.String(),
	})
}
