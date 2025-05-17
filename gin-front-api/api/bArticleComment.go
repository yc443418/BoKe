// 文章评论相关接口
// @author chen

package api

import (
	"gin-front-api/constant"
	"gin-front-api/core"
	. "gin-front-api/core"
	"gin-front-api/global"
	"gin-front-api/model"
	"gin-front-api/result"
	util "gin-front-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

// CreateBArticleComment 新增文章评论
// @Summary 新增文章评论
// @Tags 文章评论相关接口
// @Produce json
// @Description 新增文章评论
// @Param data body model.AddCommentDto true "data"
// @Success 200 {object} result.Result
// @Router /api/bArticle/add/comment [post]
// @Security ApiKeyAuth
func CreateBArticleComment(c *gin.Context) {
	// 绑定参数并校验参数
	var dto model.AddCommentDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 获取当前登录的用户
	user, _ := core.GetUser(c)
	global.Log.Infof("当前评论用户id: %d 用户名称: %s", user.ID, user.Username)
	// 判断文章是否存在
	var bArticle model.BArticle
	Db.First(&bArticle, dto.ArticleId)
	if bArticle.ID == 0 {
		result.Failed(c, int(result.ApiCode.ArticleIsNotExist), result.ApiCode.GetMessage(result.ApiCode.ArticleIsNotExist))
		global.Log.Infof("文章不存在,查询的文章id为: %d", dto.ArticleId)
		return
	}
	// 构建新增评论条件
	addBArticleComment := model.BArticleComment{
		PCommentId:   dto.PCommentId,
		ArticleId:    dto.ArticleId,
		Content:      dto.Content,
		UserId:       user.ID,
		Username:     user.Username,
		Icon:         user.Icon,
		LoginAddress: user.LoginAddress,
		TopType:      constant.NotTopType,
		CreateTime:   util.HTime{Time: time.Now()},
		Status:       constant.NotDelete,
	}
	// 新增一级和二级评论
	if dto.PCommentId != 0 {
		// 判断回复是否存在
		var replyUser model.User
		Db.First(&replyUser, dto.ReplyUserId)
		if replyUser.ID == 0 {
			result.Failed(c, int(result.ApiCode.ReplyUserNotExist), result.ApiCode.GetMessage(result.ApiCode.ReplyUserNotExist))
			global.Log.Infof("回复人不存在,查询的回复人id为: %d", dto.ReplyUserId)
			return
		}
		addBArticleComment.ReplyUserId = dto.ReplyUserId
		addBArticleComment.ReplyUsername = replyUser.Username
		addBArticleComment.ReplyIcon = replyUser.Icon
		addBArticleComment.ReplyLoginAddress = replyUser.LoginAddress
		Db.Create(&addBArticleComment)
		// 评论成功后发送消息
		addUserMessage := model.UserMessage{
			MessageType:    constant.Reply,
			CreateTime:     util.HTime{Time: time.Now()},
			ArticleId:      dto.ArticleId,
			CommentId:      addBArticleComment.ID,
			SendUserId:     user.ID,
			SendUsername:   user.Username,
			SendUserIcon:   user.Icon,
			Status:         constant.NotReadStatus,
			MessageContent: dto.Content,
			ArticleTitle:   bArticle.Title,
			ReceivedUserId: addBArticleComment.UserId,
		}
		// 排除自己给自己的评论消息
		if addBArticleComment.UserId != addUserMessage.ReceivedUserId {
			Db.Save(&addUserMessage)
		}
		// 返回二级评论
		var secondCommentList []model.BArticleComment
		Db.Table("b_article_comment bac").
			Select("bac.*").
			Where("bac.article_id = ?", dto.ArticleId).
			Where("bac.p_comment_id = ?", dto.PCommentId).
			Where("bac.user_id = ?", user.ID).
			Find(&secondCommentList)
		result.Success(c, secondCommentList)
	} else {
		Db.Create(&addBArticleComment)
		// 更新文章评论数量
		bArticle.CommentCount = bArticle.CommentCount + 1
		Db.Save(&bArticle)
		// 评论成功后发送消息
		addUserMessage := model.UserMessage{
			MessageType:    constant.Reply,
			CreateTime:     util.HTime{Time: time.Now()},
			ArticleId:      dto.ArticleId,
			CommentId:      addBArticleComment.ID,
			SendUserId:     user.ID,
			SendUsername:   user.Username,
			SendUserIcon:   user.Icon,
			Status:         constant.NotReadStatus,
			MessageContent: dto.Content,
			ArticleTitle:   bArticle.Title,
			ReceivedUserId: bArticle.UserId,
		}
		// 排除自己给自己评论消息
		if addBArticleComment.UserId != addUserMessage.ReceivedUserId {
			Db.Save(&addUserMessage)
		}
		// 返回一级评论
		result.Success(c, addBArticleComment)
	}
}

// GetBArticleCommentVoList 分页查询文章评论接口
// @Summary 分页查询文章评论接口
// @Tags 文章评论相关接口
// @Produce json
// @Description 分页查询文章评论接口
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param articleId query int true "文章id"
// @Param orderType query int true "排序参数(1-最新， 3-热榜)"
// @Param userId query uint false "用户id"
// @Success 200 {object} result.Result
// @router /api/bArticle/commentVo/list [get]
func GetBArticleCommentVoList(c *gin.Context) {
	// 查询参数
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	ArticleId, _ := strconv.Atoi(c.Query("articleId"))
	OrderType, _ := strconv.Atoi(c.Query("orderType"))
	UserId, _ := strconv.Atoi(c.Query("userId"))
	// 默认分页数
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	// 一级评论
	var bArticleComment []model.BArticleComment
	commentVoList, count, pageSize, pageNum := getCommentParentList(ArticleId, OrderType, PageSize, PageNum)
	// 组装一级和二级评论
	for _, commentFirst := range commentVoList {
		commentChildrenVoList := getCommentChildrenList(uint(UserId), commentFirst.ID)
		boolFirst := getCommentLike(uint(UserId), commentFirst.ID)
		itemFirst := model.BArticleComment{}
		itemFirst.ID = commentFirst.ID
		itemFirst.PCommentId = commentFirst.PCommentId
		itemFirst.ArticleId = commentFirst.ArticleId
		itemFirst.Content = commentFirst.Content
		itemFirst.UserId = commentFirst.UserId
		itemFirst.Username = commentFirst.Username
		itemFirst.Icon = commentFirst.Icon
		itemFirst.LoginAddress = commentFirst.LoginAddress
		itemFirst.ReplyUserId = commentFirst.ReplyUserId
		itemFirst.ReplyUsername = commentFirst.ReplyUsername
		itemFirst.ReplyIcon = commentFirst.ReplyIcon
		itemFirst.ReplyLoginAddress = commentFirst.ReplyLoginAddress
		itemFirst.TopType = commentFirst.TopType
		itemFirst.CreateTime = commentFirst.CreateTime
		itemFirst.GoodCount = commentFirst.GoodCount
		itemFirst.Status = commentFirst.Status
		itemFirst.Children = commentChildrenVoList
		if !boolFirst {
			itemFirst.LikeType = constant.LikeComment
		} else {
			itemFirst.LikeType = constant.CancelLikeComment
		}
		bArticleComment = append(bArticleComment, itemFirst)
	}
	result.Success(c, map[string]any{"total": count, "pageSize": pageSize, "pageNum": pageNum, "list": bArticleComment})
}

// DoCommentTop 文章评论评论置顶接口
// @Summary 文章评论评论置顶接口
// @Tags 文章评论相关接口
// @Produce json
// @Description 文章评论评论置顶接口
// @Param data body model.CommentTopDto true "data"
// @Success 200 {object} result.Result
// @router /api/bArticle/comment/top [post]
// @Security ApiKeyAuth
func DoCommentTop(c *gin.Context) {
	// 绑定参数
	var dto model.CommentTopDto
	_ = c.BindJSON(&dto)
	// 查询文章评论是否存在
	var bArticleComment model.BArticleComment
	Db.First(&bArticleComment, dto.Id)
	if bArticleComment.ID == 0 {
		result.Failed(c, int(result.ApiCode.BArticleCommentNotExist), result.ApiCode.GetMessage(result.ApiCode.BArticleCommentNotExist))
		return
	}
	// 查询文章是否存在
	var bArticle model.BArticle
	Db.First(&bArticle, bArticleComment.ArticleId)
	if bArticle.ID == 0 {
		result.Failed(c, int(result.ApiCode.ArticleIsNotExist), result.ApiCode.GetMessage(result.ApiCode.ArticleIsNotExist))
		return
	}
	// 二级评论不能置顶
	if bArticleComment.PCommentId != 0 {
		result.Failed(c, int(result.ApiCode.SecondCommentNotTop), result.ApiCode.GetMessage(result.ApiCode.SecondCommentNotTop))
		return
	}
	// 是否是自己发的评论
	user, _ := core.GetUser(c)
	if bArticleComment.UserId != user.ID {
		result.Failed(c, int(result.ApiCode.NotYouComment), result.ApiCode.GetMessage(result.ApiCode.NotYouComment))
		return
	}
	// 如果有置顶则要取消原来的置顶更新新的置顶
	var bArticleCommentTop model.BArticleComment
	Db.Where("top_type = ?", constant.TopType).First(&bArticleCommentTop)
	if bArticleCommentTop.ID > 0 {
		bArticleCommentTop.TopType = constant.NotTopType
		Db.Save(&bArticleCommentTop)
		bArticleComment.TopType = dto.TopType
		Db.Save(&bArticleComment)
		result.Success(c, true)
		return
	} else {
		bArticleComment.TopType = dto.TopType
		Db.Save(&bArticleComment)
		result.Success(c, true)
	}
}

// DoBArticleCommentLike 文章评论点赞和取消接口
// @Summary 文章评论点赞和取消接口
// @Tags 文章评论相关接口
// @Produce json
// @Description 文章评论点赞和取消接口
// @Param data body model.IdDto true "data"
// @Success 200 {object} result.Result
// @router /api/bArticle/comment/like [post]
// @Security ApiKeyAuth
func DoBArticleCommentLike(c *gin.Context) {
	// 绑定参数
	var dto model.IdDto
	_ = c.BindJSON(&dto)
	var bArticleComment model.BArticleComment
	// 获取当前登录用户
	user, _ := core.GetUser(c)
	// 查询当前用户是否对评论点赞
	bool := getCommentLike(user.ID, dto.Id)
	// 点赞和取消点赞
	if bool {
		// 删除点赞记录
		Db.Table("b_article_comment_star").
			Where("comment_id = ?", dto.Id).
			Where("user_id = ?", user.ID).
			Delete(&model.BArticleCommentStar{})
		// 评论记录点赞数减去1
		Db.First(&bArticleComment, dto.Id)
		bArticleComment.GoodCount = bArticleComment.GoodCount - 1
		bArticleComment.LikeType = constant.LikeComment
		Db.Save(&bArticleComment)
		result.Success(c, bArticleComment)
		return
	} else {
		// 新增点赞记录
		addBArticleCommentStar := model.BArticleCommentStar{
			CommentId: dto.Id,
			UserId:    user.ID,
		}
		Db.Create(&addBArticleCommentStar)
		// 评论的点赞量加1
		Db.First(&bArticleComment, dto.Id)
		bArticleComment.GoodCount = bArticleComment.GoodCount + 1
		bArticleComment.LikeType = constant.CancelLikeComment
		Db.Save(&bArticleComment)
		// 发送点赞消息
		var bArticle model.BArticle
		Db.First(&bArticle, bArticleComment.ArticleId)
		addUserMessage := model.UserMessage{
			MessageType:    constant.StarComment,
			CreateTime:     util.HTime{Time: time.Now()},
			ArticleId:      bArticle.ID,
			ArticleTitle:   bArticle.Title,
			CommentId:      constant.ZERO,
			SendUserId:     user.ID,
			SendUsername:   user.Username,
			SendUserIcon:   user.Icon,
			Status:         constant.NotReadStatus,
			ReceivedUserId: bArticle.UserId,
			MessageContent: bArticleComment.Content,
		}
		Db.Save(&addUserMessage)
		result.Success(c, bArticleComment)
		return
	}
}

// 查询一级评论列表
func getCommentParentList(ArticleId, OrderType, PageSize, PageNum int) (bArticleComment []model.BArticleComment, count int64, pageSize, pageNum int) {
	curDb := Db.Table("b_article_comment").
		Where("p_comment_id = ?", constant.ZERO).
		Where("article_id = ?", ArticleId).
		Where("status = ?", constant.NotDelete)
	if OrderType == constant.HOT {
		curDb.Count(&count).
			Limit(PageSize).
			Offset((PageNum - 1) * PageSize).
			Order("top_type DESC").
			Order("good_count DESC").
			Scan(&bArticleComment)
	} else if OrderType == constant.NEW {
		curDb.Count(&count).
			Limit(PageSize).
			Offset((PageNum - 1) * PageSize).
			Order("top_type DESC").
			Order("create_time DESC").
			Scan(&bArticleComment)
	}
	return bArticleComment, count, PageSize, PageNum
}

// 查询二级评论列表
func getCommentChildrenList(userId, commentId uint) (bArticleCommentSecond []model.BArticleComment) {
	Db.Raw("select *, "+
		"(case when (select user_id from b_article_comment_star where user_id = ? and comment_id = ?) is null then 1 else 2 end) as likeType "+
		"from b_article_comment where p_comment_id = ?", userId, commentId, commentId).
		Find(&bArticleCommentSecond)
	return bArticleCommentSecond
}

// 查询当前用户对文章评论是否点赞
func getCommentLike(userId, commentId uint) bool {
	tx := Db.Table("b_article_comment_star").
		Where("user_id = ?", userId).
		Where("comment_id = ?", commentId).
		Find(&model.BArticleCommentStar{})
	if tx.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
