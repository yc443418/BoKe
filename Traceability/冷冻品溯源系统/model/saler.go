package model

import "time"

type Dealer struct {
	DealerID   string    `gorm:"primaryKey;type:varchar(36)" json:"dealer_id"`
	UserID     string    `gorm:"type:varchar(36);unique;constraint:OnDelete:CASCADE" json:"user_id"`
	DealerName string    `gorm:"type:varchar(100);not null" json:"dealer_name"`
	LicenseNo  string    `gorm:"type:varchar(50);unique;not null" json:"license_no"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
