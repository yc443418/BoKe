package api

import (
	. "admin_vue3/core"
	"admin_vue3/model"
	"admin_vue3/result"
	"admin_vue3/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateSysRole 新增角色
// @Summary 新增角色
// @Tags 角色相关接口
// @Produce json
// @Description 新增角色
// @Param data body model.AddSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/add [post]
func CreateSysRole(c *gin.Context) {
	var dto model.AddSysRoleDto
	_ = c.BindJSON(&dto)
	sysRoleByName := GetSysMenuByRole(dto.RoleName)
	// 查询用户名称是否存在
	if sysRoleByName.ID > 0 {
		result.Failed(c, int(result.ApiCode.RoleAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.RoleAlreadyExists))
		return
	}
	// 查询用户key是否存在
	sysRoleByKey := GetSysMenuByKey(dto.RoleKey)
	if sysRoleByKey.ID > 0 {
		result.Failed(c, int(result.ApiCode.RoleAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.RoleAlreadyExists))
		return
	}
	addSysRole := model.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Description: dto.Description,
		Status:      dto.Status,
		CreateTime:  utils.HTime{Time: time.Now()},
	}
	tx := Db.Create(&addSysRole)
	if tx.RowsAffected > 0 {
		result.Success(c, true)
		return
	} else {
		result.Failed(c, int(result.ApiCode.Failed), result.ApiCode.GetMessage(result.ApiCode.Failed))
		return
	}
}

// GetSysRoleList 分页查询角色列表
// @Summary 分页查询角色列表
// @Tags 角色相关接口
// @Produce json
// @Description 分页查询角色列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param roleName query string false "角色名称"
// @Param status query string false "账号启用状态: 1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysRole/list [get]
func GetSysRoleList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	RoleName := c.Query("roleName")
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
	var sysRole []model.SysRole
	var count int64
	curDb := Db.Table("sys_role")
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if Status != "" {
		curDb = curDb.Where("status = ?", Status)
	}
	curDb.Count(&count).Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&sysRole)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysRole})
}

// GetSysRole 根据id查询角色
// @Summary 根据id查询角色
// @Tags 角色相关接口
// @Produce json
// @Description 根据id查询角色
// @Param id query int true "角色id"
// @Success 200 {object} result.Result
// @router /api/sysRole/info [get]
func GetSysRole(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var sysRole model.SysRole
	Db.First(&sysRole, Id)
	result.Success(c, sysRole)
}

// UpdateSysRole 修改角色
// @Summary 修改角色
// @Tags 角色相关接口
// @Produce json
// @Description 修改角色
// @Param data body model.UpdateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/update [put]
func UpdateSysRole(c *gin.Context) {
	var dto model.UpdateSysRoleDto
	_ = c.BindJSON(&dto)
	var sysRole model.SysRole
	Db.First(&sysRole, dto.ID)
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	sysRole.Status = dto.Status
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	Db.Save(&sysRole)
	result.Success(c, true)
}

// DeleteSysRole 根据ID删除角色
// @Summary 根据ID删除角色
// @Tags 角色相关接口
// @Produce json
// @Description 根据ID删除角色
// @Param data body model.SysRoleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/delete [delete]
func DeleteSysRole(c *gin.Context) {
	var dto model.SysRoleIdDto
	_ = c.BindJSON(&dto)
	sysRoleById := GetSysAdminRole(dto.ID)
	if sysRoleById.RoleId > 0 {
		result.Failed(c, int(result.ApiCode.DelSysRoleFailed), result.ApiCode.GetMessage(result.ApiCode.DelSysRoleFailed))
		return
	}
	Db.Table("sys_role").Delete(&model.SysRole{}, dto.ID)
	Db.Table("sys_role_menu").Where("role_id = ?", dto.ID).Delete(&model.SysRoleMenu{})
	result.Success(c, true)
}

// UpdateSysRoleStatus 设置角色状态
// @Summary 设置角色状态
// @Tags 角色相关接口
// @Produce json
// @Description 设置角色状态
// @Param data body model.UpdateSysRoleStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/updateStatus [put]
func UpdateSysRoleStatus(c *gin.Context) {
	var dto model.UpdateSysRoleStatusDto
	_ = c.BindJSON(&dto)
	var sysRole model.SysRole
	Db.First(&sysRole, dto.ID)
	sysRole.Status = dto.Status
	Db.Save(&sysRole)
	result.Success(c, true)
}

// GetSysRoleVoList 查询角色下拉列表
// @Summary 查询角色下拉列表
// @Tags 角色相关接口
// @Produce json
// @Description 查询角色下拉列表
// @Success 200 {object} result.Result
// @router /api/sysRole/vo/list [get]
func GetSysRoleVoList(c *gin.Context) {
	var sysRoleVo []model.SysRoleVo
	Db.Table("sys_role").Select("id, role_name").Scan(&sysRoleVo)
	result.Success(c, sysRoleVo)
}

// QueryRoleMenuIdList 根据角色id查询菜单数据
// @Summary 根据角色id查询菜单数据
// @Tags 角色相关接口
// @Produce json
// @Description 根据角色id查询菜单数据
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/sysRole/vo/idList [get]
func QueryRoleMenuIdList(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var idVo []model.IdVo
	const menuType = 3
	Db.Table("sys_menu sm").
		Select("sm.id").Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id").
		Where("sm.menu_type = ?", menuType).
		Where("sr.id = ?", Id).
		Scan(&idVo)
	var idList = make([]int, 0, 10)
	for _, id := range idVo {
		idList = append(idList, id.ID)
	}
	result.Success(c, idList)
}

// AssignPermissions 分配权限
// @Summary 分配权限
// @Tags 角色相关接口
// @Produce json
// @Description 分配权限
// @Param data body model.RoleMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysRole/assignPermissions [put]
func AssignPermissions(c *gin.Context) {
	var roleMenu model.RoleMenuDto
	_ = c.BindJSON(&roleMenu)
	Db.Table("sys_role_menu").
		Where("role_id = ?", roleMenu.ID).
		Delete(&model.SysMenu{})
	for _, value := range roleMenu.MenuIds {
		var sysRoleMenu model.SysRoleMenu
		sysRoleMenu.RoleId = roleMenu.ID
		sysRoleMenu.MenuId = value
		Db.Create(&sysRoleMenu)
	}
	result.Success(c, true)
}

// GetSysMenuByRole 根据角色名称查询角色
func GetSysMenuByRole(roleName string) (sysRole model.SysRole) {
	Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

// GetSysMenuByKey 根据角色key查询角色
func GetSysMenuByKey(roleKey string) (sysRole model.SysRole) {
	Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}

// GetSysAdminRole 查询是否分配菜单
func GetSysAdminRole(id uint) (sysAdminRole model.SysAdminRole) {
	Db.Where("role_id = ?", id).First(&sysAdminRole)
	return sysAdminRole
}
