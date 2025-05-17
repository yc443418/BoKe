// 文章相关接口
// @author chen

package api

import (
	"gin-front-api/constant"
	"gin-front-api/core"
	. "gin-front-api/core"
	"gin-front-api/model"
	"gin-front-api/result"
	util "gin-front-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// CreateBArticle 新增文章
// @Summary 新增文章
// @Tags 文章相关接口
// @Produce json
// @Description 新增文章
// @Param data body model.BArticleDto true "data"
// @Success 200 {object} result.Result
// @Router /api/bArticle/add [post]
// @Security ApiKeyAuth
func CreateBArticle(c *gin.Context) {
	var dto model.BArticleDto
	_ = c.BindJSON(&dto)
	bArticle := GetBArticleByName(dto.Title)
	if bArticle.ID > 0 {
		result.Failed(c, int(result.ApiCode.ArticleIsExist), result.ApiCode.GetMessage(result.ApiCode.ArticleIsExist))
		return
	}
	user, _ := core.GetUser(c)
	addBArticle := model.BArticle{
		UserId:          user.ID,
		CategoryId:      dto.CategoryId,
		TagId:           dto.TagId,
		Title:           dto.Title,
		Image:           dto.Image,
		Summary:         dto.Summary,
		MarkdownContent: dto.MarkdownContent,
		Content:         dto.Content,
		TopType:         constant.NotTopType,
		CreateTime:      util.HTime{Time: time.Now()},
		Status:          constant.NotDelete,
	}
	Db.Create(&addBArticle)
	result.Success(c, addBArticle)
}

// GetBArticleVoList 分页查询文章列表
// @Summary 分页查询文章列表
// @Tags 文章相关接口
// @Produce json
// @Description 分页查询文章列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param tagId query string false "标签id"
// @Param categoryId query string false "分类id"
// @Param orderType query int false "查询参数：1->热榜,2->发布,3->最新"
// @Success 200 {object} result.Result
// @Router /api/bArticle/list [get]
func GetBArticleVoList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	TagId := c.Query("tagId")
	CategoryId := c.Query("categoryId")
	OrderType, _ := strconv.Atoi(c.Query("orderType"))
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var bArticleVo []model.BArticleVo
	var count int64
	curDb := Db.Table("b_article").
		Select("b_article.*, user.username, user.icon, user.login_address, b_tags.tag_name, b_category.category_name").
		Joins("LEFT JOIN user ON user.id = b_article.user_id").
		Joins("LEFT JOIN b_tags ON b_tags.id = b_article.tag_id").
		Joins("LEFT JOIN b_category ON b_category.id = b_article.category_id")
	if TagId != "" {
		curDb = curDb.Where("b_article.tag_id = ?", TagId)
	}
	if CategoryId != "" {
		curDb = curDb.Where("b_article.category_id = ?", CategoryId)
	}
	curDb.Where("b_article.status = ?", constant.NotDelete)
	if OrderType == constant.HOT {
		curDb.Count(&count).
			Limit(PageSize).
			Offset((PageNum - 1) * PageSize).
			Order("b_article.top_type DESC").
			Order("b_article.comment_count DESC").
			Order("b_article.good_count DESC").
			Order("b_article.quantity DESC").
			Find(&bArticleVo)
	} else if OrderType == constant.PUBLISH {
		curDb.Count(&count).
			Limit(PageSize).
			Offset((PageNum - 1) * PageSize).
			Order("b_article.top_type DESC").
			Order("b_article.create_time ASC").
			Find(&bArticleVo)
	} else if OrderType == constant.NEW {
		curDb.Count(&count).
			Limit(PageSize).
			Offset((PageNum - 1) * PageSize).
			Order("b_article.top_type DESC").
			Order("b_article.create_time DESC").
			Find(&bArticleVo)
	} else {
		curDb.Count(&count).
			Limit(PageSize).
			Offset((PageNum - 1) * PageSize).
			Order("b_article.id DESC").
			Find(&bArticleVo)
	}
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bArticleVo})
}

// GetBArticleDetail 查询文章详情
// @Summary 查询文章详情
// @Tags 文章相关接口
// @Produce json
// @Description 查询文章详情
// @Param id query uint false "文章id"
// @Param userId query uint false "用户id"
// @Success 200 {object} result.Result
// @Router /api/bArticle/detail [get]
func GetBArticleDetail(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	UserId, _ := strconv.Atoi(c.Query("userId"))
	// 文章阅读数加1
	var bArticle model.BArticle
	Db.First(&bArticle, Id)
	bArticle.Quantity = bArticle.Quantity + 1
	Db.Save(&bArticle)
	// 文章详情
	var bArticleDetailVo model.BArticleDetailVo
	Db.Table("b_article").
		Select("b_article.*, user.username, user.icon, b_tags.tag_name, b_category.category_name").
		Joins("LEFT JOIN user ON user.id = b_article.user_id").
		Joins("LEFT JOIN b_tags ON b_tags.id = b_article.tag_id").
		Joins("LEFT JOIN b_category ON b_category.id = b_article.category_id").
		Where("b_article.id = ?", Id).
		Scan(&bArticleDetailVo)
	// 查询当前登录的用户是否点赞
	bool := QueryArticleIsLikeByUserId(uint(Id), uint(UserId))
	if bool {
		bArticleDetailVo.HaveLike = true
	} else {
		bArticleDetailVo.HaveLike = false
	}
	result.Success(c, map[string]interface{}{"bArticleDetailVo": bArticleDetailVo})
}

