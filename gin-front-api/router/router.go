// 初始化路由及注册
// @author chen

package router

import (
	"gin-front-api/api"
	"gin-front-api/config"
	"gin-front-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// InitRouter 初始化路由
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
	// register注入
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/captcha", api.Captcha)
	router.POST("/api/qqMail", api.SendMail)
	router.POST("/api/register", api.Register)
	router.POST("/api/login", api.Login)
	router.POST("/api/user/reset/password", api.ResetPassword)
	router.GET("/api/bCategorySelect/list", api.GetBCategoryList)
	router.GET("/api/bTagsSelect/list", api.GetBTagsList)
	router.GET("/api/bCarousel/list", api.GetBCarouselList)
	router.GET("/api/bArticle/list", api.GetBArticleVoList)
	router.GET("/api/user/article/rank/list", api.GetArticleUserRankList)
	router.GET("/api/bArticle/detail", api.GetBArticleDetail)
	router.GET("/api/bArticle/keyword/list", api.GetBArticleListByKeywords)
	router.GET("/api/bArticle/commentVo/list", api.GetBArticleCommentVoList)
	router.GET("/api/user/articleUserDetail", api.GetArticleUserDetail)
	router.GET("/api/user/bArticle/list", api.GetArticleUserList)
	router.GET("/api/user/message/count", api.GetMessageCount)
	router.GET("/api/user/message/list", api.GetUserMessageList)
	jwt := router.Group("/api", middleware.AuthMiddleware())
	{
		jwt.POST("/upload", api.Upload)
		jwt.POST("/bArticle/add", api.CreateBArticle)
		jwt.POST("/bArticle/like", api.DoBArticleLike)
		jwt.GET("/bArticle/articleId", api.GetBArticleByArticleId)
		jwt.PUT("/bArticle/update", api.UpdateBArticle)
		jwt.POST("/bArticle/add/comment", api.CreateBArticleComment)
		jwt.POST("/bArticle/comment/top", api.DoCommentTop)
		jwt.POST("/bArticle/comment/like", api.DoBArticleCommentLike)
		jwt.PUT("/user/updateUserInfo", api.UpdateUserInfo)
		jwt.PUT("/user/updateUserPassword", api.UpdateUserPassword)
	}
}
