package tableDeleteAll

type Message struct {
	CafeId string `json:"cafe-id" binding:"required" example:"uuid"`
}
