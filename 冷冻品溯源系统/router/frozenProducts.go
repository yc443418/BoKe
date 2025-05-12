package router

import (
	api "Traceability/api/v1/controller"
	"github.com/gin-gonic/gin"
)

func frozenProducts(p *gin.RouterGroup) {
	//需要权限管理中间件
	//这个是后台管理员查看所有冷冻列表的路由
	p.GET("", api.GetProductListByAdmin)

	//产家查看自己的列表
	p.GET("/:factoryId", api.GetProductListByFactory)

	//获取冷冻品详情
	p.GET("/detail/:productId", api.GetProductDetail)
	p.PUT("/detail/:productId", api.PutProductDetail)
	p.POST("/Add", api.AddProduct)

}
