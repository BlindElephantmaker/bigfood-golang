package controller

import (
	"bigfood/internal/infrastructure"
	_ "bigfood/swagger"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Controller struct {
	handlers *infrastructure.Handlers
}

func NewController(handlers *infrastructure.Handlers) *Controller {
	return &Controller{handlers: handlers}
}

func (controller *Controller) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("", controller.Auth)
		auth.POST("/sms-code", controller.SmsCode)
		auth.POST("/refresh-token", controller.RefreshToken)
	}

	//api := router.Group("/api", controller.userIdentity)
	//{
	//	apiUser := api.Group("/user")
	//	{
	//		// change name
	//	}
	//}

	router.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
