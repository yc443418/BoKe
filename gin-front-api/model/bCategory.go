// 文章分类相关模型
// @author chen

package model

import util "gin-front-api/utils"

// BCategory 文章分类模型对象
type BCategory struct {
	ID           uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                         // ID
	CategoryName string     `gorm:"column:category_name;varchar(64);comment:'分类名称';NOT NULL" json:"categoryName"` // 分类名称
	Sort         uint       `gorm:"column:sort;comment:'排序'" json:"sort"`                                         // 排序
	CreateTime   util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                 // 创建时间
}

func (BCategory) TableName() string {
	return "b_category"
}

// BCategorySelectVo 文章分类视图对象
type BCategorySelectVo struct {
	ID           uint   `json:"id"`           // ID
	CategoryName string `json:"categoryName"` // 分类名称
}
