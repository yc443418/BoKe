package model

import "time"

// 监管部门表
type Supervision struct {
	SupervisionID  string          `gorm:"primaryKey;type:varchar(36)" json:"supervision_id"`
	UserID         string          `gorm:"type:varchar(36);unique;constraint:OnDelete:CASCADE" json:"user_id"`
	DepartmentName string          `gorm:"type:varchar(100);not null" json:"department_name"`
	LicenseNo      string          `gorm:"type:varchar(50);unique;not null" json:"license_no"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
	Products       []FrozenProduct `gorm:"many2many:supervision_products;" json:"products"`
}
