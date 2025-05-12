package model

import (
	"time"
)

type AuditLog struct {
	LogID       string    `gorm:"primaryKey;type:varchar(36)" json:"log_id"`
	UserID      string    `gorm:"type:varchar(36);index;not null;column:user_id" json:"user_id"`
	RequestType string    `gorm:"type:varchar(50);not null" json:"request_type"`
	OldValue    string    `gorm:"type:text" json:"old_value"` // 旧值
	Status      int       `gorm:"type:int;default:0" json:"status"`
	AdminID     string    `gorm:"type:varchar(36)" json:"admin_id"`
	ReviewedAt  time.Time `gorm:"default:null" json:"reviewed_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
