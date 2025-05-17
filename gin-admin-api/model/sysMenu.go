// 菜单相关模型
// @author chen

package model

import util "gin-admin-api/utils"

type SysMenu struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`              // ID
	ParentId   uint       `gorm:"column:parent_id;comment:'父菜单id'" json:"parentId"`                  // 父菜单id
	MenuName   string     `gorm:"column:menu_name;varchar(100);comment:'菜单名称'" json:"menuName"`      // 菜单名称
	Icon       string     `gorm:"column:icon;varchar(100);comment:'菜单图标'" json:"icon"`               // 菜单图标
	Value      string     `gorm:"column:value;varchar(100);comment:'权限值'" json:"value"`              // 权限值
	MenuType   uint       `gorm:"column:menu_type;comment:'菜单类型：1->目录；2->菜单；3->按钮'" json:"menuType"` // 菜单类型：1->目录；2->菜单；3->按钮
	Url        string     `gorm:"column:url;varchar(100);comment:'菜单url'" json:"url"`                // 菜单url
	MenuStatus uint       `gorm:"column:menu_status;comment:'启用状态；1->启用；2->禁用'" json:"menuStatus"`   // 启用状态；1->启用；2->禁用
	Sort       uint       `gorm:"column:sort;comment:'排序'" json:"sort"`                              // 排序
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间'" json:"createTime"`               // 创建时间
	Children   []SysMenu  `json:"children" gorm:"-"`                                                 // 子集
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

// AddSysMenuDto 新增菜单参数
type AddSysMenuDto struct {
	ParentId   uint   `json:"parentId"`   // 父菜单id
	MenuName   string `json:"menuName"`   // 菜单名称
	Icon       string `json:"icon"`       // 菜单图标
	Value      string `json:"value"`      // 权限值
	MenuType   uint   `json:"menuType"`   // 菜单类型：1->目录；2->菜单；3->按钮
	Url        string `json:"url"`        // 菜单url
	MenuStatus uint   `json:"menuStatus"` // 启用状态；1->启用；2->禁用
	Sort       uint   `json:"sort"`       // 排序
}

// UpdateSysMenuDto 修改菜单参数
type UpdateSysMenuDto struct {
	ID         uint   `json:"id"`         // ID
	ParentId   uint   `json:"parentId"`   // 父菜单id
	MenuName   string `json:"menuName"`   // 菜单名称
	Icon       string `json:"icon"`       // 菜单图标
	Value      string `json:"value"`      // 权限值
	MenuType   uint   `json:"menuType"`   // 菜单类型：1->目录；2->菜单；3->按钮
	Url        string `json:"url"`        // 菜单url
	MenuStatus uint   `json:"menuStatus"` // 启用状态；1->启用；2->禁用
	Sort       uint   `json:"sort"`       // 排序
}

// SysMenuIdDto id参数
type SysMenuIdDto struct {
	ID uint `json:"id"` // ID
}

// MenuSvo 菜单级vo
type MenuSvo struct {
	MenuName string `json:"menuName"` // 菜单名称
	Icon     string `json:"icon"`     // 图标
	Url      string `json:"url"`      // url
}

// LeftMenuVo 左侧菜单vo
type LeftMenuVo struct {
	Id          uint      `json:"id"`          // ID
	MenuName    string    `json:"menuName"`    // 菜单名称
	Icon        string    `json:"icon"`        // 图标
	MenuSvoList []MenuSvo `json:"menuSvoList"` // 菜单列表
}

// ValueVo 权限vo
type ValueVo struct {
	Value string `json:"value"` // 权限
}
