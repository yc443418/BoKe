// 首页相关结构体
// @author chen

package model

// DataStatisticsVo 近周数据类目统计
type DataStatisticsVo struct {
	StatisticName  string `json:"statisticName"`  // 名称
	StatisticCount int64  `json:"statisticCount"` // 数量
	YesterdayCount int64  `json:"yesterdayCount"` // 新增数据
}

// DataCountList 新增文章和用户统计子集
type DataCountList struct {
	StatisticName string `json:"statisticName"` // 名称
	DataList      []int  `json:"dataList"`      // 数量列表
}

// DataArticleCreateVo 近一周新增文章统计
type DataArticleCreateVo struct {
	DateList         []string        `json:"dateList"`         // 时间列表
	ArticleCountList []DataCountList `json:"articleCountList"` // 文章数量列表
}

// DataUserCreateVo 近一周新增用户统计
type DataUserCreateVo struct {
	DateList      []string        `json:"dateList"`      // 时间列表
	UserCountList []DataCountList `json:"userCountList"` // 用户数量列表
}
