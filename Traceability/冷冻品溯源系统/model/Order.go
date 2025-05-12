package model

import "time"

type Order struct {
	OrderID     string    `gorm:"primaryKey;type:varchar(36)" json:"order_id"`
	UserID      string    `gorm:"type:varchar(36);index" json:"user_id"`
	ProductID   string    `gorm:"type:varchar(36);index" json:"product_id"`
	DealerID    string    `gorm:"type:varchar(36);index" json:"dealer_id"`
	LogisticsID string    `gorm:"type:varchar(36);index" json:"logistics_id"`
	Quantity    int       `gorm:"type:int;check:quantity > 0" json:"quantity"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	OrderTime   time.Time `gorm:"type:datetime;index" json:"order_time" time_format:"2006-01-02 15:04:05"`
	Status      int       `gorm:"type:int;default:1;comment:1待支付 2已发货 3已完成" json:"status"`
}
