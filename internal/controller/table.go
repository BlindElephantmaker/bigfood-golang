package controller

import (
	"bigfood/internal/cafe/cafeUser/role"
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"github.com/gin-gonic/gin"
)

type TableListResponse struct {
	Tables []*table.Table `json:"tables"`
}

func userCanEditTable(c *gin.Context, cafeId helpers.Uuid) bool {
	return userHasRole(c, cafeId, role.Admin)
}

func userCanViewTable(c *gin.Context, cafeId helpers.Uuid) bool {
	return userHasRole(c, cafeId, role.Admin) || userHasRole(c, cafeId, role.Hostess)
}
