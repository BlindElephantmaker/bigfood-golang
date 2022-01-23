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
// @Param        input  body      sendSmsCode.Message  true  "json"
// @Success      200    {object}  SmsCodeResponse
// @Failure      400    {object}  server.ResponseError  "Phone number is invalid"
// @Failure      429    {object}  server.ResponseError  "Retry count of sms code requests exceeded"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /auth/sms-code [post]
func (controller *Controller) SmsCode(c *gin.Context) {
	var message sendSmsCode.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	err = controller.handlers.SendSmsCode.Run(message)
	if err == user.ErrorPhoneNumberIsInvalid {
		server.NewResponseError(c, http.StatusBadRequest, err)
		return
	}
	if err == sendSmsCode.ErrorRetryCountExceeded {
		server.NewResponseError(c, http.StatusTooManyRequests, err)
		return
	}
	if err != nil {
		server.NewResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &SmsCodeResponse{true})
}
