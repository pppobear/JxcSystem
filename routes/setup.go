package routes

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/config"
	"pppobear.cn/jxc-backend/filters"
	"pppobear.cn/jxc-backend/filters/auth"
	"pppobear.cn/jxc-backend/module/handler"
)

func setup(router *gin.Engine) {
	if config.GetEnv().Debug {
		pprof.Register(router) // 性能分析工具
	}
	router.Use(gin.Logger())

	// 错误处理
	router.Use(handler.HandleErrors())
	// 全局session
	router.Use(filters.RegisterSession())
	// 全局cache
	router.Use(filters.RegisterCache())
	// 全局auth cookie
	router.Use(auth.RegisterGlobalAuthDriver("cookie", "web_auth"))
	// 全局auth jwt
	router.Use(auth.RegisterGlobalAuthDriver("jwt", "jwt_auth"))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "不支持该请求方法",
		})
		return
	})
}
