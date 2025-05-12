package router

import (
	"Traceability/api/v1/controller"
	"Traceability/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	userController := controller.NewUserController(db)

	// 无需认证的路由
	authGroup := r.Group("")
	{
		authGroup.POST("/register", userController.Register)
		authGroup.POST("/login", userController.Login)
		authGroup.POST("/user/type-request", userController.SubmitUserTypeRequest)
	}

	//需要认证的路由
	authRequired := r.Group("/")
	authRequired.Use(middleware.JwtAuth())
	{
		authRequired.GET("/user/info", userController.GetUserInfo)
	}
}
