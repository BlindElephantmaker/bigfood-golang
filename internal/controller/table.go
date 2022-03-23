package controller

import (
	"bigfood/internal/cafeUser"
	"bigfood/internal/helpers"
	"bigfood/internal/table"
	"github.com/gin-gonic/gin"
)

type TableListResponse struct {
	Tables []*table.Table `json:"tables"`
}

func userIsAdmin(c *gin.Context, cafeId helpers.Uuid) bool {
	return userHasRole(c, cafeId, cafeUser.Admin)
}

func userIsHostess(c *gin.Context, cafeId helpers.Uuid) bool {
	return userIsAdmin(c, cafeId) ||
		userHasRole(c, cafeId, cafeUser.Hostess)
}
