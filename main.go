package main

import (
	"runtime"

	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/config"
	_ "pppobear.cn/jxc-backend/docs"
	_ "pppobear.cn/jxc-backend/module/logger" // 日志
	"pppobear.cn/jxc-backend/module/server"
	routeRegister "pppobear.cn/jxc-backend/routes"
)

// @title 进销存API
// @version 0.0.1
// @description  进销存项目的API文档
// @BasePath /api/v1/

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config.Setup()
	router := gin.New()

	// 注册路由
	routeRegister.RegisterApiRouter(router)

	server.Run(router)
}
