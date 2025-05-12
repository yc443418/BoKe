package router

import (
	"Traceability/global"
	"Traceability/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"time"
)

// 初始化路由
func InitRouter(e *gin.Engine, db *gorm.DB) {

	runAddress := fmt.Sprintf("%s:%d", global.CONFIG.Server.Address, global.CONFIG.Server.Port)

	// 添加 Swagger 路由
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//设置登录和注册路由，该路由无需jwt验证
	SetupUserRoutes(e, db)

	//需要jwt的路由分组
	needgroup := e.Group("")
	needgroup.Use(middleware.JwtAuth()) // 启用JWT中间件
	go need_Jwt_Group(needgroup, e, db)

	time.Sleep(100 * time.Millisecond)
	e.Run(runAddress)
}

func need_Jwt_Group(needgroup *gin.RouterGroup, e *gin.Engine, db *gorm.DB) {
	// 添加管理员路由
	SetupAdminRoutes(e, db)

	//冷冻产品管理
	frozen_products := needgroup.Group("/frozen_products")
	frozenProducts(frozen_products)
}
