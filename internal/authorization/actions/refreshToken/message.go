package refreshToken

type Message struct {
	Token string `json:"token" binding:"required" example:"UUID"`
}
