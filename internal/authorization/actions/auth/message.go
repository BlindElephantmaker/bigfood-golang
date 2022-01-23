package auth

type Message struct {
	Phone   string `json:"phone" binding:"required" example:"+71234567890"`
	SmsCode string `json:"sms-code" binding:"required" example:"1234"`
}
