package controller

import (
	"bigfood/internal/authorization/actions/sendSmsCode"
	"bigfood/internal/user"
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
		// todo: message
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation
		return
	}

	err := controller.handlers.SendSmsCode.Run(message)
	if err == user.ErrorPhoneNumberIsInvalid {
		server.NewResponseError(c, http.StatusBadRequest, err) // todo: annotation
		return
	}
	if err == sendSmsCode.ErrorRetryCountExceeded {
		server.NewResponseError(c, http.StatusTooManyRequests, err) // todo: annotation
		return
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err) // todo: annotation
		return
	}

	c.JSON(http.StatusOK, &SmsCodeResponse{true})
}
