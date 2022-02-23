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
		auth.POST("/", controller.auth)
		auth.POST("/sms-code", controller.smsCode)
		auth.PUT("/refresh-token", controller.refreshToken)
		auth.DELETE("/logout", controller.userIdentity, controller.logout)
	}

	api := router.Group("/", controller.userIdentity)
	{
		apiUser := api.Group("user")
		{
			apiUser.PUT("/", controller.userEdit)
		}

		apiCafe := api.Group("cafe", controller.userIdentity)
		{
			apiCafe.POST("/", controller.cafeCreate)
		}

		apiTable := api.Group("table")
		{
			apiTable.GET("/list", controller.tableGetList)
			apiTable.PUT("/", controller.tableEdit)
			apiTable.POST("/", controller.tableCreate)
			apiTable.POST("/create-mass", controller.tableCreateMass)
			apiTable.DELETE("/", controller.tableDelete)
			apiTable.DELETE("/delete-all", controller.tableDeleteAll)
		}
	}

	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
