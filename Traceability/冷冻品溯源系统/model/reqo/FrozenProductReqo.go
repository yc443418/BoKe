package reqo

import (
	"Traceability/global"
	"time"
)

type GetFrozenProduct struct {
	ProductID            string    `json:"product_id"`
	ProductName          string    `json:"product_name"`
	FactoryID            string    `json:"factory_id"`
	ProductionDate       time.Time `json:"production_date" time_format:"2006-01-02"`
	ShelfLife            int       ` json:"shelf_life"`
	BatchNumber          string    `json:"batch_number"`
	StorageCondition     string    `json:"storage_condition"`
	TransportTemperature string    `json:"transport_temperature"`
	CreatedAt            time.Time `json:"created_at"`
	Status               int       `json:"status"`
}

func (p *GetFrozenProduct) TableName() string {
	return "frozen_products"
}

// 后台管理员的冷冻品管理
func (p *GetFrozenProduct) GetProductListByAdmin() ([]GetFrozenProduct, error) {
	var products []GetFrozenProduct
	global.DB.Debug().Find(&products)
	return products, nil
}
func (p *GetFrozenProduct) GetProductListByFactory(factoryId string) ([]GetFrozenProduct, error) {
	var products []GetFrozenProduct
	global.DB.Debug().Where("factory_id = ?", factoryId).Find(&products)
	return products, nil
}

// 获取单个冷冻品详情
func (p *GetFrozenProduct) GetProductDetail(productId string) (*GetFrozenProduct, error) {
	var product *GetFrozenProduct
	global.DB.Debug().Where("product_id = ?", productId).Find(&product)
	return product, nil
}
