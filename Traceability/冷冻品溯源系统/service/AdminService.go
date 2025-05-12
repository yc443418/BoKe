package service

import (
	"Traceability/model"
	"Traceability/model/revo"
	"Traceability/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{db: db}
}

// GetUserList 获取用户列表（分页）
func (s *AdminService) GetUserList(req revo.UserQueryRequest) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := s.db.Model(&model.User{})

	// 构建查询条件
	if req.Username != "" {
		query = query.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.UserType != "" {
		query = query.Where("user_type = ?", req.UserType)
	}
	if req.Status >= 0 {
		query = query.Where("status = ?", req.Status)
	}

	// 分页处理
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if req.Page > 0 && req.PageSize > 0 {
		offset := (req.Page - 1) * req.PageSize
		query = query.Offset(offset).Limit(req.PageSize)
	}

	// 执行查询
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// GetPendingRequests 获取待审核请求
func (s *AdminService) GetPendingRequests() ([]model.AuditLog, error) {
	var logs []model.AuditLog
	result := s.db.Preload("users").
		Where("status = 0").
		Order("created_at DESC").
		Find(&logs)
	return logs, result.Error
}

// ProcessAuditRequest 处理审核用户请求
func (s *AdminService) ProcessAuditRequest(logID string, adminID string, status int) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 获取审核记录
		var log model.AuditLog
		if err := tx.Preload("users").First(&log, "log_id = ?", logID).Error; err != nil {
			return errors.New("申请记录不存在")
		}

		// 更新审核记录
		updates := map[string]interface{}{
			"status":      status,
			"admin_id":    adminID,
			"reviewed_at": nil, // 初始值
		}

		// 如果是审核通过，设置审核时间
		if status == 1 {
			updates["reviewed_at"] = utils.HTime{Time: time.Now()}

			// 更新用户类型
			if err := tx.Model(&model.User{}).
				Where("user_id = ?", log.UserID).
				Update("user_type", log.RequestType).Error; err != nil {
				return errors.New("用户类型更新失败")
			}
		}

		// 一次性更新所有审核日志字段
		if err := tx.Model(&log).Where("log_id = ?", log.LogID).Updates(updates).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetPendingRequests 获取待审核请求
func (s *AdminService) GetPendingProducts() ([]model.AuditLog, error) {
	var logs []model.AuditLog
	result := s.db.Preload("users").
		Where("status = 0").
		Order("created_at DESC").
		Find(&logs)
	return logs, result.Error
}

// ProcessAuditProducts 处理审核产品请求
func (s *AdminService) ProcessAuditProducts(logID string, adminID string, status int) error {
	return s.db.Transaction(func(tx *gorm.DB) error { // 获取审核记录
		var application revo.ProductApplication
		if err := tx.First(&application, "log_id = ?", logID).Error; err != nil {
			return errors.New("申请记录不存在")
		}

		// 更新审核记录
		updates := map[string]interface{}{
			"status":     status,
			"admin_id":   adminID,
			"updated_at": time.Now(),
		}

		// 如果是审核通过，设置审核时间
		if status == 1 {
			// 更新用户类型
			if err := tx.Model(&model.FrozenProduct{}).
				Where("product_id = ?", application.ProductID).
				Update("status", 1).Error; err != nil {
				return errors.New("用户类型更新失败")
			}
		}

		// 一次性更新所有审核日志字段
		if err := tx.Model(&application).Where("product_id = ?", application.ProductID).Updates(updates).Error; err != nil {
			return err
		}

		return nil
	})
}
