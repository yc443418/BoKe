// 文章评论相关接口
// @author chen

package api

import (
	"gin-admin-api/constant"
	. "gin-admin-api/core"
	"gin-admin-api/model"
	"gin-admin-api/result"
	util "gin-admin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// GetBArticleCommentList 分页查询文章评论列表
// @Summary 分页查询文章评论列表
// @Tags 文章评论相关接口
// @Produce json
// @Description 分页查询文章评论列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param title query string false "文章名称"
// @Param username query string false "用户名称"
// @Param status query int false "文章状态（1-已删除，2-未删除）"
// @Success 200 {object} result.Result
// @router /api/bArticle/comment/list [get]
// @Security ApiKeyAuth
func GetBArticleCommentList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Title := c.Query("title")
	Username := c.Query("username")
	Status, _ := strconv.Atoi(c.Query("status"))
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var bArticleCommentVo []model.BArticleCommentVo
	var count int64
	curDb := Db.Table("b_article_comment").
		Select("b_article_comment.*, b_article.title").
		Joins("LEFT JOIN b_article ON b_article.id = b_article_comment.article_id")
	if Title != "" {
		curDb = curDb.Where("b_article.title = ?", Title)
	}
	if Username != "" {
		curDb = curDb.Where("b_article_comment.username = ?", Username)
	}
	if Status != 0 {
		curDb = curDb.Where("b_article_comment.status = ?", Status)
	}
	curDb.Where("b_article_comment.p_comment_id = ?", constant.PCommentId).
		Count(&count).
		Limit(PageSize).
		Offset((PageNum - 1) * PageSize).
		Order("b_article_comment.create_time DESC").
		Find(&bArticleCommentVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bArticleCommentVo})
}

// UpdateArticleCommentStatus 设置状态删除文章评论
// @Summary 设置状态删除文章评论
// @Tags 文章评论相关接口
// @Produce json
// @Description 设置状态删除文章评论
// @Param data body model.BArticleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/bArticle/comment/updateStatus [put]
// @Security ApiKeyAuth
func UpdateArticleCommentStatus(c *gin.Context) {
	// 删除评论
	var dto model.BArticleIdDto
	_ = c.BindJSON(&dto)
	var bArticleComment model.BArticleComment
	Db.First(&bArticleComment, dto.ID)
	bArticleComment.Status = constant.Delete
	Db.Save(&bArticleComment)
	// 更新文章评论数量
	var bArticle model.BArticle
	Db.First(&bArticle, bArticleComment.ArticleId)
	bArticle.CommentCount = bArticle.CommentCount - 1
	Db.Save(&bArticle)
	// 发送信息
	addUserMessage := model.UserMessage{
		ReceivedUserId: bArticleComment.UserId,
		MessageContent: "评论【" + bArticleComment.Content + "】被管理员删除",
		MessageType:    constant.SystemMessage,
		CreateTime:     util.HTime{Time: time.Now()},
		Status:         constant.NotReadStatus,
	}
	Db.Create(&addUserMessage)
	result.Success(c, true)
}
