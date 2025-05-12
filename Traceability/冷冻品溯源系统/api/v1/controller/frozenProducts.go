package controller

import (
	ErrMsg "Traceability/config"
	"Traceability/model/reqo"
	"Traceability/model/revo"
	"Traceability/service"
	"Traceability/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetProductListByAdmin(c *gin.Context) {
	data, _ := service.GetProductListByAdmin()
	utils.ResultReturnOK(c, data, ErrMsg.SUCCEE)
}
func GetProductListByFactory(c *gin.Context) {
	//接收厂家id
	factorId := c.Param("factoryId")
	//接入服务
	products, err := service.GetProductListByFactory(factorId)
	if err != nil {
		utils.ResultReturnErr(c, err, ErrMsg.ERROR)
		return
	}

	fmt.Println(products, len(products))
	//返回结果
	type ProductResponse struct {
		Products []reqo.GetFrozenProduct `json:"products"`
		Total    int                     `json:"total"`
	}

	response := ProductResponse{
		Products: products,
		Total:    len(products),
	}

	//log.Printf("处理结果: %+v", response)
	utils.ResultReturnOK(c, response, ErrMsg.SUCCEE)
}

// 获取冷冻品的详情
func GetProductDetail(c *gin.Context) {
	productId := c.Param("productId")
	data, err := service.GetProductDetails(productId)
	if err != nil {
		utils.ResultReturnErr(c, err, ErrMsg.ERROR)
		return
	}
	utils.ResultReturnOK(c, data, ErrMsg.SUCCEE)
}

// 修改冷冻品信息
func PutProductDetail(c *gin.Context) {
	productId := c.Param("productId")
	var requsestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requsestBody); err != nil {
		utils.ResultReturnErr(c, err.Error(), ErrMsg.ERROR)
		return
	}
	err := service.PutProductDetails(productId, requsestBody)
	if err != nil {
		utils.ResultReturnErr(c, err.Error(), ErrMsg.ERROR)
		return
	}
	utils.ResultReturnOK(c, "更新成功", ErrMsg.SUCCEE)
}

// 需要放到申请表中，让管理员或者监管者去同意上架
// 绑定产品，将产品写入产品表中，并将申请提交给后台进行审核，
func AddProduct(c *gin.Context) {
	var product revo.FrozenProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.ResultReturnErr(c, err.Error(), ErrMsg.ERROR)
		return
	}

	//将产家id进行uuid生成
	//将上述的信息进行sku的生成
	err := service.AddProduct(&product)
	//查看服务层错误，并将结果返回给前端
	if err != nil {
		utils.ResultReturnErr(c, err.Error(), ErrMsg.ERROR)
		return
	}
	utils.ResultReturnOK(c, "加入产品成功,等待审核", ErrMsg.SUCCEE)
}
