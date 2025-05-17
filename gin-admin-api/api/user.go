// 用户相关接口
// @author chen

package api

import (
	"gin-admin-api/constant"
	. "gin-admin-api/core"
	"gin-admin-api/model"
	"gin-admin-api/result"
	util "gin-admin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// GetUserList 分页查询前台用户
// @Summary 分页查询前台用户
// @Tags 前台用户相关接口
// @Produce json
// @Description 分页查询前台用户
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param username query string false "用户账号"
// @Param status query string false "帐号状态：1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/user/list [get]
// @Security ApiKeyAuth
func GetUserList(c *gin.Context) {
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
	var userVo []model.UserVo
	var count int64
	curDb := Db.Table("user u").
		Select("u.*, a.quantity").
		Joins("LEFT JOIN (SELECT user_id, count(id) AS quantity FROM b_article GROUP BY user_id) a ON u.id = a.user_id")
	if Username != "" {
		curDb = curDb.Where("u.username = ?", Username)
	}
	if Status != "" {
		curDb = curDb.Where("u.status = ?", Status)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("u.create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("u.create_time DESC").Find(&userVo)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize, "pageNum": PageNum, "list": userVo})
}

// UpdateUserStatus 用户状态启用/停用
// @Summary 用户状态启用/停用
// @Tags 前台用户相关接口
// @Produce json
// @Description 用户状态启用/停用
// @Param data body model.UpdateSysRoleStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/updateStatus [put]
// @Security ApiKeyAuth
func UpdateUserStatus(c *gin.Context) {
	var dto model.UpdateUserStatusDto
	_ = c.BindJSON(&dto)
	var user model.User
	Db.First(&user, dto.Id)
	user.Status = dto.Status
	Db.Save(&user)
	result.Success(c, true)
}

// SendMessage 给用户发送消息
// @Summary 给用户发送消息
// @Tags 前台用户相关接口
// @Produce json
// @Description 给用户发送消息
// @Param data body model.SendMessageDto true "data"
// @Success 200 {object} result.Result
// @router /api/user/send/message [post]
// @Security ApiKeyAuth
func SendMessage(c *gin.Context) {
	var dto model.SendMessageDto
	_ = c.BindJSON(&dto)
	addUserMessage := model.UserMessage{
		ReceivedUserId: dto.Id,
		MessageContent: dto.MessageContent,
		MessageType:    constant.SystemMessage,
		CreateTime:     util.HTime{Time: time.Now()},
		Status:         constant.NotReadStatus,
	}
	Db.Create(&addUserMessage)
	result.Success(c, true)
}
