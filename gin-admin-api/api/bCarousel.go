// 文章轮播图相关接口
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

// CreateBCarousel 新增文章轮播图
// @Summary 新增文章轮播图
// @Tags 文章轮播图相关接口
// @Produce json
// @Description 新增文章轮播图
// @Param data body model.BCarousel true "data"
// @Success 200 {object} result.Result
// @router /api/bCarousel/add [post]
// @Security ApiKeyAuth
func CreateBCarousel(c *gin.Context) {
	var dto model.BCarousel
	_ = c.BindJSON(&dto)
	addCarousel := model.BCarousel{
		Icon:       dto.Icon,
		IsPublish:  dto.IsPublish,
		CreateTime: util.HTime{Time: time.Now()},
	}
	Db.Create(&addCarousel)
	result.Success(c, true)
}

// GetBCarouselList 分页查询文章轮播图列表
// @Summary 分页查询文章轮播图列表
// @Tags 文章轮播图相关接口
// @Produce json
// @Description 分页查询文章轮播图列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param isPublish query string false "是否发布(1:否 2:是)"
// @Success 200 {object} result.Result
// @router /api/bCarousel/list [get]
// @Security ApiKeyAuth
func GetBCarouselList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	IsPublish := c.Query("isPublish")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var bCarousel []*model.BCarousel
	var count int64
	curDb := Db.Table("b_carousel")
	if IsPublish != "" {
		curDb = curDb.Where("is_publish = ?", IsPublish)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("b_carousel.create_time DESC").Find(&bCarousel)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bCarousel})
}

// GetBCarouse 根据id查询文章轮播图
// @Summary 根据id查询文章轮播图
// @Tags 文章轮播图相关接口
// @Produce json
// @Description 根据id查询文章轮播图
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/bCarousel/info [get]
// @Security ApiKeyAuth
func GetBCarouse(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var bCarousel model.BCarousel
	Db.First(&bCarousel, Id)
	result.Success(c, bCarousel)
}

// UpdateBCarouse 修改文章轮播图
// @Summary 修改文章轮播图
// @Tags 文章轮播图相关接口
// @Produce json
// @Description 修改文章轮播图
// @Param data body model.BCarousel true "data"
// @Success 200 {object} result.Result
// @router /api/bCarousel/update [put]
// @Security ApiKeyAuth
func UpdateBCarouse(c *gin.Context) {
	var dto model.BCarousel
	_ = c.BindJSON(&dto)
	var bCarousel model.BCarousel
	Db.First(&bCarousel, dto.ID)
	bCarousel.Icon = dto.Icon
	bCarousel.IsPublish = dto.IsPublish
	Db.Save(&bCarousel)
	result.Success(c, true)
}

// DeleteBCarouse 根据id删除文章轮播图
// @Summary 根据id删除文章轮播图
// @Tags 文章轮播图相关接口
// @Produce json
// @Description 根据id删除文章轮播图
// @Param data body model.BCarouselIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/bCarousel/delete [delete]
// @Security ApiKeyAuth
func DeleteBCarouse(c *gin.Context) {
	var dto model.BCarouselIdDto
	_ = c.BindJSON(&dto)
	Db.Table("b_carousel").Delete(&model.BCarousel{}, dto.Id)
	result.Success(c, true)
}
