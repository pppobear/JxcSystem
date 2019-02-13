package goods

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/filters/auth"
)

func SetGroup(rootGroup *gin.RouterGroup) {
	goodsGroup := rootGroup.Group("/goods").Use(auth.Middleware("jwt"))
	{
		goodsGroup.GET("", ListGoods)
		goodsGroup.POST("", Create)
		goodsGroup.GET("/:id", RetrieveGoods)
		goodsGroup.DELETE("/:id", Delete)
		goodsGroup.PUT("/:id", UpdateGood)
	}

	inventoryGroup := rootGroup.Group("/inventory")
	{
		inventoryGroup.GET("", ListInventories)
		inventoryGroup.GET("/:id", RetrieveInventory)
		inventoryGroup.PUT("/:id", UpdateInventory)
	}
}
