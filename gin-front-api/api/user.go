// 用户相关接口
// @author chen

package api

import (
	"gin-front-api/component"
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

// Register 用户注册
// @Summary 用户注册
// @Tags 用户相关接口
// @Produce json
// @Description 用户注册
// @Param data body model.RegisterDto true "data"
// @Success 200 {object} result.Result
// @Router /api/register [post]
func Register(c *gin.Context) {
	// 校验参数
	var dto model.RegisterDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 根据用户名查询用户
	userByUsername := GetUserByUsername(dto.Username)
	if userByUsername.ID > 0 {
		result.Failed(c, int(result.ApiCode.UsernameAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.UsernameAlreadyExists))
		return
	}
	// 根据邮箱查询邮箱名不能重复
	userByEmail := GetUserByEmail(dto.Email)
	if userByEmail.ID > 0 {
		result.Failed(c, int(result.ApiCode.EmailAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.EmailAlreadyExists))
		return
	}
	// 判断两次输入密码是否一致
	if dto.Password != dto.ResetPassword {
		result.Failed(c, int(result.ApiCode.PasswordError), result.ApiCode.GetMessage(result.ApiCode.PasswordError))
		return
	}
	// 获取邮箱验证码
	emailCode, err := core.RedisDb.Get(global.Ctx, constant.EmailCode+dto.Email).Result()
	if len(emailCode) == 0 {
		result.Failed(c, int(result.ApiCode.EmailCodeExpire), result.ApiCode.GetMessage(result.ApiCode.EmailCodeExpire))
		return
	}
	// 判断输入的邮箱验证码是否正确
	if dto.EmailCode != emailCode {
		result.Failed(c, int(result.ApiCode.EmailCodeIsFalse), result.ApiCode.GetMessage(result.ApiCode.EmailCodeIsFalse))
		return
	}
	// 创建新用户
	user := model.User{
		Username:   dto.Username,
		Password:   util.EncryptionMd5(dto.Password),
		Email:      dto.Email,
		CreateTime: util.HTime{Time: time.Now()},
	}
	Db.Create(&user)
	// 注册成功发送系统消息
	systemContent, err := core.RedisDb.Get(global.Ctx, constant.RegisterMessage).Result()
	global.Log.Infof("从redis获取的用户注册语: %s", systemContent)
	if err != nil {
		var sysConfig model.SysConfig
		Db.Table("sys_config").Where("config_key = ?", constant.RegisterMessage).Scan(&sysConfig)
		core.RedisDb.Set(global.Ctx, sysConfig.ConfigKey, sysConfig.ConfigValue, 0)
		global.Log.Infof("从mysql获取的用户注册语: %s", sysConfig.ConfigValue)
		addUserMessage := model.UserMessage{
			MessageType:    constant.SystemMessage,
			CreateTime:     util.HTime{Time: time.Now()},
			Status:         constant.NotReadStatus,
			ReceivedUserId: user.ID,
			MessageContent: sysConfig.ConfigValue,
		}
		Db.Save(&addUserMessage)
		result.Success(c, true)
	} else {
		addUserMessage := model.UserMessage{
			MessageType:    constant.SystemMessage,
			CreateTime:     util.HTime{Time: time.Now()},
			Status:         constant.NotReadStatus,
			ReceivedUserId: user.ID,
			MessageContent: systemContent,
		}
		Db.Save(&addUserMessage)
		result.Success(c, true)
	}
}

// Login 用户登录
// @Summary 用户登录
// @Tags 用户相关接口
// @Produce json
// @Description 用户登录
// @Param data body model.LoginDto true "data"
// @Success 200 {object} result.Result
// @Router /api/login [post]
func Login(c *gin.Context) {
	// 绑定json并校验参数
	var dto model.LoginDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 查询用户信息，邮箱，密码正确及状态是否正常
	user := GetUserByEmail(dto.Email)
	if user.ID > 0 {
		if user.Password != util.EncryptionMd5(dto.Password) {
			result.Failed(c, int(result.ApiCode.PasswordNotTrue), result.ApiCode.GetMessage(result.ApiCode.PasswordNotTrue))
			return
		}
		if user.Status == constant.Status {
			result.Failed(c, int(result.ApiCode.StatusIsEnable), result.ApiCode.GetMessage(result.ApiCode.StatusIsEnable))
			return
		}
	} else {
		result.Failed(c, int(result.ApiCode.UserIsNotExist), result.ApiCode.GetMessage(result.ApiCode.UserIsNotExist))
		return
	}
	// 判断验证码是否过期和是否正确
	code := component.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.CodeHasExpired), result.ApiCode.GetMessage(result.ApiCode.CodeHasExpired))
		return
	}
	verifyRes := CaptVerify(dto.IdKey, dto.CheckCode)
	if !verifyRes {
		result.Failed(c, int(result.ApiCode.CaptchaNotTrue), result.ApiCode.GetMessage(result.ApiCode.CaptchaNotTrue))
		return
	}
	// 生成token
	token, _ := core.GenerateTokenByUser(user)
	// 记录用户登录的时间和ip
	user.LoginIp = c.ClientIP()
	user.LoginAddress = util.GetRealAddressByIP(c.ClientIP())
	Db.Save(&user)
	// 生成用户信息
	userVo := model.UserVo{
		ID:           user.ID,
		Username:     user.Username,
		Status:       user.Status,
		Icon:         user.Icon,
		Sex:          user.Sex,
		Email:        user.Email,
		Note:         user.Note,
		LoginIp:      user.LoginIp,
		LoginAddress: user.LoginAddress,
		CreateTime:   user.CreateTime,
	}
	// 返回token和用户信息
	result.Success(c, map[string]any{"token": token, "user": userVo})
}

