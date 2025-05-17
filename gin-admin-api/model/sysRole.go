// 角色 相关模型
// @author chen

package model

import util "gin-admin-api/utils"

// SysRole 角色模型
type SysRole struct {
	ID          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                        // ID
	RoleName    string     `gorm:"column:role_name;varchar(64);comment:'角色名称';NOT NULL" json:"roleName"`        // 角色名称
	RoleKey     string     `gorm:"column:role_key;varchar(64);comment:'权限字符串';NOT NULL" json:"roleKey"`         // 权限字符串
	Status      int        `gorm:"column:status;default:1;comment:'帐号启用状态：1->启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态：1->启用,2->禁用
	Description string     `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`             // 描述
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                // 创建时间
}

func (SysRole) TableName() string {
	return "sys_role"
}

// AddSysRoleDto 新增角色参数
type AddSysRoleDto struct {
	RoleName    string `json:"roleName"`    // 角色名称
	RoleKey     string `json:"roleKey"`     // 权限字符串
	Status      int    `json:"status"`      // 帐号启用状态：1->启用,2->禁用
	Description string `json:"description"` // 描述
}

// UpdateSysRoleDto 修改角色参数
type UpdateSysRoleDto struct {
	ID          uint   `json:"id"`          // ID
	RoleName    string `json:"roleName"`    // 角色名称
	RoleKey     string `json:"roleKey"`     // 权限字符串
	Status      int    `json:"status"`      // 帐号启用状态：1->启用,2->禁用
	Description string `json:"description"` // 描述
}

// SysRoleIdDto Id参数
type SysRoleIdDto struct {
	ID uint `json:"id"` // ID
}

// UpdateSysRoleStatusDto 设置状态参数
type UpdateSysRoleStatusDto struct {
	ID     uint `json:"id"`     // ID
	Status int  `json:"status"` // 帐号启用状态：1->启用,2->禁用
}

// SysRoleVo 角色下拉列表
type SysRoleVo struct {
	ID       uint   `json:"id"`       // ID
	RoleName string `json:"roleName"` // 角色名称
}

// IdVo 当前角色的菜单权限id
type IdVo struct {
	ID int `json:"id"` // ID
}

// RoleMenuDto 角色id和菜单id列表参数对象
type RoleMenuDto struct {
	ID      uint   `json:"id" binding:"required"`      // 角色ID
	MenuIds []uint `json:"menuIds" binding:"required"` // 菜单id列表
}
