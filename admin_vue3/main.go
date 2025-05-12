// 启动程序
// @author chen
package main

import (
	"admin_vue3/config"
	"admin_vue3/core"
	_ "admin_vue3/docs"
	"admin_vue3/global"
	"admin_vue3/router"
	"fmt"
)

// @title 博客运营后台
// @version 1.0
// @description 博客运营后台API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 初始化logger
	global.Log = core.InitLogger()

	// 初始化mysql
	core.MysqlInit()

	// 初始化redis
	core.RedisInit()

	// 初始化路由
	router := router.InitRouter()
	address := fmt.Sprintf("%s:%d", config.Config.System.Host, config.Config.System.Port)
	global.Log.Infof("系统启动成功，运行在:%s", address)
	router.Run(address)
}
