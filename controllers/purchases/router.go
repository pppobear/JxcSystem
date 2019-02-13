package purchases

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/filters/auth"
)

func SetGroup(rootGroup *gin.RouterGroup) {
	group := rootGroup.Group("/purchases").Use(auth.Middleware("jwt"))
	{
		group.GET("", List)
		group.POST("", Create)
		group.GET("/:id", Retrieve)
		group.PUT("/:id", Update)
		group.DELETE("/:id", Delete)
		group.POST("/:id", AddDetails)
		group.PUT("/:id/:gid", UpdateDetail)
		group.DELETE("/:id/:gid", DeleteDetail)
	}
}
