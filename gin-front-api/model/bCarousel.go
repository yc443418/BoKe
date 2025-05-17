// 文章轮播图相关模型
// @author chen

package model

import util "gin-front-api/utils"

// BCarousel 轮播图模型
type BCarousel struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                           // ID
	Icon       string     `gorm:"column:icon;varchar(200);comment:'封面地址'" json:"icon"`                            //  封面地址
	IsPublish  int        `gorm:"column:is_publish;default:1;comment:'是否发布(1:否 2:是)'';NOT NULL" json:"isPublish"` // 是否发布(1:否 2:是)'
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                   // 创建时间
}

func (BCarousel) TableName() string {
	return "b_carousel"
}
