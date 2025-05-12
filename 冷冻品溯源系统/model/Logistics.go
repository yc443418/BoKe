package model

import "time"

// 物流表
type Logistics struct {
	LogisticsID     string    `gorm:"primaryKey;type:varchar(36)" json:"logistics_id"`
	ProductID       string    `gorm:"type:varchar(36);index" json:"product_id"`
	DealerID        string    `gorm:"type:varchar(36);index" json:"dealer_id"`
	TrackingNumber  string    `gorm:"type:varchar(50);unique;not null" json:"tracking_number"` //物流号
	StartLocation   string    `gorm:"type:varchar(100);not null" json:"start_location"`        //起始位置
	EndLocation     string    `gorm:"type:varchar(100);not null" json:"end_location"`          //终点
	CurrentTemp     float64   `gorm:"type:decimal(5,2);not null" json:"current_temp"`          //温度
	CurrentHumidity float64   `gorm:"type:decimal(5,2);not null" json:"current_humidity"`      //湿度
	Status          int       `gorm:"type:int;default:1;comment:1运输中 2已签收" json:"status"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
