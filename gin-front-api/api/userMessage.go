// 用户消息相关接口
// @author chen

package api

import (
	"gin-front-api/constant"
	. "gin-front-api/core"
	"gin-front-api/model"
	"gin-front-api/result"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetMessageCount 查询用户消息数量
// @Summary 查询用户消息数量
// @Tags 用户消息相关接口
// @Produce json
// @Description 查询用户消息数量
// @Param userId query uint true "用户id"
// @Success 200 {object} result.Result
// @Router /api/user/message/count [get]
func GetMessageCount(c *gin.Context) {
	UserId, _ := strconv.Atoi(c.Query("userId"))
	// 回复数
	var replyCount int64
	Db.Model(&model.UserMessage{}).
		Select("count(1)").
		Where("received_user_id = ?", UserId).
		Where("status = ?", constant.NotReadStatus).
		Where("message_type = ?", constant.Reply).
		Scan(&replyCount)
	// 点赞文章数
	var starArticleCount int64
	Db.Model(&model.UserMessage{}).
		Select("count(1)").
		Where("received_user_id = ?", UserId).
		Where("status = ?", constant.NotReadStatus).
		Where("message_type = ?", constant.StarArticle).
		Scan(&starArticleCount)
	// 点赞评论数
	var starCommentCount int64
	Db.Model(&model.UserMessage{}).
		Select("count(1)").
		Where("received_user_id = ?", UserId).
		Where("status = ?", constant.NotReadStatus).
		Where("message_type = ?", constant.StarComment).
		Scan(&starCommentCount)
	// 系统通知数
	var systemCount int64
	Db.Model(&model.UserMessage{}).
		Select("count(1)").
		Where("received_user_id = ?", UserId).
		Where("status = ?", constant.NotReadStatus).
		Where("message_type = ?", constant.SystemMessage).
		Scan(&systemCount)
	// 数据组装
	userMessageVo := model.UserMessageVo{
		ReplyCount:       replyCount,
		StarArticleCount: starArticleCount,
		StarCommentCount: starCommentCount,
		SystemCount:      systemCount,
		TotalCount:       replyCount + starArticleCount + starCommentCount + systemCount,
	}
	result.Success(c, userMessageVo)
}

// GetUserMessageList 分页查询用户消息列表
// @Summary 分页查询用户消息列表
// @Tags 用户消息相关接口
// @Produce json
// @Description 分页查询用户消息列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param userId query int true "用户id"
// @Param type query int true "类型（1-回复我的，2-赞了文章，3-赞了评论，4-系统消息）"
// @Success 200 {object} result.Result
// @router /api/user/message/list [get]
func GetUserMessageList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	UserId, _ := strconv.Atoi(c.Query("userId"))
	Type, _ := strconv.Atoi(c.Query("type"))
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var userMessage []model.UserMessage
	var count int64
	Db.Table("user_message").
		Where("received_user_id = ?", UserId).
		Where("message_type = ?", Type).
		Count(&count).
		Limit(PageSize).
		Offset((PageNum - 1) * PageSize).
		Order("id DESC").
		Find(&userMessage)
	// 更新消息状态为已读状态
	if PageNum == 1 {
		Db.Table("user_message").
			Where("received_user_id = ?", UserId).
			Where("message_type = ?", Type).
			Updates(map[string]interface{}{"status": constant.ReadStatus})
	}
	// 查询更新后回复数，点赞文章数，点赞评论数，系统通知数
	var replyCount, starArticleCount, starCommentCount, systemCount int64
	if Type == constant.Reply {
		// 查询更新后的回复数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", Type).
			Scan(&replyCount)
		// 查询更新后的点赞文章数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.StarArticle).
			Scan(&starArticleCount)
		// 查询更新后的点赞评论数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.StarComment).
			Scan(&starCommentCount)
		// 查询更新后的系统通知数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.SystemMessage).
			Scan(&systemCount)
	} else if Type == constant.StarArticle {
		// 查询更新后的回复数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.Reply).
			Scan(&replyCount)
		// 查询更新后的点赞文章数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", Type).
			Scan(&starArticleCount)
		// 查询更新后的点赞评论数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.StarComment).
			Scan(&starCommentCount)
		// 查询更新后的系统通知数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.SystemMessage).
			Scan(&systemCount)
	} else if Type == constant.StarComment {
		// 查询更新后的回复数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.Reply).
			Scan(&replyCount)
		// 查询更新后的点赞文章数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.StarArticle).
			Scan(&starArticleCount)
		// 查询更新后的点赞评论数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", Type).
			Scan(&starCommentCount)
		// 查询更新后的系统通知数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.SystemMessage).
			Scan(&systemCount)
	} else if Type == constant.SystemMessage {
		// 查询更新后的回复数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.Reply).
			Scan(&replyCount)
		// 查询更新后的点赞文章数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.StarArticle).
			Scan(&starArticleCount)
		// 查询更新后的点赞评论数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", constant.StarComment).
			Scan(&starCommentCount)
		// 查询更新后的系统通知数
		Db.Model(&model.UserMessage{}).
			Select("count(1)").
			Where("received_user_id = ?", UserId).
			Where("status = ?", constant.NotReadStatus).
			Where("message_type = ?", Type).
			Scan(&systemCount)
	}
	// 数据组装
	userMessageVo := model.UserMessageVo{
		ReplyCount:       replyCount,
		StarCommentCount: starCommentCount,
		StarArticleCount: starArticleCount,
		SystemCount:      systemCount,
		TotalCount:       replyCount + starCommentCount + starArticleCount + systemCount,
	}
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": userMessage, "userMessageVo": userMessageVo})
}
