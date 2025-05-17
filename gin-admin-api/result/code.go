// 状态码和状态信息定义
// @author chen

package result

// Codes 定义状态
type Codes struct {
	Message                 map[uint]string
	Success                 uint
	Failed                  uint
	NoAuth                  uint
	AuthFormatError         uint
	InvalidToken            uint
	SysMenuIsExist          uint
	DelSysMenuFailed        uint
	NorDeleteMenu           uint
	RoleAlreadyExists       uint
	DelSysRoleFailed        uint
	MissParameter           uint
	UsernameAlreadyExists   uint
	StatusDisabled          uint
	SysAdminIsNotExists     uint
	PasswordNotTrue         uint
	FileUploadError         uint
	ResetPasswordNotTrue    uint
	ConfigNameAlreadyExists uint
	BTagsHasExist           uint
	BCategoryIsExist        uint
	CategoryHasArticle      uint
	DeletedCannotOperate    uint
}

// ApiCode 状态码
var ApiCode = &Codes{
	Success:                 200,
	Failed:                  501,
	NoAuth:                  401,
	AuthFormatError:         402,
	InvalidToken:            403,
	SysMenuIsExist:          600,
	DelSysMenuFailed:        601,
	NorDeleteMenu:           602,
	RoleAlreadyExists:       603,
	DelSysRoleFailed:        604,
	MissParameter:           605,
	UsernameAlreadyExists:   606,
	StatusDisabled:          607,
	SysAdminIsNotExists:     608,
	PasswordNotTrue:         609,
	FileUploadError:         610,
	ResetPasswordNotTrue:    611,
	ConfigNameAlreadyExists: 612,
	BTagsHasExist:           613,
	BCategoryIsExist:        614,
	CategoryHasArticle:      615,
	DeletedCannotOperate:    616,
}

// 状态信息初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Success:                 "成功",
		ApiCode.Failed:                  "失败",
		ApiCode.NoAuth:                  "无权限",
		ApiCode.AuthFormatError:         "请求头中的auth格式错误",
		ApiCode.InvalidToken:            "无效的token,请重新登录",
		ApiCode.SysMenuIsExist:          "菜单名称已存在，请重新输入！",
		ApiCode.DelSysMenuFailed:        "菜单已分配不能删除",
		ApiCode.NorDeleteMenu:           "存在子菜单不能删除",
		ApiCode.RoleAlreadyExists:       "角色名称或角色key已存在，请重新输入！",
		ApiCode.DelSysRoleFailed:        "角色已分配不能删除",
		ApiCode.MissParameter:           "缺少必填参数",
		ApiCode.UsernameAlreadyExists:   "用户名称已存在，请重新输入！",
		ApiCode.StatusDisabled:          "您的账号已被禁用，请联系管理员",
		ApiCode.SysAdminIsNotExists:     "用户不存在",
		ApiCode.PasswordNotTrue:         "密码不正确",
		ApiCode.FileUploadError:         "文件上传错误",
		ApiCode.ResetPasswordNotTrue:    "两次输入密码不一样，请重新输入！",
		ApiCode.ConfigNameAlreadyExists: "参数名称已存在，请重新输入！",
		ApiCode.BTagsHasExist:           "该标签已存在，请重新输入。",
		ApiCode.BCategoryIsExist:        "文章分类名称已存在，请重新输入。",
		ApiCode.CategoryHasArticle:      "该分类下有文章不能删除。",
		ApiCode.DeletedCannotOperate:    "该文章已删除不能操作",
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
