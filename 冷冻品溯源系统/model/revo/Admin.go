package revo

import "time"

// UserQueryRequest 用户查询请求
type UserQueryRequest struct {
	Page     int    `form:"page" binding:"min=1"`
	PageSize int    `form:"page_size" binding:"min=1,max=100"`
	Username string `form:"username"`
	UserType string `form:"user_type"`
	Status   int    `form:"status"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Total int64          `json:"total"`
	Users []UserResponse `json:"users"`
}

// UserTypeRequest 用户类型变更请求
type UserTypeRequest struct {
	RequestType string `json:"request_type" binding:"required,oneof=factory dealer supervision consumer"`
}

// AuditProcessRequest 审核处理请求
type AuditProcessRequest struct {
	Status int `json:"status" binding:"required,oneof=1 2"`
}

// AuditLogResponse 审核记录响应
type AuditLogResponse struct {
	LogID       string `json:"log_id"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	RequestType string `json:"request_type"`
	OldValue    string `json:"old_value"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"created_at"`
	ReviewedAt  string `json:"reviewed_at,omitempty"`
}

type ProductApplication struct {
	LogId     string    `gorm:"primaryKey;type:varchar(36)" json:"log_id"`
	ProductID string    `gorm:"type:varchar(36);index" json:"product_id"`
	FactoryID string    `gorm:"type:varchar(36);unique;constraint:OnDelete:CASCADE" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:null" json:"reviewed_at"`
	AdminID   string    `gorm:"type:varchar(36)" json:"admin_id"`
	Status    int       `gorm:"type:int;default:0" json:"status"` //默认1为同意上架 0为待审核  2为下架
}
