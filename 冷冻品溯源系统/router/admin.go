package router

import (
	"Traceability/api/v1/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAdminRoutes(r *gin.Engine, db *gorm.DB) {
	adminController := controller.NewAdminController(db)

	adminGroup := r.Group("/admin")
	//adminGroup.Use(middleware.JwtAuth(), middleware.AdminOnly()) // 权限拦截
	{
		adminGroup.GET("/users", adminController.GetUserList)
		adminGroup.GET("/pending", adminController.GetPendingRequests)
		adminGroup.PUT("/:log_id", adminController.ProcessAudit)
		adminGroup.PUT("/product/:log_id", adminController.ProcessProducts)
	}
}
