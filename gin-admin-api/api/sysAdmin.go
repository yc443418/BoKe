// 用户相关接口
// @author chen

package api

import (
	"gin-admin-api/constant"
	"gin-admin-api/core"
	. "gin-admin-api/core"
	"gin-admin-api/global"
	"gin-admin-api/model"
	"gin-admin-api/result"
	util "gin-admin-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

// CreateSysAdmin 新增用户
// @Summary 新增用户
// @Tags 用户相关接口
// @Produce json
// @Description 新增用户
// @Param data body model.AddSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/add [post]
// @Security ApiKeyAuth
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
	sysAdminByUsername := GetSysAdminByUsername(dto.Username)
	if sysAdminByUsername.ID > 0 {
		result.Failed(c, int(result.ApiCode.UsernameAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.UsernameAlreadyExists))
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
	sysAdminExist := GetSysAdminByUsername(dto.Username)
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
// @Security ApiKeyAuth
func GetSysAdmin(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var sysAdminInfo model.SysAdminInfo
	Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id = sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_Admin_role.role_id = sys_role.id").
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
// @Security ApiKeyAuth
func UpdateSysAdmin(c *gin.Context) {
	// 绑定参数并校验参数必填
	var dto model.UpdateSysAdminDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 先查询在修改
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
	// 删除之前的角色在分配新的角色
	var sysAdminRole model.SysAdminRole
	Db.Where("admin_id = ?", dto.Id).Delete(&model.SysAdminRole{})
	sysAdminRole.AdminId = dto.Id
	sysAdminRole.RoleId = dto.RoleId
	Db.Create(&sysAdminRole)
	result.Success(c, true)
}

// DeleteSysAdmin 根据id删除用户
// @Summary 根据id删除用户
// @Tags 用户相关接口
// @Produce json
// @Description 根据id删除用户
// @Param data body model.SysAdminIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/delete [delete]
// @Security ApiKeyAuth
func DeleteSysAdmin(c *gin.Context) {
	var dto model.SysAdminIdDto
	_ = c.BindJSON(&dto)
	Db.Delete(&model.SysAdmin{}, dto.Id)
	Db.Where("admin_id = ?", dto.Id).Delete(&model.SysAdminRole{})
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
// @Security ApiKeyAuth
func UpdateSysAdminStatus(c *gin.Context) {
	var dto model.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	var sysAdmin model.SysAdmin
	Db.First(&sysAdmin, dto.Id)
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
// @Security ApiKeyAuth
func ResetSysAdminPassword(c *gin.Context) {
	var dto model.ResetSysAdminPasswordDto
	_ = c.BindJSON(&dto)
	var sysAdmin model.SysAdmin
	Db.First(&sysAdmin, dto.Id)
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
// @Param status query string false "帐号启用状态：1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/list [get]
// @Security ApiKeyAuth
func GetSysAdminList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var sysAdminVo []model.SysAdminVo
	var count int64
	curDb := Db.Table("sys_admin").
		Select("sys_admin.*, sys_role.role_name").
		Joins("LEFT JOIN sys_Admin_role ON sys_admin.id = sys_admin_role.admin_id").
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
	// 绑定json参数并校验参数
	var dto model.LoginDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	// 查询用户信息
	sysAdmin := GetSysAdminByUsername(dto.Username)
	if sysAdmin.Status == constant.StatusDisabled {
		result.Failed(c, int(result.ApiCode.StatusDisabled), result.ApiCode.GetMessage(result.ApiCode.StatusDisabled))
		return
	}
	if sysAdmin.ID > 0 {
		// 校验密码是否正确
		if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
			result.Failed(c, int(result.ApiCode.PasswordNotTrue), result.ApiCode.GetMessage(result.ApiCode.PasswordNotTrue))
			return
		}
		// 生成token
		token, _ := core.GenerateTokenByAdmin(sysAdmin)
		global.Log.Infof("用户token: %s", token)
		// 组装用户信息
		sysAdminVo := model.SysAdminVo{
			ID:         sysAdmin.ID,
			Username:   sysAdmin.Username,
			Nickname:   sysAdmin.Nickname,
			Status:     sysAdmin.Status,
			Icon:       sysAdmin.Icon,
			Sex:        sysAdmin.Sex,
			Email:      sysAdmin.Email,
			Phone:      sysAdmin.Phone,
			Note:       sysAdmin.Note,
			CreateTime: sysAdmin.CreateTime,
		}
		// 左侧菜单数据
		var leftMenuVo []model.LeftMenuVo
		leftMenuList := QueryLeftMenuList(sysAdmin.ID)
		for _, vo := range leftMenuList {
			menuSvoList := QueryMenuVoList(sysAdmin.ID, vo.Id)
			item := model.LeftMenuVo{}
			item.MenuSvoList = menuSvoList
			item.Id = vo.Id
			item.MenuName = vo.MenuName
			item.Icon = vo.Icon
			leftMenuVo = append(leftMenuVo, item)
		}
		// 权限列表
		permissionList := QueryPermissionList(sysAdmin.ID)
		var stringList = make([]string, 0, 10)
		for _, vo := range permissionList {
			stringList = append(stringList, vo.Value)
		}
		result.Success(c, map[string]any{"token": token, "sysAdmin": sysAdminVo, "leftMenuList": leftMenuVo, "permissionList": stringList})
	} else {
		result.Failed(c, int(result.ApiCode.SysAdminIsNotExists), result.ApiCode.GetMessage(result.ApiCode.SysAdminIsNotExists))
		return
	}
}

// UpdatePersonal 修改个人信息
// @Summary 修改个人信息
// @Tags 用户相关接口
// @Produce json
// @Description 修改个人信息
// @Param data body model.UpdatePersonalDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/updatePersonal [put]
// @Security ApiKeyAuth
func UpdatePersonal(c *gin.Context) {
	var dto model.UpdatePersonalDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	admin, _ := core.GetAdmin(c)
	var sysAdmin model.SysAdmin
	Db.First(&sysAdmin, admin.ID)
	if dto.Icon != "" {
		sysAdmin.Icon = dto.Icon
	}
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
	Db.Save(&sysAdmin)
	result.Success(c, sysAdmin)
}

// UpdatePersonalPassword 修改密码
// @Summary 修改密码
// @Tags 用户相关接口
// @Produce json
// @Description 修改密码
// @Param data body model.UpdatePersonalPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysAdmin/updatePersonalPassword [put]
// @Security ApiKeyAuth
func UpdatePersonalPassword(c *gin.Context) {
	var dto model.UpdatePersonalPasswordDto
	_ = c.BindJSON(&dto)
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissParameter), result.ApiCode.GetMessage(result.ApiCode.MissParameter))
		return
	}
	admin, _ := core.GetAdmin(c)
	dto.Id = admin.ID
	sysAdminExist := GetSysAdminByUsername(admin.Username)
	if sysAdminExist.Password != util.EncryptionMd5(dto.OldPassword) {
		result.Failed(c, int(result.ApiCode.PasswordNotTrue), result.ApiCode.GetMessage(result.ApiCode.PasswordNotTrue))
		return
	}
	if dto.NewPassword != dto.ResetPassword {
		result.Failed(c, int(result.ApiCode.ResetPasswordNotTrue), result.ApiCode.GetMessage(result.ApiCode.ResetPasswordNotTrue))
		return
	}
	dto.NewPassword = util.EncryptionMd5(dto.NewPassword)
	UpdatePwd(dto)
	result.Success(c, true)
}

// QueryMenuVoList 当前登录用户的左侧菜单级列表
func QueryMenuVoList(AdminId, MenuId uint) (menuSvo []model.MenuSvo) {
	const status, menuStatus, menuType = 1, 1, 2
	Db.Table("sys_menu sm").
		Select("sm.menu_name, sm.icon, sm.url").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sm.parent_id = ?", MenuId).
		Where("sa.id = ?", AdminId).
		Order("sm.sort").
		Scan(&menuSvo)
	return menuSvo
}

// QueryLeftMenuList 当前登录用户左侧菜单列表
func QueryLeftMenuList(Id uint) (leftMenuVo []model.LeftMenuVo) {
	const status, menuStatus, menuType uint = 1, 1, 1
	Db.Table("sys_menu sm").
		Select("sm.id, sm.menu_name, sm.url, sm.icon").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Order("sm.sort").
		Scan(&leftMenuVo)
	return leftMenuVo
}

// QueryPermissionList 当前登录用户的权限列表
func QueryPermissionList(Id uint) (valueVo []model.ValueVo) {
	const status, menuStatus, menuType uint = 1, 1, 3
	Db.Table("sys_menu sm").
		Select("sm.value").
		Joins("LEFT JOIN sys_role_menu srm ON sm.id = srm.menu_id").
		Joins("LEFT JOIN sys_role sr On sr.id = srm.role_id").
		Joins("LEFT JOIN sys_admin_role sar ON sar.role_id = sr.id").
		Joins("LEFT JOIN sys_admin sa ON sa.id = sar.admin_id").
		Where("sr.status = ?", status).
		Where("sm.menu_status = ?", menuStatus).
		Where("sm.menu_type = ?", menuType).
		Where("sa.id = ?", Id).
		Scan(&valueVo)
	return valueVo
}

// GetSysAdminByUsername 根据用户名称查询用户
func GetSysAdminByUsername(username string) (sysAdmin model.SysAdmin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// UpdatePwd 修改个人密码
func UpdatePwd(dto model.UpdatePersonalPasswordDto) (sysAdmin model.SysAdmin) {
	Db.First(&sysAdmin, dto.Id)
	sysAdmin.Password = dto.NewPassword
	Db.Save(&sysAdmin)
	return sysAdmin
}
