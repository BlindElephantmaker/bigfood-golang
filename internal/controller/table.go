package controller

import "bigfood/internal/table"

type TableListResponse struct {
	Tables []*table.Table `json:"tables"`
}
