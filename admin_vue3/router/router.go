// 初始化路由及注册
// @author chen
package router

import (
	"admin_vue3/api"
	"admin_vue3/config"
	_ "admin_vue3/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// 设置启动模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()

	// 跌机时恢复
	router.Use(gin.Recovery())
	// register注册
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {
	// todo 后续接口url
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api")
	{
		v1.GET("/success", api.Success)
		v1.GET("/failed", api.Failed)

		//菜单
		menu := v1.Group("sysMenu")
		sysMenu(menu)

		v1.POST("/sysRole/add", api.CreateSysRole)
		v1.GET("/sysRole/list", api.GetSysRoleList)
		v1.GET("/sysRole/info", api.GetSysRole)
		v1.PUT("/sysRole/update", api.UpdateSysRole)
		v1.DELETE("/sysRole/delete", api.DeleteSysRole)
		v1.PUT("/sysRole/updateStatus", api.UpdateSysRoleStatus)
		v1.GET("/sysRole/vo/list", api.GetSysRoleVoList)
		v1.GET("/sysRole/vo/idList", api.QueryRoleMenuIdList)
		v1.PUT("/sysRole/assignPermissions", api.AssignPermissions)
		v1.POST("/sysAdmin/add", api.CreateSysAdmin)
		v1.GET("/sysAdmin/info", api.GetSysAdmin)
		v1.GET("/sysAdmin/update", api.UpdateSysAdmin)
		v1.DELETE("/sysAdmin/delete", api.DeleteSysAdmin)
		v1.PUT("/sysAdmin/updateStatus", api.UpdateSysAdminStatus)
		v1.PUT("/sysAdmin/updatePassword", api.ResetSysAdminPassword)
		v1.GET("/sysAdmin/list", api.GetSysAdminList)
		v1.POST("/sysAdmin/login", api.Login)
	}
}
