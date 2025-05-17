// 标签相关模型
// @author chen

package model

import util "gin-front-api/utils"

// BTags 标签模型
type BTags struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`               // ID
	TagName    string     `gorm:"column:tag_name;varchar(32);comment:'标签名称';NOT NULL" json:"tagName"` //  标签名称
	Sort       uint       `gorm:"column:sort;comment:'排序';DEFAULT NULL" json:"sort"`                  // 排序
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`       // 创建时间
}

func (BTags) TableName() string {
	return "b_tags"
}

// BTagsSelectVo 标签视图对象
type BTagsSelectVo struct {
	ID      uint   `json:"id"`      // ID
	TagName string `json:"tagName"` //  标签名称
}
