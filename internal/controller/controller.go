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
				apiCafeUser.GET("list/:cafe-id", controller.cafeUserList)
				apiCafeUser.POST("", controller.cafeUserCreate)
				apiCafeUser.PUT("", controller.cafeUserEdit)
				apiCafeUser.DELETE("", controller.cafeUserDelete)
			}
		}

		apiTable := api.Group("table")
		{
			apiTable.GET("/list/:cafe-id", controller.tableList)
			apiTable.GET("/list/:cafe-id/available", controller.tableListAvailable)
			apiTable.POST("/", controller.tableCreate)
			apiTable.POST("/create-mass", controller.tableCreateMass)
			apiTable.PUT("/", controller.tableEdit)
			apiTable.DELETE("/", controller.tableDelete)
			apiTable.DELETE("/delete-all", controller.tableDeleteAll)
			apiTable.GET(":table-id", controller.tableGet)
		}

		apiReserve := api.Group("reserve")
		{
			apiReserve.GET("/:reserve-id", controller.reserveGet)
			apiReserve.GET("/table/:table-id", controller.reserveListByTable)
			apiReserve.GET("/table/:table-id/history", controller.reserveHistoryByTable)
			apiReserve.POST("/", controller.reserveCreate)
			apiReserve.PUT("/", controller.reserveEdit)
			apiReserve.PUT("/undelete", controller.reserveUndelete)
			apiReserve.DELETE("/", controller.reserveDelete)
		}
	}

	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
