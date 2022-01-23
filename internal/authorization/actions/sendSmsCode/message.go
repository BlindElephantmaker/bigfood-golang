package sendSmsCode

type Message struct {
	Phone string `json:"phone" binding:"required" example:"+71234567890"`
	// todo: captcha
}