// ResetPassword 重置密码
// @Summary 重置密码
// @Tags 用户相关接口
// @Produce json
// @Description 重置密码
// @Param data body model.ResetPasswordDto true "data"
// @Success 200 {object} result.Result
// @Router /api/user/reset/password [post]
func ResetPassword(c *gin.Context) {
	// 校验参数
	var dto model.ResetPasswordDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 根据邮箱查询
	userByEmail := GetUserByEmail(dto.Email)
	if userByEmail.ID == 0 {
		result.Failed(c, int(result.ApiCode.UserIsNotExist), result.ApiCode.GetMessage(result.ApiCode.UserIsNotExist))
		return
	}
	// 发送邮箱验证码
	emailCode, err := core.RedisDb.Get(global.Ctx, constant.EmailCode+dto.Email).Result()
	if len(emailCode) == 0 {
		result.Failed(c, int(result.ApiCode.EmailCodeExpire), result.ApiCode.GetMessage(result.ApiCode.EmailCodeExpire))
		return
	}
	if dto.EmailCode != emailCode {
		result.Failed(c, int(result.ApiCode.EmailCodeIsFalse), result.ApiCode.GetMessage(result.ApiCode.EmailCodeIsFalse))
		return
	}
	// 校验两次输入的密码是否一致
	if dto.Password != dto.ResetPassword {
		result.Failed(c, int(result.ApiCode.PasswordError), result.ApiCode.GetMessage(result.ApiCode.PasswordError))
		return
	}
	// 修改密码
	userByEmail.Password = util.EncryptionMd5(dto.Password)
	Db.Save(&userByEmail)
	result.Success(c, true)
}

// GetArticleUserRankList 文章用户排行列表
// @Summary 文章用户排行列表
// @Tags 用户相关接口
// @Produce json
// @Description 文章用户排行列表
// @Success 200 {object} result.Result
// @Router /api/user/article/rank/list [get]
func GetArticleUserRankList(c *gin.Context) {
	// 查询列表
	var articleUserRankVo []model.ArticleUserRankVo
	Db.Table("user").
		Select("DISTINCT user.id, user.username, user.icon, b_article.status").
		Joins("INNER JOIN b_article ON b_article.user_id = user.id").
		Limit(10).
		Where("b_article.status = ?", constant.NotDelete).
		Order("id ASC").
		Scan(&articleUserRankVo)
	// 定义查询的数据
	var articleCount, quantity, goodCount, commentCount uint
	// 定义list存储查询的数据
	list := make([]model.ArticleUserRankVo, 0, 10)
	// 遍历数据并查询
	for _, vo := range articleUserRankVo {
		Db.Model(&model.BArticle{}).Select("count(1)").Where("user_id = ?", vo.Id).Scan(&articleCount)
		Db.Model(&model.BArticle{}).Select("sum(quantity)").Where("user_id = ?", vo.Id).Scan(&quantity)
		Db.Model(&model.BArticle{}).Select("sum(good_count)").Where("user_id = ?", vo.Id).Scan(&goodCount)
		Db.Model(&model.BArticle{}).Select("sum(comment_count)").Where("user_id = ?", vo.Id).Scan(&commentCount)
		item := model.ArticleUserRankVo{
			Id:           vo.Id,
			Username:     vo.Username,
			Icon:         vo.Icon,
			ArticleCount: articleCount,
			Quantity:     quantity,
			GoodCount:    goodCount,
			CommentCount: commentCount,
		}
		// 查询的数据添加到list中
		list = append(list, item)
	}
	result.Success(c, list)
}

// UpdateUserInfo 修改个人信息
// @Summary 修改个人信息
// @Tags 用户相关接口
// @Produce json
// @Description 修改个人信息
// @Param data body model.UpdateUserDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/updateUserInfo [put]
// @Security ApiKeyAuth
func UpdateUserInfo(c *gin.Context) {
	var dto model.UpdateUserDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	userCurrent, _ := core.GetUser(c)
	dto.Id = userCurrent.ID
	var user model.User
	Db.First(&user, dto.Id)
	if dto.Icon != "" {
		user.Icon = dto.Icon
	}
	if dto.Username != "" {
		user.Username = dto.Username
	}
	if dto.Note != "" {
		user.Note = dto.Note
	}
	Db.Save(&user)
	result.Success(c, user)
}

