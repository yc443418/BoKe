// 标签相关接口
// @author chen

package api

import (
	. "gin-admin-api/core"
	"gin-admin-api/model"
	"gin-admin-api/result"
	util "gin-admin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// CreateBTags 新增文章标签
// @Summary 新增文章标签
// @Tags 文章标签相关接口
// @Produce json
// @Description 新增文章标签
// @Param data body model.BTags true "data"
// @Success 200 {object} result.Result
// @router /api/bTags/add [post]
// @Security ApiKeyAuth
func CreateBTags(c *gin.Context) {
	var dto model.BTags
	_ = c.BindJSON(&dto)
	bTags := GetBTagsByName(dto.TagName)
	if bTags.ID > 0 {
		result.Failed(c, int(result.ApiCode.BTagsHasExist), result.ApiCode.GetMessage(result.ApiCode.BTagsHasExist))
		return
	}
	addBTags := model.BTags{
		TagName:    dto.TagName,
		Sort:       dto.Sort,
		CreateTime: util.HTime{Time: time.Now()},
	}
	Db.Create(&addBTags)
	result.Success(c, true)
}

// GetBTagsList 分页查询文章标签列表
// @Summary 分页查询文章标签列表
// @Tags 文章标签相关接口
// @Produce json
// @Description 分页查询文章标签列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param tagName query string false "标签名称"
// @Success 200 {object} result.Result
// @router /api/bTags/list [get]
// @Security ApiKeyAuth
func GetBTagsList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	TagName := c.Query("tagName")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var bTagsVo []*model.BTagsVo
	var count int64
	curDb := Db.Table("b_tags bt").
		Select("bt.*, ba.quantity").
		Joins("LEFT JOIN (SELECT tag_id,count(id) AS quantity FROM b_article GROUP BY tag_id) ba ON bt.id = ba.tag_id")
	if TagName != "" {
		curDb = curDb.Where("bt.tag_name = ?", TagName)
	}
	curDb.Count(&count).Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("bt.create_time DESC").Find(&bTagsVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bTagsVo})
}

// GetBTags 根据id查询文章标签
// @Summary 根据id查询文章标签
// @Tags 文章标签相关接口
// @Produce json
// @Description 根据id查询文章标签
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/bTags/info [get]
// @Security ApiKeyAuth
func GetBTags(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var bTags model.BTags
	Db.First(&bTags, Id)
	result.Success(c, bTags)
}

// UpdateBTags 修改文章标签
// @Summary 修改文章标签
// @Tags 文章标签相关接口
// @Produce json
// @Description 修改文章标签
// @Param data body model.BTags true "data"
// @Success 200 {object} result.Result
// @router /api/bTags/update [put]
// @Security ApiKeyAuth
func UpdateBTags(c *gin.Context) {
	var dto model.BTags
	_ = c.BindJSON(&dto)
	var bTags model.BTags
	Db.First(&bTags, dto.ID)
	bTags.TagName = dto.TagName
	bTags.Sort = dto.Sort
	Db.Save(&bTags)
	result.Success(c, true)
}

// DeleteBTags 根据id删除文章标签
// @Summary 根据id删除文章标签
// @Tags 文章标签相关接口
// @Produce json
// @Description 根据id删除文章标签
// @Param data body model.BTagsIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/bTags/delete [delete]
// @Security ApiKeyAuth
func DeleteBTags(c *gin.Context) {
	var dto model.BTagsIdDto
	_ = c.BindJSON(&dto)
	Db.Table("b_tags").Delete(&model.BTags{}, dto.Id)
	result.Success(c, true)
}

// GetBTagsSelectList 查询文章标签列表
// @Summary 查询文章标签列表
// @Tags 文章标签相关接口
// @Produce json
// @Description 查询文章标签列表
// @Success 200 {object} result.Result
// @router /api/bTagsSelect/list [get]
// @Security ApiKeyAuth
func GetBTagsSelectList(c *gin.Context) {
	var bTagsSelectVo []model.BTagsSelectVo
	Db.Table("b_tags").Order("sort").Find(&bTagsSelectVo)
	result.Success(c, bTagsSelectVo)
}

// GetBTagsByName 根据名称查询
func GetBTagsByName(name string) (bTags model.BTags) {
	Db.Where("name = ?", name).First(&bTags)
	return bTags
}
