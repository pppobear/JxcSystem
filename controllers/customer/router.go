package customer

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/filters/auth"
)

func SetGroup(rootGroup *gin.RouterGroup) {
	group := rootGroup.Group("/customer").Use(auth.Middleware("jwt"))
	{
		group.GET("", List)
		group.POST("", Create)
		group.GET("/:id", Retrieve)
		group.DELETE("/:id", Delete)
		group.PUT("/:id", Update)
	}
}
