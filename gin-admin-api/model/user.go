// 前台用户相关模型
// @author chen

package model

import util "gin-admin-api/utils"

// User 前台用户模型
type User struct {
	ID           uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        // ID
	Username     string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`         // 用户账号
	Password     string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`           // 密码
	Status       int        `gorm:"column:status;default:1;comment:'帐号启用状态：1->启用,2->禁用';NOT NULL" json:"status"` // 帐号状态：1->启用,2->禁用
	Icon         string     `gorm:"column:icon;varchar(500);comment:'头像'" json:"icon"`                           //  头像
	Sex          int        `gorm:"column:sex;default:1;comment:'性别：1->男,2->女';NOT NULL" json:"sex"`             // 性别：1->男,2->女
	Email        string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`                          // 邮箱
	Note         string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`                           // 备注
	LoginIp      string     `gorm:"column:login_ip;varchar(255);comment:'登录ip'" json:"loginIp"`                  // 登录ip
	LoginAddress string     `gorm:"column:login_address;varchar(255);comment:'登录地址'" json:"loginAddress"`        // 登录地址
	CreateTime   util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                // 创建时间
}

func (User) TableName() string {
	return "user"
}

// UserVo 前台用户对象
type UserVo struct {
	ID           uint       `json:"id"`           // ID
	Username     string     `json:"username"`     // 用户账号
	Status       int        `json:"status"`       // 帐号状态：1->启用,2->禁用
	Icon         string     `json:"icon"`         //  头像
	Sex          int        `json:"sex"`          // 性别：1->男,2->女
	Quantity     uint       `json:"quantity"`     // 发表文章数量
	Email        string     `json:"email"`        // 邮箱
	Note         string     `json:"note"`         // 备注
	LoginIp      string     `json:"loginIp"`      // 登录ip
	LoginAddress string     `json:"loginAddress"` // 登录地址
	CreateTime   util.HTime `json:"createTime"`   // 创建时间
}

// UpdateUserStatusDto 设置状态参数
type UpdateUserStatusDto struct {
	Id     uint // ID
	Status int  // 状态：1->启用,2->禁用
}
