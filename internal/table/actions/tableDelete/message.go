package tableDelete

type Message struct {
	TableId string `json:"table-id" binding:"required" example:"uuid"`
}
