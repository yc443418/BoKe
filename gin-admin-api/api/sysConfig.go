// 参数配置相关接口
// @author chen

package api

import (
	"gin-admin-api/core"
	. "gin-admin-api/core"
	"gin-admin-api/global"
	"gin-admin-api/model"
	"gin-admin-api/result"
	util "gin-admin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// CreateSysConfig 新增参数
// @Summary 新增参数
// @Tags 参数相关接口
// @Produce json
// @Description 新增参数
// @Param data body model.SysConfig true "data"
// @Success 200 {object} result.Result
// @router /api/sysConfig/add [post]
// @Security ApiKeyAuth
func CreateSysConfig(c *gin.Context) {
	var dto model.SysConfig
	_ = c.BindJSON(&dto)
	sysConfig := GetSysConfigByName(dto.Name)
	if sysConfig.ID > 0 {
		result.Failed(c, int(result.ApiCode.ConfigNameAlreadyExists), result.ApiCode.GetMessage(result.ApiCode.ConfigNameAlreadyExists))
		return
	}
	addSysConfig := model.SysConfig{
		Name:        dto.Name,
		ConfigKey:   dto.ConfigKey,
		ConfigValue: dto.ConfigValue,
		Sort:        dto.Sort,
		Remark:      dto.Remark,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	Db.Create(&addSysConfig)
	// 存redis缓存
	core.RedisDb.Set(global.Ctx, dto.ConfigKey, dto.ConfigValue, 0)
	result.Success(c, true)
}

// GetSysConfigList 分页查询参数列表
// @Summary 分页查询参数列表
// @Tags 参数相关接口
// @Produce json
// @Description 分页查询参数列表
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param name query string false "参数名称"
// @Success 200 {object} result.Result
// @router /api/sysConfig/list [get]
// @Security ApiKeyAuth
func GetSysConfigList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Name := c.Query("name")
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	var sysConfig []model.SysConfig
	var count int64
	curDb := Db.Table("sys_config")
	if Name != "" {
		curDb = curDb.Where("name = ?", Name)
	}
	curDb.Count(&count).Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("sort DESC").Find(&sysConfig)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": sysConfig})
}

// GetSysConfig 根据id查询参数
// @Summary 根据id查询参数
// @Tags 参数相关接口
// @Produce json
// @Description 根据id查询参数
// @Param id query int true "参数id"
// @Success 200 {object} result.Result
// @router /api/sysConfig/info [get]
// @Security ApiKeyAuth
func GetSysConfig(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	var sysConfig model.SysConfig
	Db.First(&sysConfig, Id)
	result.Success(c, sysConfig)
}

// UpdateSysConfig 修改参数
// @Summary 修改参数
// @Tags 参数相关接口
// @Produce json
// @Description 修改参数
// @Param data body model.SysConfig  true "data"
// @Success 200 {object} result.Result
// @router /api/sysConfig/update [put]
// @Security ApiKeyAuth
func UpdateSysConfig(c *gin.Context) {
	var dto model.SysConfig
	_ = c.BindJSON(&dto)
	var sysConfig model.SysConfig
	Db.First(&sysConfig, dto.ID)
	// 删除之前的缓存
	core.RedisDb.Del(global.Ctx, sysConfig.ConfigKey)
	// 保存新的值
	sysConfig.Name = dto.Name
	sysConfig.ConfigKey = dto.ConfigKey
	sysConfig.ConfigValue = dto.ConfigValue
	sysConfig.Sort = dto.Sort
	sysConfig.Remark = dto.Remark
	Db.Save(&sysConfig)
	core.RedisDb.Set(global.Ctx, dto.ConfigKey, dto.ConfigValue, 0)
	result.Success(c, true)
}

// DeleteSysConfig 删除参数
// @Summary 删除参数
// @Tags 参数相关接口
// @Produce json
// @Description 删除参数
// @Param data body model.SysConfigIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysConfig/delete [delete]
// @Security ApiKeyAuth
func DeleteSysConfig(c *gin.Context) {
	var dto model.SysConfigIdDto
	_ = c.BindJSON(&dto)
	// 删除redis缓存
	var sysConfig model.SysConfig
	Db.First(&sysConfig, dto.Id)
	core.RedisDb.Del(global.Ctx, sysConfig.ConfigKey)
	// 删除数据
	Db.Table("sys_config").Delete(&sysConfig, dto.Id)
	result.Success(c, true)
}

// Refresh 刷新
// @Summary 刷新
// @Tags 参数相关接口
// @Produce json
// @Description 刷新
// @Success 200 {object} result.Result
// @router /api/sysConfig/refresh [delete]
// @Security ApiKeyAuth
func Refresh(c *gin.Context) {
	// 先清除缓存
	keys, err := core.RedisDb.Keys(global.Ctx, "*").Result()
	if err != nil {
		return
	}
	global.Log.Infof("所有的缓存key值：%s", keys)
	for _, key := range keys {
		core.RedisDb.Del(global.Ctx, key)
	}
	// 添加的新的缓存
	var sysConfig []model.SysConfig
	Db.Table("sys_config").Find(&sysConfig)
	for _, config := range sysConfig {
		core.RedisDb.Set(global.Ctx, config.ConfigKey, config.ConfigValue, 0)
	}
	result.Success(c, true)
}

// GetSysConfigByName 根据名称查询
func GetSysConfigByName(name string) (sysConfig model.SysConfig) {
	Db.Where("name = ?", name).First(&sysConfig)
	return sysConfig
}
