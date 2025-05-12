package revo

import (
	"Traceability/global"
	"Traceability/model/reqo"
	"fmt"
	"reflect"
	"time"
)

type FrozenProduct struct {
	ProductID            string    `json:"product_id"`
	ProductName          string    `json:"product_name"`
	FactoryID            string    `json:"factory_id"`
	ProductionDate       time.Time `json:"production_date" time_format:"2025-04-27 15:12:21"`
	ShelfLife            int       `json:"shelf_life"`
	BatchNumber          string    `json:"batch_number"`
	StorageCondition     string    `json:"storage_condition"`
	TransportTemperature string    `json:"transport_temperature"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	Status               int       `gorm:"type:int;default:0;comment:0待审核，1已上架，2已下架" json:"status"`
}

// ProductApplications 审核记录响应
type ProductApplications struct {
	LogID     string    `json:"log_id"`
	ProductID string    `json:"product_id"`
	FactoryID string    `json:"factory_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	AdminID   string    `json:"admin_id"`
	Status    int       `gorm:"type:int;default:0;comment:0待审核，1已上架，2已下架" json:"status"`
}

func (p *FrozenProduct) TableName() string {
	return "frozen_products"
}
func (p FrozenProduct) UpdateProduct(product *reqo.GetFrozenProduct, requsestBody map[string]interface{}) error {
	// 获取指针的值
	productValue := reflect.ValueOf(product).Elem()
	// 反射遍历这个值
	for fieldName, value := range requsestBody {
		field := productValue.FieldByName(fieldName)
		if !field.IsValid() || !field.CanSet() {
			return fmt.Errorf("无法修改字段： %s ", fieldName)
		}
		switch field.Kind() {
		case reflect.Int:
			if v, ok := value.(int); ok {
				field.SetInt(int64(v))
			} else if v, ok := value.(int64); ok {
				field.SetInt(v)
			} else {
				return fmt.Errorf("无法将 %v 转换为 int 类型", value)
			}
		case reflect.String:
			if v, ok := value.(string); ok {
				field.SetString(v)
			} else {
				return fmt.Errorf("无法将 %v 转换为 string 类型", value)
			}
		default:
			return fmt.Errorf("不支持的字段类型 %s ", fieldName)
		}
	}
	// 传递结构体指针给Save方法
	if err := global.DB.Save(productValue.Addr().Interface()).Error; err != nil {
		return fmt.Errorf("更新失败: %v", err)
	}
	return nil
}

// 将产品写入
func (p *FrozenProduct) AddProduct() (FrozenProduct, error) {
	var err error
	var product FrozenProduct
	p.Status = 0
	now := time.Now()
	p.CreatedAt = now
	fmt.Println(p)
	err = global.DB.Create(&p).Find(&product).Error
	if err != nil {
		return FrozenProduct{}, err
	}
	return product, nil
}

// 将产品添加进Frozen表中，进行管理员审核
func (p *FrozenProduct) AddFrozen(log ProductApplications) error {
	err := global.DB.Create(&log).Error
	return err
}
