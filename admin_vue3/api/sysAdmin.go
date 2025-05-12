// 用户相关接口
package api

import (
	"admin_vue3/constant"
	"admin_vue3/core"
	. "admin_vue3/core"
	"admin_vue3/global"
	"admin_vue3/model"
	"admin_vue3/result"
	util "admin_vue3/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateSysAdmin 新增用户
// @Summary 新增用户
// @Tags 用户相关接口
// @Produce json
// @Description 新增用户
// @Param data body model.AddSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/add [post]
func CreateSysAdmin(c *gin.Context) {
	// 绑定参数并校验参数必填
	var dto model.AddSysAdminDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 用户名称不能重复
	sysAdminByUsername := GetAdminByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		result.Failed(c, int(result.ApiCode.UsernameAlready), result.ApiCode.GetMessage(result.ApiCode.UsernameAlready))
		return
	}
	// 新增用户
	sysAdmin := model.SysAdmin{
		Username:   dto.Username,
		Nickname:   dto.Nickname,
		Password:   util.EncryptionMd5(dto.Password),
		Phone:      dto.Phone,
		Email:      dto.Email,
		Sex:        dto.Sex,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: util.HTime{Time: time.Now()},
	}
	Db.Create(&sysAdmin)
	// 增加用户角色表信息
	sysAdminExist := GetAdminByUsername(dto.Username)
	var sysAdminRole model.SysAdminRole
	sysAdminRole.AdminId = sysAdminExist.ID
	sysAdminRole.RoleId = dto.RoleId
	Db.Create(&sysAdminRole)
	result.Success(c, true)
}

// GetSysAdmin 根据id查询用户
// @Summary 根据id查询用户
// @Tags 用户相关接口
// @Produce json
// @Description 根据id查询用户
// @Param id query int true "用户id"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/info [get]
func GetSysAdmin(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var sysAdminInfo model.SysAdminInfo
	Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_role.id").
		First(&sysAdminInfo, Id)
	result.Success(c, sysAdminInfo)
}

// UpdateSysAdmin 修改用户
// @Summary 修改用户
// @Tags 用户相关接口
// @Produce json
// @Description 修改用户
// @Param data body model.UpdateSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/update [put]
func UpdateSysAdmin(c *gin.Context) {
	// 绑定参数并校验参数必填
	var dto model.UpdateSysAdminDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 先查询再修改
	var sysAdmin model.SysAdmin
	Db.First(&sysAdmin, dto.Id)
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	if dto.Note != "" {
		sysAdmin.Note = dto.Note
	}
	sysAdmin.Status = dto.Status
	sysAdmin.Sex = dto.Sex
	Db.Save(&sysAdmin)
	// 删除之前的角色再分配新的角色
	var sysAdminRole model.SysAdminRole
	Db.Where("admin_id = ?", dto.Id).Delete(&model.SysAdminRole{})
	sysAdminRole.AdminId = dto.Id
	sysAdminRole.RoleId = dto.RoleId
	Db.Create(&sysAdminRole)
	result.Success(c, true)
}

// DeleteSysAdmin 根据ID删除用户
// @Summary 根据ID删除用户
// @Tags 用户相关接口
// @Produce json
// @Description 根据ID删除用户
// @Param data body model.SysAdminIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/delete [delete]
func DeleteSysAdmin(c *gin.Context) {
	var dto model.SysAdminIdDto
	_ = c.BindJSON(&dto)
	Db.Delete(&model.SysAdmin{}, dto.ID)
	Db.Where("admin_id = ?", dto.ID).Delete(&model.SysAdminRole{})
	result.Success(c, true)
}

// UpdateSysAdminStatus 设置用户状态
// @Summary 设置用户状态
// @Tags 用户相关接口
// @Produce json
// @Description 设置用户状态
// @Param data body model.UpdateSysAdminStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/updateStatus [put]
func UpdateSysAdminStatus(c *gin.Context) {
	var dto model.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	var sysAdmin model.SysAdmin
	Db.First(&sysAdmin, dto.ID)
	sysAdmin.Status = dto.Status
	Db.Save(&sysAdmin)
	result.Success(c, true)
}

// ResetSysAdminPassword 重置密码
// @Summary 重置密码
// @Tags 用户相关接口
// @Produce json
// @Description 重置密码
// @Param data body model.ResetSysAdminPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/updatePassword [put]
func ResetSysAdminPassword(c *gin.Context) {
	var dto model.ResetSysAdminPasswordDto
	_ = c.BindJSON(&dto)
	var sysAdmin model.SysAdmin
	Db.First(sysAdmin, dto.Id)
	sysAdmin.Password = util.EncryptionMd5(dto.Password)
	Db.Save(&sysAdmin)
	result.Success(c, true)
}

