// 用户相关模型
// @author chen

package model

import util "gin-front-api/utils"

// User 用户模型
type User struct {
	ID           uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        // ID
	Username     string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`         // 用户账号
	Password     string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`           // 密码
	Status       int        `gorm:"column:status;default:1;comment:'帐号启用状态：1->启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态：1->启用,2->禁用
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

// RegisterDto 注册参数
type RegisterDto struct {
	Username      string `json:"username" validate:"required"`      // 用户账号
	Email         string `json:"email" validate:"required"`         // 邮箱
	EmailCode     string `json:"emailCode" validate:"required"`     // 邮箱验证码
	Password      string `json:"password" validate:"required"`      // 密码
	ResetPassword string `json:"resetPassword" validate:"required"` // 重复密码
}

// JwtUser 鉴权用户结构体
type JwtUser struct {
	ID           uint   `json:"id"`           // ID
	Username     string `json:"username"`     // 用户名
	Icon         string `json:"icon"`         // 头像
	Email        string `json:"email"`        // 邮箱
	Note         string `json:"note"`         // 备注
	ExpireTime   int64  `yaml:"expireTime"`   // 过期时间
	LoginAddress string `json:"loginAddress"` // 登录地址
}

// LoginDto 登录参数结构体
type LoginDto struct {
	Email     string `json:"email" validate:"required"`                 // 邮箱
	Password  string `json:"password" validate:"required"`              // 密码
	CheckCode string `json:"checkCode" validate:"required,min=4,max=6"` // 验证码
	IdKey     string `json:"idKey"`                                     // uuid
}

// UserVo 用户详情
type UserVo struct {
	ID           uint       `json:"id"`           // ID
	Username     string     `json:"username"`     // 用户账号
	Status       int        `json:"status"`       // 帐号启用状态：1->启用,2->禁用
	Icon         string     `json:"icon"`         // 头像
	Sex          int        `json:"sex"`          // 性别：1->男,2->女
	Email        string     `json:"email"`        // 邮箱
	Note         string     `json:"note"`         // 备注
	LoginIp      string     `json:"loginIp"`      // 登录ip
	LoginAddress string     `json:"loginAddress"` // 登录地址
	CreateTime   util.HTime `json:"createTime"`   // 创建时间
}

// ResetPasswordDto 重置密码参数
type ResetPasswordDto struct {
	Email         string `json:"email" validate:"required"`         // 邮箱
	EmailCode     string `json:"emailCode" validate:"required"`     // 邮箱验证码
	Password      string `json:"password" validate:"required"`      // 密码
	ResetPassword string `json:"resetPassword" validate:"required"` // 重复密码
}

// ArticleUserRankVo 文章用户排行
type ArticleUserRankVo struct {
	Id           uint   `json:"id"`                                       // 用户id
	Username     string `json:"username"`                                 // 用户名称
	Icon         string `json:"icon"`                                     // 头像
	ArticleCount uint   `json:"articleCount"`                             // 用户文章数量
	Quantity     uint   `json:"quantity"`                                 // 文章阅读量
	GoodCount    uint   `gorm:"column:good_count" json:"goodCount"`       // 文章点赞量
	CommentCount uint   `gorm:"column:comment_count" json:"commentCount"` // 文章评论量
}

// IdDto 参数
type IdDto struct {
	Id uint `json:"id"` // ID
}

// UpdateUserDto 修改个人信息参数
type UpdateUserDto struct {
	Id       uint   //ID
	Icon     string // 头像
	Username string `validate:"required"` //用户名
	Note     string `validate:"required"` // 备注
}

// UpdateUserPasswordDto 修改个人密码
type UpdateUserPasswordDto struct {
	Id            uint   //ID
	OldPassword   string `validate:"required"` // 密码
	NewPassword   string `validate:"required"` // 新密码
	ResetPassword string `validate:"required"` // 重复密码
}

// ArticleUserDetailVo 文章用户详情
type ArticleUserDetailVo struct {
	ID           uint       `json:"id"`           // 用户ID
	Username     string     `json:"username"`     // 用户账号
	Icon         string     `json:"icon"`         // 头像
	Email        string     `json:"email"`        // 邮箱
	LoginIp      string     `json:"loginIp"`      // 登录ip
	LoginAddress string     `json:"loginAddress"` // 登录地址
	CreateTime   util.HTime `json:"createTime"`   // 创建时间
	Note         string     `json:"note"`         // 备注
	ArticleCount int64      `json:"articleCount"` // 发布文章数量
	LikeCount    int64      `json:"likeCount"`    // 文章点赞数量
}
