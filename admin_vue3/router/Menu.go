package router

import (
	"admin_vue3/api"

	"github.com/gin-gonic/gin"
)

func sysMenu(menu *gin.RouterGroup) {
	menu.POST("/add", api.CreateSysMenu)
	menu.GET("/list", api.GetSysMenuList)
	menu.GET("/info", api.GetSysMenu)
	menu.PUT("/update", api.UpdateSysMenu)
	menu.DELETE("/delete", api.DeleteSysMenu)
}
