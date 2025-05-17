// 首页相关接口
// @author chen

package api

import (
	"gin-admin-api/constant"
	. "gin-admin-api/core"
	"gin-admin-api/model"
	"gin-admin-api/result"
	"github.com/gin-gonic/gin"
)

// GetDataStatisticsList 数据类目统计
// @Summary 数据类目统计
// @Tags 首页相关接口
// @Produce json
// @Description 数据类目统计
// @Success 200 {object} result.Result
// @router /api/index/statistics/list [get]
// @Security ApiKeyAuth
func GetDataStatisticsList(c *gin.Context) {
	// 定义list数组
	list := make([]model.DataStatisticsVo, 0, 10)
	// 数量定义
	var articleCount, categoryCount, tagsCount, userCount int64
	var yesterdayArticleCount, yesterdayCategoryCount, yesterdayTagsCount, yesterdayUserCount int64
	// 文章统计
	Db.Model(&model.BArticle{}).
		Select("count(1)").
		Where("b_article.status = ?", constant.NotDelete).
		Scan(&articleCount)
	Db.Model(&model.BArticle{}).
		Select("count(1)").
		Where("date_sub(curdate(), interval 7 day) <= create_time").
		Scan(&yesterdayArticleCount)
	dataStatisticsVoArticle := model.DataStatisticsVo{
		StatisticName:  constant.ArticleStatistic,
		StatisticCount: articleCount,
		YesterdayCount: yesterdayArticleCount,
	}
	list = append(list, dataStatisticsVoArticle)
	// 文章分类统计
	Db.Model(&model.BCategory{}).
		Select("count(1)").
		Scan(&categoryCount)
	Db.Model(&model.BCategory{}).
		Select("count(1)").
		Where("date_sub(curdate(), interval 7 day) <= create_time").
		Scan(&yesterdayCategoryCount)
	dataStatisticsVoCategory := model.DataStatisticsVo{
		StatisticName:  constant.CategoryStatistic,
		StatisticCount: categoryCount,
		YesterdayCount: yesterdayCategoryCount,
	}
	list = append(list, dataStatisticsVoCategory)
	// 文章标签统计
	Db.Model(&model.BTags{}).
		Select("count(1)").
		Scan(&tagsCount)
	Db.Model(&model.BTags{}).
		Select("count(1)").
		Where("date_sub(curdate(), interval 7 day) <= create_time").
		Scan(&yesterdayTagsCount)
	dataStatisticsVoTags := model.DataStatisticsVo{
		StatisticName:  constant.TagsStatistic,
		StatisticCount: tagsCount,
		YesterdayCount: yesterdayTagsCount,
	}
	list = append(list, dataStatisticsVoTags)
	// 文章用户统计
	Db.Model(&model.User{}).
		Select("count(1)").
		Scan(&userCount)
	Db.Model(&model.User{}).
		Select("count(1)").
		Where("date_sub(curdate(), interval 7 day) <= create_time").
		Scan(&yesterdayUserCount)
	dataStatisticsVoUser := model.DataStatisticsVo{
		StatisticName:  constant.UserStatistic,
		StatisticCount: userCount,
		YesterdayCount: yesterdayUserCount,
	}
	list = append(list, dataStatisticsVoUser)
	result.Success(c, map[string]interface{}{"list": list})
}

// GetDataArticleCreateList 近一周新增文章统计
// @Summary 近一周新增文章统计
// @Tags 首页相关接口
// @Produce json
// @Description 近一周新增文章统计
// @Success 200 {object} result.Result
// @router /api/index/data/article/create/list [get]
// @Security ApiKeyAuth
func GetDataArticleCreateList(c *gin.Context) {
	// 查询近一周的时间（时间格式：年/月/日-2024/03/24）
	var dateList []string
	Db.Raw("SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 7 DAY), '%Y-%m-%d') as createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 6 DAY), '%Y-%m-%d') as createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 5 DAY), '%Y-%m-%d') as createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 4 DAY), '%Y-%m-%d') as createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 3 DAY), '%Y-%m-%d') as createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 2 DAY), '%Y-%m-%d') as createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 1 DAY), '%Y-%m-%d') as createTime \n" +
		"FROM DUAL").
		Scan(&dateList)
	// 根据时间查询时间内的新增文章数量
	var articleCount int
	var dataList []int
	for _, date := range dateList {
		Db.Model(&model.BArticle{}).
			Select("count(1)").
			Where("date_format(create_time, '%Y-%m-%d') = ?", date).
			Scan(&articleCount)
		// 查询到的数据添加到数组中
		dataList = append(dataList, articleCount)
	}
	// 添加到list中
	dataCountList := model.DataCountList{
		StatisticName: constant.ArticleStatistic,
		DataList:      dataList,
	}
	list := make([]model.DataCountList, 0, 10)
	list = append(list, dataCountList)
	// 数组组装返回
	dataArticleCreateVo := model.DataArticleCreateVo{
		DateList:         dateList,
		ArticleCountList: list,
	}
	result.Success(c, dataArticleCreateVo)
}

// GetDataUserCreateList 近一周新增用户统计
// @Summary 近一周新增用户统计
// @Tags 首页相关接口
// @Produce json
// @Description 近一周新增用户统计
// @Success 200 {object} result.Result
// @router /api/index/data/user/create/list [get]
// @Security ApiKeyAuth
func GetDataUserCreateList(c *gin.Context) {
	// 查询近一周的时间（时间格式：年/月/日-2024/03/24）
	var dateList []string
	Db.Raw("SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 7 DAY), '%Y-%m-%d') AS createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 6 DAY), '%Y-%m-%d') AS createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 5 DAY), '%Y-%m-%d') AS createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 4 DAY), '%Y-%m-%d') AS createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 3 DAY), '%Y-%m-%d') AS createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 2 DAY), '%Y-%m-%d') AS createTime UNION\n" +
		"SELECT DATE_FORMAT(DATE_SUB(CURDATE(), INTERVAL 1 DAY), '%Y-%m-%d') AS createTime \n" +
		"FROM DUAL").
		Scan(&dateList)
	// 根据时间查询时间内的新增用户数量
	var userCount int
	var dataList []int
	for _, date := range dateList {
		Db.Model(&model.User{}).
			Select("count(1)").
			Where("date_format(create_time, '%Y-%m-%d') = ?", date).
			Scan(&userCount)
		// 查询到的数据添加到数组中
		dataList = append(dataList, userCount)
	}
	// 添加到list中
	dataCountList := model.DataCountList{
		StatisticName: constant.UserStatistic,
		DataList:      dataList,
	}
	list := make([]model.DataCountList, 0, 10)
	list = append(list, dataCountList)
	// 数据组装返回
	dataUserCreateVo := model.DataUserCreateVo{
		DateList:      dateList,
		UserCountList: list,
	}
	result.Success(c, dataUserCreateVo)
}
