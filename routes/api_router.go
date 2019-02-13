package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"pppobear.cn/jxc-backend/controllers/customer"
	"pppobear.cn/jxc-backend/controllers/goods"
	"pppobear.cn/jxc-backend/controllers/purchases"
	"pppobear.cn/jxc-backend/controllers/sales"
	"pppobear.cn/jxc-backend/controllers/staff"
	"pppobear.cn/jxc-backend/controllers/user"
)

func RegisterApiRouter(router *gin.Engine) {
	setup(router)

	apiRouter := router.Group("/api/v1")
	{
		purchases.SetGroup(apiRouter)
		sales.SetGroup(apiRouter)
		goods.SetGroup(apiRouter)
		customer.SetGroup(apiRouter)
		staff.SetGroup(apiRouter)
		user.SetGroup(apiRouter)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
