package main

import (
	_ "Traceability/docs"
	"Traceability/global"
	"Traceability/initialize"
	"time"
)

// @title 冷冻品溯源系统
// @version 1.0
// @description 冷冻品溯源系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	go initialize.MustConfig()
	time.Sleep(1 * time.Second)
	initialize.MustLoadGorm()
	initialize.AutoMigrate(global.DB)
	initialize.InitRedis()
	initialize.MustRunWebServer(global.DB)
	defer initialize.CloseRedis()
}
