package createMass

type Message struct {
	CafeId   string `json:"cafe-id" binding:"required" example:"uuid"`
	Quantity int    `json:"quantity" binding:"required" example:"10"`
}
