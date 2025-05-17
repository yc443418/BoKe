// 启动程序
// @author chen

package main

import (
	"fmt"
	"gin-admin-api/config"
	"gin-admin-api/core"
	. "gin-admin-api/core"
	_ "gin-admin-api/docs"
	"gin-admin-api/global"
	"gin-admin-api/middleware"
	"gin-admin-api/router"
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
	// 创建数据库表结构
	if err := middleware.AutoMigrate(Db); err != nil {
		panic(fmt.Sprintf("数据库迁移失败: %v", err))
	}

	// 数据库数据初始化
	if err := middleware.SeedData(Db); err != nil {
		panic(fmt.Sprintf("数据初始化失败: %v", err))
	}
	// 初始化redis
	core.RedisInit()
	// 初始化路由
	router := router.InitRouter()
	address := fmt.Sprintf("%s:%d", config.Config.System.Host, config.Config.System.Port)
	global.Log.Infof("系统启动成功，运行在: %s", address)
	// 打印token配置
	global.Log.Infof("token配置: %s:%d:%s:%s", config.Config.Token.Headers, config.Config.Token.ExpireTime, config.Config.Token.Secret, config.Config.Token.Issuer)
	router.Run(address)
}
