package tableEdit

type Message struct {
	TableId string  `json:"table-id" binding:"required" example:"uuid"`
	Title   *string `json:"title"` // todo: check *
	Comment *string `json:"comment"`
	Seats   *int    `json:"seats"`
}
