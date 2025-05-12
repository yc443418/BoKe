package service

import (
	"Traceability/global"
	"Traceability/model"
	"Traceability/model/revo"
	"Traceability/utils"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"log"
	"time"
)

// UserService 用户服务结构体
type UserService struct {
	DB *gorm.DB
}

// NewUserService 创建用户服务实例
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// Register 用户注册服务
func (s *UserService) Register(req revo.RegisterRequest) (*model.User, error) {
	// 检查用户名是否已存在
	var existingUser model.User
	if err := s.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 创建用户对象
	newUser := model.User{
		Username: req.Username,
		Password: req.Password, // 自动加密
		UserType: req.UserType,
		Status:   0, // 默认启用
		ContactInfo: model.Contact{
			Name:  req.ContactInfo.Name,
			Email: req.ContactInfo.Email,
			Phone: req.ContactInfo.Phone,
		},
	}

	// 保存到数据库
	if err := s.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}

	// 创建审计记录
	log := model.AuditLog{
		LogID:       utils.GenerateUUID(),
		UserID:      newUser.UserID,
		RequestType: req.UserType,
		OldValue:    "新用户",
		Status:      0, // 待审核
	}

	// 保存到数据库
	if err := s.DB.Create(&log).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

// Login 用户登录服务
func (s *UserService) Login(req revo.LoginRequest) (string, error) {
	var user model.User
	// 1. 查询用户
	if err := s.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return "", errors.New("用户不存在")
	}

	// 2. 验证密码
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("密码错误")
	}

	// 3. 生成 Token
	claims := utils.NeedClaims{
		UserID:   user.UserID,
		UserType: user.UserType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.CONFIG.Jwt.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(utils.GetExpireTime())),
			NotBefore: jwt.NewNumericDate(time.Now().Add(utils.GetNotBefore())),
		},
	}

	token, err := utils.GenRegisterToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(userID string) (*model.User, error) {
	var user model.User
	// 启用 Debug 模式，打印 SQL
	result := s.DB.Debug().Where("user_id = ?", userID).First(&user)
	if result.Error != nil {
		log.Printf("数据库错误: %v, SQL: %s", result.Error, result.Statement.SQL.String())
		return nil, result.Error
	}
	return &user, nil
}

// SubmitUserTypeRequest 用户提交类型变更申请
func (s *UserService) SubmitUserTypeRequest(userID string, newType string) error {
	// 获取当前用户类型
	var user model.User
	if err := s.DB.First(&user, "user_id = ?", userID).Error; err != nil {
		return err
	}

	// 创建审计记录
	log := model.AuditLog{
		LogID:       utils.GenerateUUID(),
		UserID:      userID,
		RequestType: newType,
		OldValue:    user.UserType,
		Status:      0, // 待审核
	}
	return s.DB.Create(&log).Error
}
