package model

import "time"

// 用于商品商上架的审核
type ProductApplication struct {
	LogId     string    `gorm:"primaryKey;type:varchar(36)" json:"log_id"`
	ProductID string    `gorm:"type:varchar(36);index" json:"product_id"`
	FactoryID string    `gorm:"type:varchar(36);unique;constraint:OnDelete:CASCADE" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:null" json:"reviewed_at"`
	AdminID   string    `gorm:"type:varchar(36)" json:"admin_id"`
	Status    int       `gorm:"type:int;default:0" json:"status"` //默认1为同意上架 0为待审核  2为下架
}

func (p *ProductApplication) TableName() string {
	return "product_applications"
}
