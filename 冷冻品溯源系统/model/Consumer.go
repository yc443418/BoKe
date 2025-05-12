package model

import "time"

type Consumer struct {
	ConsumerID string          `gorm:"primaryKey;type:varchar(36)" json:"consumer_id"`
	UserID     string          `gorm:"type:varchar(36);unique;constraint:OnDelete:CASCADE" json:"user_id"`
	RealName   string          `gorm:"type:varchar(50);not null" json:"real_name"`
	IDCardNo   string          `gorm:"type:varchar(18);unique;not null" json:"id_card_no"`
	CreatedAt  time.Time       `gorm:"autoCreateTime" json:"created_at"`
	Products   []FrozenProduct `gorm:"many2many:consumer_products;" json:"products"`
}