// DoBArticleLike 文章点赞和取消
// @Summary 文章点赞和取消
// @Tags 文章相关接口
// @Produce json
// @Description 文章点赞和取消
// @Param data body model.IdDto true "data"
// @Success 200 {object} result.Result
// @Router /api/bArticle/like [post]
// @Security ApiKeyAuth
func DoBArticleLike(c *gin.Context) {
	// 绑定参数
	var dto model.IdDto
	var bArticle model.BArticle
	_ = c.BindJSON(&dto)
	// 获取当前登录用户
	user, _ := core.GetUser(c)
	// 查询是否点赞
	bool := QueryArticleIsLikeByUserId(dto.Id, user.ID)
	// 点赞及取消点赞
	if bool {
		// 删除点赞记录
		Db.Table("b_article_star").
			Where("article_id = ?", dto.Id).
			Where("user_id = ?", user.ID).
			Delete(&model.BArticleStar{})
		// 文章点赞数量减去1
		Db.First(&bArticle, dto.Id)
		bArticle.GoodCount = bArticle.GoodCount - 1
		Db.Save(&bArticle)
		result.Success(c, constant.NotLikeSuccess)
		return
	} else {
		// 新增点赞记录
		bArticleStar := model.BArticleStar{
			UserId:    user.ID,
			ArticleId: dto.Id,
		}
		Db.Create(&bArticleStar)
		// 文章点赞数量加1
		Db.First(&bArticle, dto.Id)
		bArticle.GoodCount = bArticle.GoodCount + 1
		Db.Save(&bArticle)
		// 发送点赞消息
		addUserMessage := model.UserMessage{
			MessageType:    constant.StarArticle,
			CreateTime:     util.HTime{Time: time.Now()},
			ArticleId:      bArticle.ID,
			ArticleTitle:   bArticle.Title,
			CommentId:      constant.ZERO,
			SendUserId:     user.ID,
			SendUsername:   user.Username,
			SendUserIcon:   user.Icon,
			Status:         constant.NotReadStatus,
			ReceivedUserId: bArticle.UserId,
		}
		Db.Save(&addUserMessage)
		result.Success(c, constant.LikeSuccess)
		return
	}
}

// GetBArticleByArticleId 根据id查询文章
// @Summary 根据id查询文章
// @Tags 文章相关接口
// @Produce json
// @Description 根据id查询文章
// @Param id query uint false "文章id"
// @Success 200 {object} result.Result
// @Router /api/bArticle/articleId [get]
// @Security ApiKeyAuth
func GetBArticleByArticleId(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	User, _ := core.GetUser(c)
	// 查询文章是否存在
	var bArticle model.BArticle
	Db.First(&bArticle, Id)
	if bArticle.ID == 0 {
		result.Failed(c, int(result.ApiCode.ArticleIsNotExist), result.ApiCode.GetMessage(result.ApiCode.ArticleIsNotExist))
		return
	}
	// 是否是当前登录用户的文章
	if bArticle.UserId != User.ID {
		result.Failed(c, int(result.ApiCode.ArticleIsNotYou), result.ApiCode.GetMessage(result.ApiCode.ArticleIsNotYou))
		return
	}
	result.Success(c, bArticle)
}

// UpdateBArticle 修改文章
// @Summary 修改文章
// @Tags 文章相关接口
// @Produce json
// @Description 修改文章
// @Param data body model.BArticleDto true "data"
// @Success 200 {object} result.Result
// @router /api/bArticle/update [put]
// @Security ApiKeyAuth
func UpdateBArticle(c *gin.Context) {
	var dto model.BArticleDto
	_ = c.BindJSON(&dto)
	var bArticle model.BArticle
	Db.First(&bArticle, dto.Id)
	bArticle.CategoryId = dto.CategoryId
	bArticle.TagId = dto.TagId
	bArticle.Title = dto.Title
	bArticle.Image = dto.Image
	bArticle.Summary = dto.Summary
	bArticle.MarkdownContent = dto.MarkdownContent
	bArticle.Content = dto.Content
	Db.Save(&bArticle)
	result.Success(c, bArticle)
}

// GetBArticleListByKeywords 根据关键字分页查询文章列表
// @Summary 根据关键字分页查询文章列表
// @Tags 文章相关接口
// @Produce json
// @Description 根据关键字分页查询文章列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param keywords query string true "文章标题或用户名称"
// @Success 200 {object} result.Result
// @Router /api/bArticle/keyword/list [get]
func GetBArticleListByKeywords(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Keywords := c.Query("keywords")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var bArticleVo []model.BArticleVo
	var count int64
	Db.Table("b_article").
		Select("b_article.*, user.username, user.icon, user.login_address, b_tags.tag_name, b_category.category_name").
		Joins("LEFT JOIN user ON user.id = b_article.user_id").
		Joins("LEFT JOIN b_tags ON b_tags.id = b_article.tag_id").
		Joins("LEFT JOIN b_category ON b_category.id = b_article.category_id").
		Where("b_article.status = ?", constant.NotDelete).
		Where("(b_article.title = ?) OR (user.username = ?)", Keywords, Keywords).
		Count(&count).
		Limit(PageSize).
		Offset((PageNum - 1) * PageSize).
		Order("b_article.create_time DESC").
		Find(&bArticleVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bArticleVo})
}

// GetBArticleByName 根据文章名称查询文章
func GetBArticleByName(title string) (bArticle model.BArticle) {
	Db.Where("title = ?", title).First(&bArticle)
	return bArticle
}

// QueryArticleIsLikeByUserId 查询用户是否点赞
func QueryArticleIsLikeByUserId(articleId, userId uint) bool {
	tx := Db.Table("b_article_star").
		Where("user_id = ?", userId).
		Where("article_id = ?", articleId).
		Find(&model.BArticleStar{})
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
