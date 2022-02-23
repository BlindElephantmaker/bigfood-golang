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

func userCanEditTable(c *gin.Context, cafeId string) bool {
	return userHasRole(c, helpers.Uuid(cafeId), role.Admin)
}

func userCanViewTable(c *gin.Context, cafeId string) bool {
	return userHasRole(c, helpers.Uuid(cafeId), role.Admin) ||
		userHasRole(c, helpers.Uuid(cafeId), role.Hostess)
}
