package service

import (
	"Traceability/model/reqo"
	"Traceability/model/revo"
	"Traceability/utils"
	"fmt"
	"time"
)

func GetProductListByAdmin() ([]reqo.GetFrozenProduct, error) {
	var FrozenProduct reqo.GetFrozenProduct
	data, err := FrozenProduct.GetProductListByAdmin()
	if err != nil {
		return []reqo.GetFrozenProduct{}, err
	}
	return data, nil
}

func GetProductListByFactory(factoryId string) ([]reqo.GetFrozenProduct, error) {
	//接收厂家id
	var FrozenProduct reqo.GetFrozenProduct
	data, err := FrozenProduct.GetProductListByFactory(factoryId)
	if err != nil {
		return []reqo.GetFrozenProduct{}, err
	}
	return data, nil
}

// 获取冷冻品详情
func GetProductDetails(productId string) (*reqo.GetFrozenProduct, error) {
	var FrozenProduct *reqo.GetFrozenProduct
	data, err := FrozenProduct.GetProductDetail(productId)
	if err != nil {
		return FrozenProduct, err
	}
	return data, nil
}

func PutProductDetails(productId string, requsestBody map[string]interface{}) error {
	var FrozenProduct reqo.GetFrozenProduct
	// 查询该id的值
	data, err := FrozenProduct.GetProductDetail(productId)
	if err != nil {
		return err
	}
	if data == nil {
		return fmt.Errorf("未找到产品ID为 %s 的冷冻品信息", productId)
	}
	var Respone revo.FrozenProduct
	err = Respone.UpdateProduct(data, requsestBody)
	if err != nil {
		return err
	}
	return nil
}

// 将产品放入Frozen表中，当让status初始话为待审核，只有等后台管理员修改该字段
func AddProduct(procuct *revo.FrozenProduct) (err error) {
	//产品上架前均需要审核
	result, err := procuct.AddProduct()
	if err != nil {
		return err
	}

	//将其写入申请表中
	log := revo.ProductApplications{
		LogID:     utils.GenerateUUID(),
		ProductID: result.ProductID,
		FactoryID: result.FactoryID,
		CreatedAt: time.Now(),
		Status:    0,
	}

	err = procuct.AddFrozen(log)
	if err != nil {
		return err
	}
	return nil
}
