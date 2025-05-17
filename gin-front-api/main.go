// 程序启动入口
// @author chen

package main

import (
	"fmt"
	"gin-front-api/config"
	"gin-front-api/core"
	. "gin-front-api/core"
	_ "gin-front-api/docs"
	"gin-front-api/global"
	"gin-front-api/middleware"
	"gin-front-api/router"
)

// @title 博客用户后台
// @version 1.0
// @description 博客用户后台API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化mysql
	core.MysqlInit()
	// 创建数据库表结构
	if err := middleware.AutoMigrate(Db); err != nil {
		panic(fmt.Sprintf("数据库迁移失败: %v", err))
	}
	// 初始化redis
	core.RedisInit()
	// 初始化路由
	router := router.InitRouter()
	address := fmt.Sprintf("%s:%s", config.Config.System.Host, config.Config.System.Port)
	global.Log.Infof("系统启动成功，运行在: %s", address)
	router.Run(address)
}
