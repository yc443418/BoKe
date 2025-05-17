// 文章相关接口
// @author chen

package api

import (
	"gin-admin-api/constant"
	. "gin-admin-api/core"
	"gin-admin-api/model"
	"gin-admin-api/result"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetBArticleList 分页查询文章列表
// @Summary 分页查询文章列表
// @Tags 文章相关接口
// @Produce json
// @Description 分页查询文章列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param title query string false "文章名称"
// @Param tagId query string false "标签名称"
// @Param categoryId query string false "分类名称"
// @Param status query int false "文章状态：1->已删除,2->未删除"
// @Param topType query string false "是否置顶：1->未置顶,2->已置顶"
// @Success 200 {object} result.Result
// @router /api/bArticle/list [get]
// @Security ApiKeyAuth
func GetBArticleList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Title := c.Query("title")
	TagId := c.Query("tagId")
	CategoryId := c.Query("categoryId")
	Status, _ := strconv.Atoi(c.Query("status"))
	TopType := c.Query("topType")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var bArticleVo []model.BArticleVo
	var count int64
	curDb := Db.Table("b_article").
		Select("b_article.*, user.username, b_tags.tag_name, b_category.category_name").
		Joins("LEFT JOIN user ON user.id = b_article.user_id").
		Joins("LEFT JOIN b_tags ON b_tags.id = b_article.tag_id").
		Joins("LEFT JOIN b_category ON b_category.id = b_article.category_id")
	if Title != "" {
		curDb = curDb.Where("b_article.title = ?", Title)
	}
	if TagId != "" {
		curDb = curDb.Where("b_article.tag_id = ?", TagId)
	}
	if CategoryId != "" {
		curDb = curDb.Where("b_article.category_id = ?", CategoryId)
	}
	if Status != 0 {
		curDb = curDb.Where("b_article.status = ?", Status)
	}
	if TopType != "" {
		curDb = curDb.Where("b_article.top_type = ?", TopType)
	}
	curDb.Count(&count).
		Limit(PageSize).
		Offset((PageNum - 1) * PageSize).
		Order("b_article.create_time DESC").
		Find(&bArticleVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bArticleVo})
}

// UpdateToType 设置置顶
// @Summary 设置置顶
// @Tags 文章相关接口
// @Produce json
// @Description 设置置顶
// @Param data body model.UpdateTopTypeDto true "data"
// @Success 200 {object} result.Result
// @router /api/bArticle/updateTopType [put]
// @Security ApiKeyAuth
func UpdateToType(c *gin.Context) {
	var dto model.UpdateTopTypeDto
	_ = c.BindJSON(&dto)
	var bArticle model.BArticle
	// 已删除的不能操作
	Db.Where("status = ?", constant.Delete).First(&bArticle, dto.ID)
	if bArticle.ID > 0 {
		result.Failed(c, int(result.ApiCode.DeletedCannotOperate), result.ApiCode.GetMessage(result.ApiCode.DeletedCannotOperate))
		return
	}
	// 查询是否有置顶
	Db.Where("top_type = ?", constant.TopType).First(&bArticle)
	if bArticle.ID > 0 {
		// 取消原来置顶
		bArticle.TopType = constant.NotTopType
		Db.Save(&bArticle)
		// 更新新的置顶
		var newBArticle model.BArticle
		Db.First(&newBArticle, dto.ID)
		newBArticle.TopType = dto.TopType
		Db.Save(&newBArticle)
		result.Success(c, true)
	} else {
		Db.First(&bArticle, dto.ID)
		bArticle.TopType = dto.TopType
		Db.Save(&bArticle)
		result.Success(c, true)
	}
}

// UpdateStatus 设置状态删除
// @Summary 设置状态删除
// @Tags 文章相关接口
// @Produce json
// @Description 设置状态删除
// @Param data body model.BArticleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/bArticle/updateStatus [put]
// @Security ApiKeyAuth
func UpdateStatus(c *gin.Context) {
	var dto model.BArticleIdDto
	_ = c.BindJSON(&dto)
	var bArticle model.BArticle
	Db.First(&bArticle, dto.ID)
	bArticle.Status = constant.Delete
	Db.Save(&bArticle)
	result.Success(c, true)
}

// GetBArticleCommentByArticleId 根据文章的id查询评论列表
// @Summary 根据文章的id查询评论列表
// @Tags 文章相关接口
// @Produce json
// @Description 根据文章的id查询评论列表
// @Param id query uint ture "文章id"
// @Success 200 {object} result.Result
// @router /api/bArticle/comment [get]
// @Security ApiKeyAuth
func GetBArticleCommentByArticleId(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var bArticleComment []model.BArticleComment
	Db.Table("b_article_comment").
		Select("content, username, icon, login_address, create_time").
		Where("article_id = ?", Id).
		Where("p_comment_id = ?", constant.PCommentId).
		Scan(&bArticleComment)
	var commentVo []model.CommentVo
	for _, comment := range bArticleComment {
		item := model.CommentVo{
			Content:      comment.Content,
			Username:     comment.Username,
			Icon:         comment.Icon,
			LoginAddress: comment.LoginAddress,
			CreateTime:   comment.CreateTime,
		}
		commentVo = append(commentVo, item)
	}
	result.Success(c, commentVo)
}
