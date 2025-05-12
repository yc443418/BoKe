package controller

import (
	ErrMsg "Traceability/config"
	"Traceability/model/revo"
	"Traceability/service"
	"Traceability/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		userService: service.NewUserService(db),
	}
}

// Register 用户注册接口
// @Summary 用户注册
// @Description 新用户注册接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body revo.RegisterRequest true "注册请求"
// @Success 200 {object} utils.Result
// @Router /register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req revo.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_AUTH_FAILED)
		return
	}

	user, err := c.userService.Register(req)
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_DATABASE_ERROR)
		return
	}

	response := revo.UserResponse{
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
	}
	utils.ResultReturnOK(ctx, response, ErrMsg.SUCCEE)
}

// Login 用户登录接口
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body revo.LoginRequest true "登录请求"
// @Success 200 {object} utils.Result
// @Router /login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req revo.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_AUTH_FAILED)
		return
	}

	token, err := c.userService.Login(req)
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_AUTH_FAILED)
		return
	}

	utils.ResultReturnOK(ctx, gin.H{"token": token}, ErrMsg.SUCCEE)
}

// GetUserInfo 获取用户信息接口
// @Summary 获取用户信息
// @Description 获取当前登录用户信息
// @Tags 用户管理
// @Param Authorization header string true "Bearer Token"
// @Security ApiKeyAuth
// @Success 200 {object} utils.Result
// @Router /user/info [get]
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		log.Printf("Token解析失败: %v", err)
		utils.ResultReturnErr(ctx, "无效的Token", ErrMsg.ERROR) // 更友好的错误提示
		return
	}

	// 使用 JWT 中的 UserID 获取用户信息
	userID := claims.UserID
	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_USER_NOT_EXIST)
		return
	}

	// 构建响应
	response := revo.UserResponse{
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
	}

	// 返回用户信息
	utils.ResultReturnOK(ctx, response, ErrMsg.SUCCEE)
}

// SubmitUserTypeRequest 用户提交类型变更申请
// @Summary 提交用户类型变更
// @Description 普通用户提交用户类型变更申请
// @Tags 用户管理
// @Security ApiKeyAuth
// @Param request body revo.UserTypeRequest true "变更请求"
// @Success 200 {object} utils.Result
// @Router /user/type-request [post]
func (c *UserController) SubmitUserTypeRequest(ctx *gin.Context) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_AUTH_FAILED)
	}
	var req revo.UserTypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR)
		return
	}

	if err := c.userService.SubmitUserTypeRequest(claims.UserID, req.RequestType); err != nil {
		utils.ResultReturnErr(ctx, err.Error(), ErrMsg.ERROR_DATABASE_ERROR)
		return
	}
	utils.ResultReturnOK(ctx, nil, ErrMsg.SUCCEE)
}
