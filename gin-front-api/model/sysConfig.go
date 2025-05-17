// 系统参数相关模型
// @author chen

package model

import util "gin-front-api/utils"

type SysConfig struct {
	ID          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                       // ID
	Name        string     `gorm:"column:name;varchar(32);comment:'参数名称';NOT NULL" json:"name"`                // 参数名称
	ConfigKey   string     `gorm:"column:config_key;varchar(32);comment:'参数键名';NOT NULL" json:"configKey"`     // 参数键名
	ConfigValue string     `gorm:"column:config_value;varchar(32);comment:'参数键值';NOT NULL" json:"configValue"` // 参数键值
	Sort        uint       `gorm:"column:sort;comment:'排序';DEFAULT NULL" json:"sort"`                          // 排序
	Remark      string     `gorm:"column:remark;varchar(32);comment:'备注';NOT NULL" json:"remark"`              // 备注
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`               // 创建时间
}

func (SysConfig) TableName() string {
	return "sys_config"
}
