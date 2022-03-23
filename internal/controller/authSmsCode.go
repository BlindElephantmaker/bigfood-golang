package controller

import (
	"bigfood/internal/authorization/action/sendSmsCode"
	"bigfood/pkg/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// smsCode
// @Summary      Send SMS code
// @Description  Send SMS code to user at authorization
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body  sendSmsCode.Message  true  "Body"
// @Success      200    "Success"
// @Failure      400    {object}  server.ResponseError  "Phone number is invalid"
// @Failure      429    {object}  server.ResponseError  "Retry count of sms code requests exceeded"
// @Failure      500    {object}  server.ResponseError  "Internal Server Error"
// @Router       /auth/sms-code [post]
func (controller *Controller) smsCode(c *gin.Context) {
	var message sendSmsCode.Message
	err := server.ParseJsonRequestToMessage(c, &message)
	if err != nil {
		return
	}

	err = controller.handlers.SendSmsCode.Run(message)
	if err == sendSmsCode.ErrorRetryCountExceeded {
		server.NewResponseError(c, http.StatusTooManyRequests, err)
		return
	}
	if err != nil {
		server.InternalServerError(c, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
