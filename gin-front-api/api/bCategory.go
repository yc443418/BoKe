// 文章分类相关接口
// @author chen

package api

import (
	. "gin-front-api/core"
	"gin-front-api/model"
	"gin-front-api/result"
	"github.com/gin-gonic/gin"
)

// GetBCategoryList 查询文章分类列表
// @Summary 查询文章分类列表
// @Tags 文章分类相关接口
// @Produce json
// @Description 查询文章分类列表
// @Success 200 {object} result.Result
// @Router /api/bCategorySelect/list [get]
func GetBCategoryList(c *gin.Context) {
	var bCategorySelectVo []model.BCategorySelectVo
	Db.Table("b_category").Order("sort").Find(&bCategorySelectVo)
	result.Success(c, bCategorySelectVo)
}
