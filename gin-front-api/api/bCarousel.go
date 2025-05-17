// 文章轮播图相关接口
// @author chen

package api

import (
	. "gin-front-api/core"
	"gin-front-api/model"
	"gin-front-api/result"
	"github.com/gin-gonic/gin"
)

// GetBCarouselList 查询文章轮播图列表
// @Summary 查询文章轮播图列表
// @Tags 文章轮播图相关接口
// @Produce json
// @Description 查询文章轮播图列表
// @Success 200 {object} result.Result
// @Router /api/bCarousel/list [get]
func GetBCarouselList(c *gin.Context) {
	var bCarousel []model.BCarousel
	Db.Table("b_carousel").Order("create_time").Where("is_publish = ?", 2).Find(&bCarousel)
	result.Success(c, bCarousel)
}
