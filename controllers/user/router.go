package user

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/filters/auth"
)

func SetGroup(rootGroup *gin.RouterGroup) {
	group := rootGroup.Group("/auth")
	group.POST("/login", Login)
	group.Any("/logout", Logout).Use(auth.Middleware("jwt"))
}
