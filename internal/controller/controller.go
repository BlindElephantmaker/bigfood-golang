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

	auth := router.Group("auth")
	{
		auth.POST("", controller.auth)
		auth.POST("sms-code", controller.smsCode)
		auth.PUT("refresh-token", controller.refreshToken)
	}

	api := router.Group("/", controller.userIdentity)
	{
		apiUser := api.Group("user")
		{
			apiUser.PUT("", controller.userEdit)
		}

		apiCafe := api.Group("cafe", controller.userIdentity)
		{
			apiCafe.POST("", controller.cafeCreate)

			apiCafeUser := apiCafe.Group("user")
			{
				apiCafeUser.GET("list", controller.cafeUserList)
				apiCafeUser.POST("", controller.cafeUserCreate)
				apiCafeUser.DELETE("", controller.cafeUserDelete)
				apiCafeUser.PUT("", controller.cafeUserEdit)
			}
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

		apiReserve := api.Group("reserve")
		{
			apiReserve.GET("/:reserve-id", controller.reserveGet)
			apiReserve.POST("/", controller.reserveCreate)
			apiReserve.DELETE("/", controller.reserveDelete)
			apiReserve.PUT("/undelete", controller.reserveUndelete)
		}
	}

	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
