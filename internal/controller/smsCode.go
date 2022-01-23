package controller

import (
	"bigfood/internal/authorization/actions/sendSmsCode"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SmsCodeResponse struct {
	Success bool `json:"success" example:"true"`
}

// SmsCode
// @Summary      Send SMS code
// @Description  Send SMS code to user at authorization
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      sendSmsCode.Message  true  "Phone number"
// @Success      200    {object}  SmsCodeResponse
// @Failure      500    {object}  server.ResponseError
// @Router       /auth/sms-code [post]
func (controller *Controller) SmsCode(c *gin.Context) {
	var message sendSmsCode.Message
	if err := c.BindJSON(&message); err != nil {
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation
		return
	}

	err := controller.handlers.SendSmsCode.Run(message)
	if err != nil {
		// todo: how handle errors by type and user friendly messages?
		server.NewResponseError(c, http.StatusInternalServerError, err) // todo: change code response
		return
	}

	c.JSON(http.StatusOK, &SmsCodeResponse{true})
}
