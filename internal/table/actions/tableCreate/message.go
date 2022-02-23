package tableCreate

type Message struct {
	CafeId  string  `json:"cafe-id" binding:"required" example:"uuid"`
	Title   *string `json:"title"`
	Comment *string `json:"comment"`
	Seats   *int    `json:"seats"`
}
