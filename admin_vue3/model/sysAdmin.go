package model

import util "admin_vue3/utils"

// SysAdmin 用户模型对象
type SysAdmin struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        // ID
	Username   string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`         // 用户账号
	Password   string     `gorm:"column:password;varchar(64);comment:'密码';NOT NULL" json:"password"`           // 密码
	Nickname   string     `gorm:"column:nickname;varchar(64);comment:'昵称'" json:"nickname"`                    // 昵称
	Status     int        `gorm:"column:status;default:1;comment:'帐号启用状态：1->启用,2->禁用';NOT NULL" json:"status"` // 账号状态
	Icon       string     `gorm:"column:icon;varchar(500);comment:'头像'" json:"icon"`                           // 头像
	Sex        int        `gorm:"column:sex;default:1;comment:'性别：1->男,2->女';NOT NULL" json:"sex"`             // 性别
	Email      string     `gorm:"column:email;varchar(64);comment:'邮箱'" json:"email"`                          // 邮箱
	Phone      string     `gorm:"column:phone;varchar(64);comment:'电话'" json:"phone"`                          // 电话
	Note       string     `gorm:"column:note;varchar(500);comment:'备注'" json:"note"`                           // 备注
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                // 创建时间
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// AddSysAdminDto 新增参数
type AddSysAdminDto struct {
	RoleId   uint   `validate:"required"` // 角色id
	Username string `validate:"required"` // 用户名
	Password string `validate:"required"` // 密码
	Nickname string `validate:"required"` // 昵称
	Sex      int    // 性别: 1->男,2->女
	Phone    string `validate:"required"` // 手机号
	Email    string `validate:"required"` // 邮箱
	Note     string // 备注
	Status   int    `validate:"required"` // 状态: 1->启用,2->禁用
}

// SysAdminInfo 根据id查询用户
type SysAdminInfo struct {
	ID       uint   `json:"id"`       // ID
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Status   int    `json:"status"`   // 状态：1->启用,2->禁用
	RoleId   uint   `json:"roleId"`   // 角色id
	Sex      int    `json:"sex"`      // 性别：1->男,2->女
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Note     string `json:"note"`     // 备注
}

// UpdateSysAdminDto 修改参数
type UpdateSysAdminDto struct {
	Id       uint   // ID
	RoleId   uint   `validate:"required"` // 角色id
	Username string `validate:"required"` // 用户名
	Nickname string `validate:"required"` // 昵称
	Sex      int    // 性别: 1->男,2->女
	Phone    string `validate:"required"` // 手机号
	Email    string `validate:"required"` // 邮箱
	Note     string // 备注
	Status   int    `validate:"required"` // 状态: 1->启用,2->禁用
}

// SysRoleIdDto 根据Id查询用户
type SysAdminIdDto struct {
	ID uint `json:"id"` // ID
}

// UpdateSysRoleStatusDto 设置状态参数
type UpdateSysAdminStatusDto struct {
	ID     uint `json:"id"`     // ID
	Status int  `json:"status"` // 账号启用状态: 1->启用,2->禁用
}

// ResetSysAdminPasswordDto 重置密码参数对象
type ResetSysAdminPasswordDto struct {
	Id       uint   `json:"id"`       // ID
	Password string `json:"password"` // 密码
}

// SysAdminVo 系统用户对象
type SysAdminVo struct {
	ID         uint       `json:"id"`         // ID
	Username   string     `json:"username"`   // 用户账号
	Nickname   string     `json:"nickname"`   // 昵称
	Status     int        `json:"status"`     // 账号启用状态：1->启用,2->禁用
	Icon       string     `json:"icon"`       // 头像
	Sex        int        `json:"sex"`        // 性别：1->男,2->女
	Email      string     `json:"email"`      // 邮箱
	Phone      string     `json:"phone"`      // 电话
	Note       string     `json:"note"`       // 备注
	RoleName   string     `json:"roleName"`   // 角色名称
	CreateTime util.HTime `json:"createTime"` // 创建时间
}

// JwtAdmin 鉴权用户结构体
type JwtAdmin struct {
	ID         uint   `json:"id"`         // ID
	Username   string `json:"username"`   // 用户账号
	Nickname   string `json:"nickname"`   // 昵称
	Status     int    `json:"status"`     // 账号启用状态：1->启用,2->禁用
	Icon       string `json:"icon"`       // 头像
	Email      string `json:"email"`      // 邮箱
	Phone      string `json:"phone"`      // 电话
	Note       string `json:"note"`       // 备注
	ExpireTime int64  `json:"expireTime"` // 过期时间
}

// LoginDto 登录参数对象
type LoginDto struct {
	Username string `json:"username"` // 用户账号
	Password string `json:"password"` // 用户密码
}
