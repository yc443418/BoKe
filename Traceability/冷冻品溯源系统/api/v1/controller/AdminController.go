package controller

import (
	ErrMsg "Traceability/config"
	"Traceability/model"
	"Traceability/model/revo"
	"Traceability/service"
	"Traceability/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{
		adminService: service.NewAdminService(db),
	}
}

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Description 管理员获取用户列表（分页）
// @Tags 管理员管理
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param username query string false "用户名"
// @Param user_type query string false "用户类型"
// @Param status query int false "状态"
// @Success 200 {object} utils.Result
// @Router /admin/users [get]
func (c *AdminController) GetUserList(ctx *gin.Context) {
	var req revo.UserQueryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_AUTH_FAILED)
		return
	}

	users, total, err := c.adminService.GetUserList(req)
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_DATABASE_ERROR)
		return
	}

	response := buildUserListResponse(users)
	utils.ResultReturnOK(ctx, gin.H{
		"total": total,
		"users": response,
	}, ErrMsg.SUCCEE)
}

// GetPendingRequests 获取待审核请求
// @Summary 获取待审核请求
// @Description 管理员获取所有未处理的申请
// @Tags 管理员管理
// @Security ApiKeyAuth
// @Success 200 {object} utils.Result
// @Router /admin/pending [get]
func (c *AdminController) GetPendingRequests(ctx *gin.Context) {
	logs, err := c.adminService.GetPendingRequests()
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_DATABASE_ERROR)
		return
	}
	utils.ResultReturnOK(ctx, logs, ErrMsg.SUCCEE)
}

// ProcessAudit 处理审核用户请求
// @Summary 处理审核请求
// @Description 管理员处理用户申请
// @Tags 管理员管理
// @Security ApiKeyAuth
// @Param log_id path string true "申请记录ID"
// @Param request body revo.AuditProcessRequest true "审核请求"
// @Success 200 {object} utils.Result
// @Router /admin/{log_id} [put]
func (c *AdminController) ProcessAudit(ctx *gin.Context) {
	logID := ctx.Param("log_id")
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_AUTH_FAILED)
	}
	var req revo.AuditProcessRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR)
		return
	}

	if err := c.adminService.ProcessAuditRequest(logID, claims.UserID, req.Status); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_DATABASE_ERROR)
		return
	}
	utils.ResultReturnOK(ctx, nil, ErrMsg.SUCCEE)
}

// ProcessProducts 处理审核产品请求
// @Summary 处理审核产品请求
// @Description 管理员处理产品申请
// @Tags 管理员管理
// @Security ApiKeyAuth
// @Param log_id path string true "申请记录ID"
// @Param request body revo.AuditProcessRequest true "审核请求"
// @Success 200 {object} utils.Result
// @Router /admin/product/{log_id} [put]
func (c *AdminController) ProcessProducts(ctx *gin.Context) {
	logID := ctx.Param("log_id")
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_AUTH_FAILED)
	}
	if claims.UserType != "admin" {
		utils.ResultReturnErr(ctx, "权限不足", ErrMsg.ERROR_AUTH_FAILED)
		return
	}
	var req revo.FrozenProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR)
		return
	}

	if err := c.adminService.ProcessAuditProducts(logID, claims.UserID, req.Status); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_DATABASE_ERROR)
		return
	}
	utils.ResultReturnOK(ctx, nil, ErrMsg.SUCCEE)
}

// // 构建用户列表响应
func buildUserListResponse(users []model.User) []revo.UserResponse {
	var response []revo.UserResponse
	for _, user := range users {
		response = append(response, revo.UserResponse{
			UserID:    user.UserID,
			Username:  user.Username,
			UserType:  user.UserType,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			Contact: revo.Contact{
				Name:  user.ContactInfo.Name,
				Email: user.ContactInfo.Email,
				Phone: user.ContactInfo.Phone,
			},
		})
	}
	return response
}
