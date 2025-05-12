// 状态码和状态信息定义
// @author chen
package result

// Codes 定义状态
type Codes struct {
	Message             map[uint]string
	Success             uint
	Failed              uint
	SysMenuIsExist      uint
	DelSysMenuFailed    uint
	NorDeletMenu        uint
	RoleAlreadyExists   uint
	DelSysRoleFailed    uint
	MissParameter       uint
	UsernameAlready     uint
	StatusDisabled      uint
	SysAdminIsNotExists uint
	PasswordIsNotExists uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	Success:             200,
	Failed:              501,
	SysMenuIsExist:      600,
	DelSysMenuFailed:    601,
	NorDeletMenu:        602,
	RoleAlreadyExists:   603,
	DelSysRoleFailed:    604,
	MissParameter:       605,
	UsernameAlready:     606,
	StatusDisabled:      607,
	SysAdminIsNotExists: 608,
	PasswordIsNotExists: 609,
}

// 状态信息初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success:             "成功",
		ApiCode.Failed:              "失败",
		ApiCode.SysMenuIsExist:      "菜单名称已存在，请重新输入",
		ApiCode.DelSysMenuFailed:    "菜单已分配不能删除",
		ApiCode.NorDeletMenu:        "存在子菜单不能删除",
		ApiCode.RoleAlreadyExists:   "角色名称或角色key已注册,请重新输入",
		ApiCode.DelSysRoleFailed:    "角色已分配不能删除",
		ApiCode.MissParameter:       "必填参数未填",
		ApiCode.UsernameAlready:     "用户名称已存在，请重新输入",
		ApiCode.StatusDisabled:      "您的账号已停用，请联系管理员",
		ApiCode.SysAdminIsNotExists: "用户名错误，请重新输入",
		ApiCode.PasswordIsNotExists: "密码错误，请重新输入",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
