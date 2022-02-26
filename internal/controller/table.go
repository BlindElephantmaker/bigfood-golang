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

func userIsAdmin(c *gin.Context, cafeId string) bool {
	return userHasRole(c, helpers.Uuid(cafeId), cafeUser.Admin)
}

func userIsHostess(c *gin.Context, cafeId string) bool {
	return userIsAdmin(c, cafeId) ||
		userHasRole(c, helpers.Uuid(cafeId), cafeUser.Hostess)
}
