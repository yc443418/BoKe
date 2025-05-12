package model

import "time"

// 厂家表
type Factory struct {
	FactoryID   string    `gorm:"primaryKey;type:varchar(36)" json:"factory_id"`
	UserID      string    `gorm:"type:varchar(36);unique;constraint:OnDelete:CASCADE" json:"user_id"`
	FactoryName string    `gorm:"type:varchar(100);not null" json:"factory_name"`
	LicenseNo   string    `gorm:"type:varchar(50);unique;not null" json:"license_no"` //证书
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