// UpdateUserPassword 修改密码
// @Summary 修改密码
// @Tags 用户相关接口
// @Produce json
// @Description 修改密码
// @Param data body model.UpdateUserPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/updateUserPassword [put]
// @Security ApiKeyAuth
func UpdateUserPassword(c *gin.Context) {
	// 绑定json参数并校验
	var dto model.UpdateUserPasswordDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 获取当前登录的用户并校验密码
	userCurrent, _ := core.GetUser(c)
	dto.Id = userCurrent.ID
	var user model.User
	Db.First(&user, dto.Id)
	if user.Password != util.EncryptionMd5(dto.OldPassword) {
		result.Failed(c, int(result.ApiCode.PasswordNotTrue), result.ApiCode.GetMessage(result.ApiCode.PasswordNotTrue))
		return
	}
	if dto.NewPassword != dto.ResetPassword {
		result.Failed(c, int(result.ApiCode.PasswordError), result.ApiCode.GetMessage(result.ApiCode.PasswordError))
		return
	}
	dto.NewPassword = util.EncryptionMd5(dto.NewPassword)
	// 修改密码
	user.Password = dto.NewPassword
	Db.Save(&user)
	result.Success(c, map[string]interface{}{"user": user})
}

// GetArticleUserDetail 根据用户的id查询文章用户
// @Summary 根据用户的id查询文章用户
// @Tags 用户相关接口
// @Produce json
// @Description 根据用户的id查询文章用户
// @Param userId query uint true "用户id"
// @Success 200 {object} result.Result
// @router /api/user/articleUserDetail [get]
func GetArticleUserDetail(c *gin.Context) {
	// 用户id
	UserId, _ := strconv.Atoi(c.Query("userId"))
	// 查询用户信息
	var user model.User
	Db.First(&user, UserId)
	// 查询文章的发布数量
	var articleCount int64
	Db.Model(&model.BArticle{}).Select("count(1)").Where("user_id = ?", UserId).Scan(&articleCount)
	// 查询文章的点赞数
	var likeCount int64
	Db.Model(&model.BArticle{}).Select("count(1)").Where("user_id = ?", UserId).Scan(&likeCount)
	// 数据组装
	articleUserDetailVo := model.ArticleUserDetailVo{
		ID:           uint(UserId),
		Username:     user.Username,
		Icon:         user.Icon,
		Email:        user.Email,
		LoginIp:      user.LoginIp,
		LoginAddress: user.LoginAddress,
		CreateTime:   user.CreateTime,
		Note:         user.Note,
		ArticleCount: articleCount,
		LikeCount:    likeCount,
	}
	result.Success(c, articleUserDetailVo)
}

// GetArticleUserList 根据id查询文章用户列表
// @Summary 根据id查询文章用户列表
// @Tags 用户相关接口
// @Produce json
// @Description 根据id查询文章用户列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param userId query int true "用户id"
// @Param type query int true "类型(1-文章列表，2-评论列表，3-点赞列表)"
// @Success 200 {object} result.Result
// @router /api/user/bArticle/list [get]
func GetArticleUserList(c *gin.Context) {
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
	var bArticleVo []model.BArticleVo
	var count int64
	curDb := Db.Table("b_article").
		Select("b_article.*, user.username, user.icon, user.login_address, b_tags.tag_name, b_category.category_name").
		Joins("LEFT JOIN user ON user.id = b_article.user_id").
		Joins("LEFT JOIN b_tags ON b_tags.id = b_article.tag_id").
		Joins("LEFT JOIN b_category ON b_category.id = b_article.category_id")
	// 类型(1-文章列表，2-评论列表，3-点赞列表)
	if Type == constant.ArticleList {
		curDb = curDb.Where("b_article.user_id = ?", UserId)
	} else if Type == constant.CommentList {
		curDb = curDb.Where("b_article.user_id = ?", UserId).Where("b_article.comment_count != 0")
	} else if Type == constant.StarList {
		curDb = curDb.Where("b_article.user_id = ?", UserId).Where("b_article.good_count != 0")
	}
	curDb.Where("b_article.status = ?", constant.NotDelete).
		Count(&count).
		Limit(PageSize).
		Offset((PageNum - 1) * PageSize).
		Order("b_article.create_time DESC").
		Find(&bArticleVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": bArticleVo})
}

// GetUserByUsername 根据用户名查询用户
func GetUserByUsername(username string) (user model.User) {
	Db.Where("username = ?", username).First(&user)
	return user
}

// GetUserByEmail 根据邮箱查询用户
func GetUserByEmail(email string) (user model.User) {
	Db.Where("email = ?", email).First(&user)
	return user
}

// CaptVerify 验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
