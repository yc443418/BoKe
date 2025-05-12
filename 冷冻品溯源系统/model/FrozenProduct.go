package model

import "time"

type FrozenProduct struct {
	ProductID            string      `gorm:"primaryKey;type:varchar(36)" json:"product_id"`
	ProductName          string      `gorm:"type:varchar(100);not null;index" json:"product_name"`
	FactoryID            string      `gorm:"type:varchar(36);index" json:"factory_id"`
	ProductionDate       time.Time   `gorm:"type:datetime;index" json:"production_date"`
	ShelfLife            int         `gorm:"type:int;not null" json:"shelf_life"`
	BatchNumber          string      `gorm:"type:varchar(50);not null;index" json:"batch_number"`
	StorageCondition     string      `gorm:"type:varchar(100);not null" json:"storage_condition"`
	TransportTemperature string      `gorm:"type:varchar(50);not null" json:"transport_temperature"`
	CreatedAt            time.Time   `gorm:"autoCreateTime" json:"created_at"`
	Logistics            []Logistics `gorm:"foreignKey:ProductID" json:"logistics"`
	Status               int         `gorm:"type:int;default:0;comment:0待审核，1已上架，2已下架" json:"status"`
}
