// 文章分类相关接口
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

// CreateBCategory 新增文章分类
// @Summary 新增文章分类
// @Tags 文章分类相关接口
// @Produce json
// @Description 新增文章分类
// @Param data body model.BCategory true "data"
// @Success 200 {object} result.Result
// @router /api/bCategory/add [post]
// @Security ApiKeyAuth
func CreateBCategory(c *gin.Context) {
	var dto model.BCategory
	_ = c.BindJSON(&dto)
	category := GetCategoryByName(dto.CategoryName)
	if category.ID > 0 {
		result.Failed(c, int(result.ApiCode.BCategoryIsExist), result.ApiCode.GetMessage(result.ApiCode.BCategoryIsExist))
		return
	}
	addCategory := model.BCategory{
		CategoryName: dto.CategoryName,
		Sort:         dto.Sort,
		CreateTime:   util.HTime{Time: time.Now()},
	}
	Db.Create(&addCategory)
	result.Success(c, true)
}

// GetCategoryList 分页查询文章分类列表
// @Summary 分页查询文章分类列表
// @Tags 文章分类相关接口
// @Produce json
// @Description 分页查询文章分类列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param categoryName query string false "分类名称"
// @Success 200 {object} result.Result
// @router /api/bCategory/list [get]
// @Security ApiKeyAuth
func GetCategoryList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	CategoryName := c.Query("categoryName")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var bCategoryVo []*model.BCategoryVo
	var count int64
	curDb := Db.Table("b_category bc").
		Select("bc.*, ba.quantity").
		Joins("LEFT JOIN (SELECT category_id, count(id) AS quantity FROM b_article GROUP BY category_id) ba ON bc.id = ba.category_id")
	if CategoryName != "" {
		curDb = curDb.Where("bc.category_name = ?", CategoryName)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("bc.sort DESC").Find(&bCategoryVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bCategoryVo})
}

// GetBCategory 根据id查询文章分类
// @Summary 根据id查询文章分类
// @Tags 文章分类相关接口
// @Produce json
// @Description 根据id查询文章分类
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/bCategory/info [get]
// @Security ApiKeyAuth
func GetBCategory(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var bCategory model.BCategory
	Db.First(&bCategory, Id)
	result.Success(c, bCategory)
}

// UpdateBCategory 修改文章分类
// @Summary 修改文章分类
// @Tags 文章分类相关接口
// @Produce json
// @Description 修改文章分类
// @Param data body model.BCategory true "data"
// @Success 200 {object} result.Result
// @router /api/bCategory/update [put]
// @Security ApiKeyAuth
func UpdateBCategory(c *gin.Context) {
	var dto model.BCategory
	_ = c.BindJSON(&dto)
	var bCategory model.BCategory
	Db.First(&bCategory, dto.ID)
	bCategory.CategoryName = dto.CategoryName
	bCategory.Sort = dto.Sort
	Db.Save(&bCategory)
	result.Success(c, true)
}

// DeleteBCategory 根据id删除文章分类
// @Summary 根据id删除文章分类
// @Tags 文章分类相关接口
// @Produce json
// @Description 根据id删除文章分类
// @Param data body model.BCategoryIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/bCategory/delete [delete]
// @Security ApiKeyAuth
func DeleteBCategory(c *gin.Context) {
	var dto model.BCategoryIdDto
	_ = c.BindJSON(&dto)
	article := GetCategoryArticle(dto.Id)
	if article.ID > 0 {
		result.Failed(c, int(result.ApiCode.CategoryHasArticle), result.ApiCode.GetMessage(result.ApiCode.CategoryHasArticle))
		return
	}
	Db.Table("b_category").Delete(&model.BCategory{}, dto.Id)
	result.Success(c, true)
}

// GetBCategorySelectList 查询文章分类列表
// @Summary 查询文章分类列表
// @Tags 文章分类相关接口
// @Produce json
// @Description 查询文章分类列表
// @Success 200 {object} result.Result
// @router /api/bCategorySelect/list [get]
// @Security ApiKeyAuth
func GetBCategorySelectList(c *gin.Context) {
	var bCategorySelectVo []model.BCategorySelectVo
	Db.Table("b_category").Order("sort").Find(&bCategorySelectVo)
	result.Success(c, bCategorySelectVo)
}

// GetCategoryByName 根据分类名称查询
func GetCategoryByName(categoryName string) (bCategory model.BCategory) {
	Db.Where("category_name = ?", categoryName).First(&bCategory)
	return bCategory
}

// GetCategoryArticle 根据分类id查询该分类是否有文章
func GetCategoryArticle(categoryId uint) (article model.BArticle) {
	Db.Where("category_id = ?", categoryId).First(&article)
	return article
}