// GetSysAdminList 分页查询用户列表
// @Summary 分页查询用户列表
// @Tags 用户相关接口
// @Produce json
// @Description 分页查询用户列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param username query string false "用户名称"
// @Param status query string false "账号启用状态: 1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/list [get]
func GetSysAdminList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("roleName")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	// 提供默认值
	if PageNum < 1 {
		PageNum = 1
	}
	if PageSize < 1 {
		PageSize = 10
	}
	var sysAdminVo []model.SysAdminVo
	var count int64
	curDb := Db.Table("sys_admin").
		Select("sys_admin.*, sys_role.role_name").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_role.id = sys_admin_role.role_id")
	if Username != "" {
		curDb = curDb.Where("sys_admin.username = ?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("sys_admin.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if Status != "" {
		curDb = curDb.Where("sys_admin.status = ?", Status)
	}
	curDb.Count(&count).Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sys_admin.create_time DESC").Find(&sysAdminVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysAdminVo})
}

// Login 登录
// @Summary 登录
// @Tags 用户相关接口
// @Produce json
// @Description 登录
// @Param data body model.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/login [post]
func Login(c *gin.Context) {
	var dto model.LoginDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}

	// 查询用户信息
	sysAdmin := GetAdminByUsername(dto.Username)
	if sysAdmin.Status == constant.StatusDisabled {
		result.Failed(c, int(result.ApiCode.StatusDisabled), result.ApiCode.GetMessage(result.ApiCode.StatusDisabled))
		return
	}
	if sysAdmin.ID > 0 {
		if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
			result.Failed(c, int(result.ApiCode.PasswordIsNotExists), result.ApiCode.GetMessage(result.ApiCode.PasswordIsNotExists))
			return
		}
		// 生成token值
		token, _ := core.GenerateTokenByAdmin(sysAdmin)
		global.Log.Infof("用户token：%s", token)
		// 1. 组装用户基本信息
		userInfo := gin.H{
			"id":         sysAdmin.ID,
			"username":   sysAdmin.Username,
			"nickname":   sysAdmin.Nickname,
			"avatar":     sysAdmin.Icon,
			"email":      sysAdmin.Email,
			"phone":      sysAdmin.Phone,
			"roleName":   getRoleNameByAdminId(sysAdmin.ID), // 获取角色名称
			"createTime": sysAdmin.CreateTime.Format("2006-01-02 15:04:05"),
		}
		// 2. 获取菜单树形结构
		menuTree := GetMenuTreeByAdminId(sysAdmin.ID)

		// 3. 获取权限标识列表
		permissions := GetPermissionsByAdminId(sysAdmin.ID)

		// 4. 返回完整信息
		result.Success(c, gin.H{
			"token":          token,
			"sysAdmin":       userInfo,
			"leftMenuList":   menuTree,
			"permissionList": permissions,
		})
	} else {
		result.Failed(c, int(result.ApiCode.SysAdminIsNotExists), result.ApiCode.GetMessage(result.ApiCode.SysAdminIsNotExists))
		return
	}
	//	result.Success(c, map[string]any{"token": token})
	//}
}

// GetAdminByUsername 根据角色名称查询角色
func GetAdminByUsername(username string) (sysAdmin model.SysAdmin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// getRoleNameByAdminId 根据用户ID获取角色名称
func getRoleNameByAdminId(adminId uint) string {
	var roleName string
	Db.Table("sys_admin_role").
		Select("sys_role.role_name").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_role.id").
		Where("sys_admin_role.admin_id = ?", adminId).
		Pluck("sys_role.role_name", &roleName)
	return roleName
}

// GetMenuTreeByAdminId 根据用户ID获取菜单树
func GetMenuTreeByAdminId(adminId uint) []model.SysMenu {
	var menus []model.SysMenu
	Db.Table("sys_admin_role").
		Select("DISTINCT sys_menu.*").
		Joins("LEFT JOIN sys_role_menu ON sys_admin_role.role_id = sys_role_menu.role_id").
		Joins("LEFT JOIN sys_menu ON sys_role_menu.menu_id = sys_menu.id").
		Where("sys_admin_role.admin_id = ?", adminId).
		Where("sys_menu.menu_status = 1").    // 只查询启用状态的菜单
		Where("sys_menu.menu_type IN (1,2)"). // 只获取目录和菜单类型
		Order("sys_menu.sort ASC").           // 按排序字段升序
		Find(&menus)

	return buildMenuTree(menus, 0) // 从根节点开始构建
}

// buildMenuTree 递归构建菜单树
func buildMenuTree(menus []model.SysMenu, parentId uint) []model.SysMenu {
	var tree []model.SysMenu
	for _, menu := range menus {
		if menu.ParentId == parentId {
			// 递归查找子菜单
			children := buildMenuTree(menus, menu.ID)
			if children != nil {
				menu.Children = children
			}
			// 过滤按钮类型，只保留目录和菜单
			if menu.MenuType != 3 {
				tree = append(tree, menu)
			}
		}
	}
	return tree
}

// GetPermissionsByAdminId 根据用户ID获取权限列表
func GetPermissionsByAdminId(adminId uint) []string {
	var permissions []string
	Db.Table("sys_admin_role").
		Select("DISTINCT sys_menu.value").
		Joins("LEFT JOIN sys_role_menu ON sys_admin_role.role_id = sys_role_menu.role_id").
		Joins("LEFT JOIN sys_menu ON sys_role_menu.menu_id = sys_menu.id").
		Where("sys_admin_role.admin_id = ?", adminId).
		Where("sys_menu.menu_type = 3"). // 只获取按钮类型权限
		Where("sys_menu.value IS NOT NULL").
		Where("sys_menu.value != ''").
		Pluck("sys_menu.value", &permissions)
	return permissions
}
