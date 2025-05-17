// 初始化路由及注册
// @author chen

package router

import (
	"gin-admin-api/api"
	"gin-admin-api/config"
	"gin-admin-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 设置启动模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	// 跌机时恢复
	router.Use(gin.Recovery())
	// 跨域
	router.Use(middleware.Cors())
	// 上传配置
	router.StaticFS(config.Config.UploadSettings.UploadDir, http.Dir(config.Config.UploadSettings.UploadDir))
	// register 注册
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/sysAdmin/login", api.Login)
	jwt := router.Group("/api", middleware.AuthMiddleware())
	{
		jwt.GET("/success", api.Success)
		jwt.GET("/failed", api.Failed)
		jwt.POST("/sysMenu/add", api.CreateSysMenu)
		jwt.GET("/sysMenu/list", api.GetSysMenuList)
		jwt.GET("/sysMenu/info", api.GetSysMenu)
		jwt.PUT("//sysMenu/update", api.UpdateSysMenu)
		jwt.DELETE("/sysMenu/delete", api.DeleteSysMenu)
		jwt.POST("/sysRole/add", api.CreateSysRole)
		jwt.GET("/sysRole/list", api.GetSysRoleList)
		jwt.GET("/sysRole/info", api.GetSysRole)
		jwt.PUT("/sysRole/update", api.UpdateSysRole)
		jwt.DELETE("/sysRole/delete", api.DeleteSysRole)
		jwt.PUT("/sysRole/updateStatus", api.UpdateSysRoleStatus)
		jwt.GET("/sysRole/vo/list", api.GetSysRoleVoList)
		jwt.GET("/sysRole/vo/idList", api.QueryRoleMenuIdList)
		jwt.PUT("/sysRole/assignPermissions", api.AssignPermissions)
		jwt.POST("/sysAdmin/add", api.CreateSysAdmin)
		jwt.GET("/sysAdmin/info", api.GetSysAdmin)
		jwt.PUT("/sysAdmin/update", api.UpdateSysAdmin)
		jwt.DELETE("/sysAdmin/delete", api.DeleteSysAdmin)
		jwt.PUT("/sysAdmin/updateStatus", api.UpdateSysAdminStatus)
		jwt.PUT("/sysAdmin/updatePassword", api.ResetSysAdminPassword)
		jwt.GET("/sysAdmin/list", api.GetSysAdminList)
		jwt.POST("/upload", api.Upload)
		jwt.PUT("/sysAdmin/updatePersonal", api.UpdatePersonal)
		jwt.PUT("/sysAdmin/updatePersonalPassword", api.UpdatePersonalPassword)
		jwt.POST("/sysConfig/add", api.CreateSysConfig)
		jwt.GET("/sysConfig/list", api.GetSysConfigList)
		jwt.GET("/sysConfig/info", api.GetSysConfig)
		jwt.PUT("/sysConfig/update", api.UpdateSysConfig)
		jwt.DELETE("/sysConfig/delete", api.DeleteSysConfig)
		jwt.DELETE("/sysConfig/refresh", api.Refresh)
		jwt.POST("/bCarousel/add", api.CreateBCarousel)
		jwt.GET("/bCarousel/list", api.GetBCarouselList)
		jwt.GET("/bCarousel/info", api.GetBCarouse)
		jwt.PUT("/bCarousel/update", api.UpdateBCarouse)
		jwt.DELETE("/bCarousel/delete", api.DeleteBCarouse)
		jwt.POST("/bTags/add", api.CreateBTags)
		jwt.GET("/bTags/list", api.GetBTagsList)
		jwt.GET("/bTags/info", api.GetBTags)
		jwt.PUT("/bTags/update", api.UpdateBTags)
		jwt.DELETE("/bTags/delete", api.DeleteBTags)
		jwt.GET("/bTagsSelect/list", api.GetBTagsSelectList)
		jwt.POST("/bCategory/add", api.CreateBCategory)
		jwt.GET("/bCategory/list", api.GetCategoryList)
		jwt.GET("/bCategory/info", api.GetBCategory)
		jwt.PUT("/bCategory/update", api.UpdateBCategory)
		jwt.DELETE("/bCategory/delete", api.DeleteBCategory)
		jwt.GET("/bCategorySelect/list", api.GetBCategorySelectList)
		jwt.GET("/bArticle/list", api.GetBArticleList)
		jwt.PUT("/bArticle/updateTopType", api.UpdateToType)
		jwt.PUT("/bArticle/updateStatus", api.UpdateStatus)
		jwt.GET("/bArticle/comment", api.GetBArticleCommentByArticleId)
		jwt.GET("/bArticle/comment/list", api.GetBArticleCommentList)
		jwt.PUT("/bArticle/comment/updateStatus", api.UpdateArticleCommentStatus)
		jwt.GET("/user/list", api.GetUserList)
		jwt.PUT("/user/updateStatus", api.UpdateUserStatus)
		jwt.POST("/user/send/message", api.SendMessage)
		jwt.GET("/index/statistics/list", api.GetDataStatisticsList)
		jwt.GET("/index/data/article/create/list", api.GetDataArticleCreateList)
		jwt.GET("/index/data/user/create/list", api.GetDataUserCreateList)
	}
}
