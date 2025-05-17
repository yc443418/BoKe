// 文章标签相关接口
// @author chen

package api

import (
	. "gin-front-api/core"
	"gin-front-api/model"
	"gin-front-api/result"
	"github.com/gin-gonic/gin"
)

// GetBTagsList 查询文章标签列表
// @Summary 查询文章标签列表
// @Tags 文章标签相关接口
// @Produce json
// @Description 查询文章标签列表
// @Success 200 {object} result.Result
// @Router /api/bTagsSelect/list [get]
func GetBTagsList(c *gin.Context) {
	var bTagsSelectVo []model.BTagsSelectVo
	Db.Table("b_tags").Order("sort").Find(&bTagsSelectVo)
	result.Success(c, bTagsSelectVo)
}
